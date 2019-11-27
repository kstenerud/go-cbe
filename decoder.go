package cbe

import (
	"fmt"
	"time"
)

// Callback functions that must be present in the receiver object.
type DecoderCallbacks interface {
	OnNil() error
	OnBool(value bool) error
	OnPositiveInt(value uint64) error
	OnNegativeInt(value uint64) error
	OnFloat(value float64) error
	OnDate(value time.Time) error
	OnTime(value time.Time) error
	OnTimestamp(value time.Time) error
	OnListBegin() error
	OnMapBegin() error
	OnMetadataBegin() error
	OnContainerEnd() error
	OnBytesBegin(byteCount uint64) error
	OnStringBegin(byteCount uint64) error
	OnURIBegin(byteCount uint64) error
	OnCommentBegin(byteCount uint64) error
	OnArrayData(bytes []byte) error
}

// Biggest item is timestamp (10 bytes), longest tz is "America/Argentina/ComodRivadavia"
const maxPartialReadLength = 50

type decoderError struct {
	err error
}

type callbackError struct {
	err error
}

type containerData struct {
	inlineContainerPhase        containerPhase
	inlineContainerInitialized bool
	depth                      int
	currentPhase                []containerPhase
	hasProcessedMapKey         []bool
}

type arrayData struct {
	currentType        arrayType
	remainingByteCount int64
}

type Decoder struct {
	streamOffset     int64
	buffer           *decodeBuffer
	mainBuffer       *decodeBuffer
	underflowBuffer  *decodeBuffer
	container        containerData
	array            arrayData
	callbacks        DecoderCallbacks
	firstItemDecoded bool
	charValidator    Utf8Validator
}

func panicOnCallbackError(err error) {
	if err != nil {
		panic(callbackError{err})
	}
}

func (this *Decoder) isExpectingMapKey() bool {
	return this.container.currentPhase[this.container.depth] == containerPhaseMap &&
		!this.container.hasProcessedMapKey[this.container.depth]
}

func (this *Decoder) isExpectingMapValue() bool {
	return this.container.currentPhase[this.container.depth] == containerPhaseMap &&
		this.container.hasProcessedMapKey[this.container.depth]
}

func (this *Decoder) flipMapKeyValueState() {
	this.container.hasProcessedMapKey[this.container.depth] = !this.container.hasProcessedMapKey[this.container.depth]
}

func (this *Decoder) assertNotExpectingMapKey(keyType string) {
	if this.isExpectingMapKey() {
		panic(decoderError{fmt.Errorf("Cannot use type %v as a map key", keyType)})
	}
}

func (this *Decoder) containerBegin(nextPhase containerPhase) {
	if this.container.depth+1 >= len(this.container.currentPhase) {
		panic(decoderError{fmt.Errorf("Exceeded max container depth of %v", len(this.container.currentPhase))})
	}
	this.container.depth++
	this.container.currentPhase[this.container.depth] = nextPhase
	this.container.hasProcessedMapKey[this.container.depth] = false
}

func (this *Decoder) containerEnd() containerPhase {
	if this.container.depth <= 0 {
		panic(decoderError{fmt.Errorf("Got container end but not in a container")})
	}
	if this.container.inlineContainerPhase != containerPhaseNone && this.container.depth <= 1 {
		panic(decoderError{fmt.Errorf("Got container end but not in a container")})
	}
	if this.isExpectingMapValue() {
		panic(decoderError{fmt.Errorf("Expecting map value for already processed key")})
	}

	this.container.depth--
	return this.container.currentPhase[this.container.depth+1]
}

func (this *Decoder) arrayBegin(newArrayType arrayType, length int64) {
	this.array.currentType = newArrayType
	this.charValidator.Reset()
	this.array.remainingByteCount = length

	switch newArrayType {
	case arrayTypeBytes:
		panicOnCallbackError(this.callbacks.OnBytesBegin(uint64(length)))
	case arrayTypeComment:
		panicOnCallbackError(this.callbacks.OnCommentBegin(uint64(length)))
	case arrayTypeString:
		panicOnCallbackError(this.callbacks.OnStringBegin(uint64(length)))
	case arrayTypeURI:
		panicOnCallbackError(this.callbacks.OnURIBegin(uint64(length)))
	default:
		panic(fmt.Errorf("BUG: Unhandled array type: %v", newArrayType))
	}
}

func (this *Decoder) decodeArrayLength(buffer *decodeBuffer) {
	this.array.remainingByteCount = buffer.DecodeArrayLength()
}

func (this *Decoder) decodeArrayData() {
	if this.array.currentType == arrayTypeNone {
		return
	}

	if this.array.remainingByteCount > 0 {
		decodeByteCount := this.buffer.RemainingByteCount()
		if int64(decodeByteCount) > this.array.remainingByteCount {
			decodeByteCount = int(this.array.remainingByteCount)
		}
		bytes := this.buffer.DecodeBytes(decodeByteCount)
		if this.array.currentType == arrayTypeString || this.array.currentType == arrayTypeComment {
			for _, ch := range bytes {
				if err := this.charValidator.AddByte(int(ch)); err != nil {
					panic(decoderError{err})
				}
				if this.charValidator.IsCompleteCharacter() && this.array.currentType == arrayTypeComment {
					if err := ValidateCommentCharacter(this.charValidator.Character()); err != nil {
						panic(decoderError{err})
					}
				}
			}
		}
		this.array.remainingByteCount -= int64(decodeByteCount)

		panicOnCallbackError(this.callbacks.OnArrayData(bytes))
		this.buffer.Commit()

		if this.array.remainingByteCount > 0 {
			panic(notEnoughBytesToDecodeArrayData(this.array.remainingByteCount))
		}
	}

	this.array.currentType = arrayTypeNone
	this.flipMapKeyValueState()
}

func (this *Decoder) decodeStringOfLength(length int64) {
	this.arrayBegin(arrayTypeString, length)
	this.decodeArrayData()
}

func (this *Decoder) decodeObject(dataType typeField) {
	asSmallInt := int8(dataType)
	if int64(asSmallInt) >= smallIntMin && int64(asSmallInt) <= smallIntMax {
		if asSmallInt >= 0 {
			panicOnCallbackError(this.callbacks.OnPositiveInt(uint64(asSmallInt)))
		} else {
			panicOnCallbackError(this.callbacks.OnNegativeInt(uint64(-asSmallInt)))
		}
		this.buffer.Commit()
		this.flipMapKeyValueState()
		return
	}

	switch dataType {
	case typeTrue:
		panicOnCallbackError(this.callbacks.OnBool(true))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeFalse:
		panicOnCallbackError(this.callbacks.OnBool(false))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeFloat32:
		panicOnCallbackError(this.callbacks.OnFloat(float64(this.buffer.DecodeFloat32())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeFloat64:
		panicOnCallbackError(this.callbacks.OnFloat(this.buffer.DecodeFloat64()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeDecimal:
		panicOnCallbackError(this.callbacks.OnFloat(this.buffer.DecodeFloat()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typePosInt8:
		panicOnCallbackError(this.callbacks.OnPositiveInt(uint64(this.buffer.DecodeUint8())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typePosInt16:
		panicOnCallbackError(this.callbacks.OnPositiveInt(uint64(this.buffer.DecodeUint16())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typePosInt32:
		panicOnCallbackError(this.callbacks.OnPositiveInt(uint64(this.buffer.DecodeUint32())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typePosInt64:
		panicOnCallbackError(this.callbacks.OnPositiveInt(this.buffer.DecodeUint64()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typePosInt:
		panicOnCallbackError(this.callbacks.OnPositiveInt(this.buffer.DecodeUint()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeNegInt8:
		panicOnCallbackError(this.callbacks.OnNegativeInt(uint64(this.buffer.DecodeUint8())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeNegInt16:
		panicOnCallbackError(this.callbacks.OnNegativeInt(uint64(this.buffer.DecodeUint16())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeNegInt32:
		panicOnCallbackError(this.callbacks.OnNegativeInt(uint64(this.buffer.DecodeUint32())))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeNegInt64:
		panicOnCallbackError(this.callbacks.OnNegativeInt(this.buffer.DecodeUint64()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeNegInt:
		panicOnCallbackError(this.callbacks.OnNegativeInt(this.buffer.DecodeUint()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeDate:
		panicOnCallbackError(this.callbacks.OnDate(this.buffer.DecodeDate()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeTime:
		panicOnCallbackError(this.callbacks.OnTime(this.buffer.DecodeTime()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeTimestamp:
		panicOnCallbackError(this.callbacks.OnTimestamp(this.buffer.DecodeTimestamp()))
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeNil:
		this.assertNotExpectingMapKey("nil")
		panicOnCallbackError(this.callbacks.OnNil())
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typePadding:
		// Ignore
	case typeList:
		this.assertNotExpectingMapKey("list")
		this.containerBegin(containerPhaseList)
		panicOnCallbackError(this.callbacks.OnListBegin())
		this.buffer.Commit()
	case typeMap:
		this.assertNotExpectingMapKey("map")
		this.containerBegin(containerPhaseMap)
		panicOnCallbackError(this.callbacks.OnMapBegin())
		this.buffer.Commit()
	case typeMetadata:
		this.assertNotExpectingMapKey("metadata")
		this.containerBegin(containerPhaseMetadata)
		panicOnCallbackError(this.callbacks.OnMetadataBegin())
		this.buffer.Commit()
	case typeEndContainer:
		this.containerEnd()
		panicOnCallbackError(this.callbacks.OnContainerEnd())
		this.buffer.Commit()
		this.flipMapKeyValueState()
	case typeBytes:
		this.arrayBegin(arrayTypeBytes, this.buffer.DecodeArrayLength())
		this.decodeArrayData()
	case typeComment:
		this.arrayBegin(arrayTypeComment, this.buffer.DecodeArrayLength())
		this.decodeArrayData()
	case typeURI:
		this.arrayBegin(arrayTypeURI, this.buffer.DecodeArrayLength())
		this.decodeArrayData()
	case typeString:
		this.arrayBegin(arrayTypeString, this.buffer.DecodeArrayLength())
		this.decodeArrayData()
	case typeString0:
		this.decodeStringOfLength(0)
	case typeString1:
		this.decodeStringOfLength(1)
	case typeString2:
		this.decodeStringOfLength(2)
	case typeString3:
		this.decodeStringOfLength(3)
	case typeString4:
		this.decodeStringOfLength(4)
	case typeString5:
		this.decodeStringOfLength(5)
	case typeString6:
		this.decodeStringOfLength(6)
	case typeString7:
		this.decodeStringOfLength(7)
	case typeString8:
		this.decodeStringOfLength(8)
	case typeString9:
		this.decodeStringOfLength(9)
	case typeString10:
		this.decodeStringOfLength(10)
	case typeString11:
		this.decodeStringOfLength(11)
	case typeString12:
		this.decodeStringOfLength(12)
	case typeString13:
		this.decodeStringOfLength(13)
	case typeString14:
		this.decodeStringOfLength(14)
	case typeString15:
		this.decodeStringOfLength(15)
	}
	// TODO: 128 bit and decimal
}

func (this *Decoder) beginInlineContainer() {
	if this.container.inlineContainerPhase != containerPhaseNone && !this.container.inlineContainerInitialized {
		this.containerBegin(this.container.inlineContainerPhase)
		switch this.container.inlineContainerPhase {
		case containerPhaseList:
			panicOnCallbackError(this.callbacks.OnListBegin())
		case containerPhaseMap:
			panicOnCallbackError(this.callbacks.OnMapBegin())
		}
		this.container.inlineContainerInitialized = true
	}
}

func (this *Decoder) endInlineContainer() {
	if this.container.inlineContainerInitialized {
		this.callbacks.OnContainerEnd()
		this.container.depth--
	}
}

func (this *Decoder) assertOnlyOneTopLevelObject() {
	if this.container.depth == 0 && this.firstItemDecoded {
		panic(decoderError{fmt.Errorf("Extra top level object detected")})
	}
}

// ----------
// Public API
// ----------

func NewCbeDecoder(inlineContainerType InlineContainerType, maxContainerDepth int, callbacks DecoderCallbacks) *Decoder {
	this := new(Decoder)

	switch inlineContainerType {
	case InlineContainerTypeList:
		this.container.inlineContainerPhase = containerPhaseList
		maxContainerDepth++
	case InlineContainerTypeMap:
		this.container.inlineContainerPhase = containerPhaseMap
		maxContainerDepth++
	case InlineContainerTypeNone:
	// Nothing to do
	default:
		panic(fmt.Errorf("Unhandled container type: %v", inlineContainerType))
	}

	this.underflowBuffer = NewDecodeBuffer(make([]byte, maxPartialReadLength))
	this.underflowBuffer.Clear()
	this.mainBuffer = NewDecodeBuffer(make([]byte, 0))
	this.buffer = this.mainBuffer
	this.callbacks = callbacks
	this.container.currentPhase = make([]containerPhase, maxContainerDepth)
	this.container.hasProcessedMapKey = make([]bool, maxContainerDepth)
	return this
}

// Feed bytes into the decoder to be decoded.
func (this *Decoder) Feed(bytesToDecode []byte) (err error) {
	defer func() {
		this.streamOffset += int64(this.buffer.lastCommitPosition)
		if r := recover(); r != nil {
			switch r.(type) {
			case notEnoughBytesToDecodeType:
				// Return as if nothing's wrong
			case notEnoughBytesToDecodeArrayData:
				// Return as if nothing's wrong
			case notEnoughBytesToDecodeObject:
				this.underflowBuffer.AddContents(this.mainBuffer.GetUncommittedBytes())
				this.buffer = this.underflowBuffer
				this.mainBuffer.Clear()
			case callbackError:
				err = fmt.Errorf("cbe: offset %v: Error from callback: %v", this.streamOffset, r.(callbackError).err)
			case decoderError:
				err = fmt.Errorf("cbe: offset %v: Decode error: %v", this.streamOffset, r.(decoderError).err)
			default:
				// Unexpected panics are passed as-is
				panic(r)
			}
		}
	}()

	this.buffer.Rollback()
	this.mainBuffer.ReplaceBuffer(bytesToDecode)

	this.beginInlineContainer()

	if this.buffer == this.underflowBuffer && this.buffer.RemainingByteCount() > 0 {
		underflowByteCount := len(this.underflowBuffer.data)
		bytesFilled := this.buffer.FillFromBuffer(this.mainBuffer, maxPartialReadLength)
		this.mainBuffer.lastCommitPosition += bytesFilled
		this.mainBuffer.position = this.mainBuffer.lastCommitPosition
		objectType := this.buffer.DecodeType()
		if objectType != typePadding {
			this.assertOnlyOneTopLevelObject()
			this.decodeObject(objectType)
			this.firstItemDecoded = true
		}
		mainBytesUsed := this.underflowBuffer.lastCommitPosition - underflowByteCount
		this.mainBuffer.lastCommitPosition = mainBytesUsed
		this.mainBuffer.position = this.mainBuffer.lastCommitPosition

		this.underflowBuffer.Clear()
		this.buffer = this.mainBuffer
	}

	// TODO: Does this handle end of buffer?
	this.decodeArrayData()

	for {
		objectType := this.buffer.DecodeType()
		if objectType != typePadding {
			this.assertOnlyOneTopLevelObject()
			this.decodeObject(objectType)
			this.firstItemDecoded = true
		}
	}

	return err
}

// End the decoding process, doing some final structural tests to make sure it's valid.
func (this *Decoder) End() error {
	this.endInlineContainer()

	if this.container.depth > 0 {
		return fmt.Errorf("Document still has %v open container(s)", this.container.depth)
	}
	if this.array.remainingByteCount > 0 {
		return fmt.Errorf("Array is still open, expecting %d more bytes", this.array.remainingByteCount)
	}
	if this.buffer == this.underflowBuffer {
		return fmt.Errorf("Document has not been completely decoded. %v bytes of underflow data remain", len(this.underflowBuffer.data))
	}
	return nil
}

// Convenience function to decode an entire document in a single call.
func (this *Decoder) Decode(document []byte) error {
	if err := this.Feed(document); err != nil {
		return err
	}
	return this.End()
}
