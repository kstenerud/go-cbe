// Copyright 2019 Karl Stenerud
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

// Imposes the structural rules that enforce a well-formed concise encoding
// document.
package rules

import (
	"fmt"
	"unicode/utf8"

	"github.com/kstenerud/go-concise-encoding/internal/chars"

	"github.com/kstenerud/go-concise-encoding/events"
	"github.com/kstenerud/go-concise-encoding/internal/common"
)

func (_this *Context) beginArray(arrayType events.ArrayType, rule EventRule, dataType DataType, maxByteCount uint64, validatorFunc func([]byte)) {
	_this.arrayTotalByteCount = 0
	_this.builtArrayBuffer = _this.builtArrayBuffer[:0]
	_this.utf8RemainderBuffer = _this.utf8RemainderBacking[:0]

	_this.stackRule(rule, dataType)
	_this.arrayType = arrayType
	_this.arrayMaxByteCount = maxByteCount
	_this.ValidateArrayDataFunc = validatorFunc
}

func (_this *Context) endArray() {
	_this.endContainerLike()
}

func (_this *Context) validateArrayTotalByteCount(byteCount uint64, maxByteCount uint64) {
	if byteCount > _this.arrayMaxByteCount && maxByteCount > 0 {
		panic(fmt.Errorf("%v byte count %d exceeds maximum of %d", _this.arrayType, byteCount, maxByteCount))
	}
}

func (_this *Context) markUpcomingChunkByteCount(byteCount uint64) {
	_this.arrayTotalByteCount += byteCount
	_this.validateArrayTotalByteCount(_this.arrayTotalByteCount, _this.arrayMaxByteCount)
}

func (_this *Context) MarkCompletedChunkByteCount(byteCount uint64) {
	_this.chunkActualByteCount += byteCount
	if _this.chunkActualByteCount > _this.chunkExpectedByteCount {
		panic(fmt.Errorf("Expected array chunk to have %d bytes, but got %d bytes", _this.chunkExpectedByteCount, _this.chunkActualByteCount))
	}
}

func (_this *Context) AddBuiltArrayBytes(bytes []byte) {
	_this.builtArrayBuffer = append(_this.builtArrayBuffer, bytes...)
}

func (_this *Context) GetBuiltArrayAsString() string {
	return string(_this.builtArrayBuffer)
}

// Takes streamed byte data, partitioning it into buffers that have only
// complete UTF-8 runes in them. This handles cases where a rune is split
// across multiple incoming data buffers.
//
// This function performs no allocations.
//
// Return:
// - firstRuneBytes will contain either a single complete rune (the next in
//   order in the stream) or an empty slice.
// - nextRunesBytes will contain the remaining complete runes. The last
//   incomplete rune, if any, will be stripped out and buffered for the next
//   call.
func (_this Context) StreamStringData(data []byte) (firstRuneBytes []byte, nextRunesBytes []byte) {
	nextRunesBytes = data

	remainderLength := len(_this.utf8RemainderBuffer)
	if remainderLength > 0 {
		requiredByteCount := chars.CalculateRuneByteCount(_this.utf8RemainderBuffer[0])
		_this.utf8RemainderBuffer = _this.utf8RemainderBuffer[:requiredByteCount]
		bytesCopied := copy(_this.utf8RemainderBuffer[remainderLength:], nextRunesBytes)
		nextRunesBytes = nextRunesBytes[bytesCopied:]
		if remainderLength+bytesCopied < requiredByteCount {
			_this.utf8RemainderBuffer = _this.utf8RemainderBuffer[:remainderLength+bytesCopied]
			return
		}
		firstRuneBytes = _this.utf8RemainderBuffer
		_this.utf8RemainderBuffer = _this.utf8RemainderBuffer[:0]
	}

	lastIndex, isComplete := chars.IndexOfLastRuneStart(nextRunesBytes)
	if !isComplete {
		remainderBytes := nextRunesBytes[lastIndex:]
		_this.utf8RemainderBuffer = _this.utf8RemainderBacking[:len(remainderBytes)]
		copy(_this.utf8RemainderBuffer, remainderBytes)
		nextRunesBytes = nextRunesBytes[:lastIndex]
	}
	return
}

// Chunking

func (_this *Context) BeginArrayAnyType(arrayType events.ArrayType) {
	switch arrayType {
	case events.ArrayTypeString:
		_this.beginArray(arrayType, &stringRule, DataTypeKeyable, _this.opts.MaxStringByteLength, _this.ValidateContentsString)
	case events.ArrayTypeResourceID:
		_this.beginArray(arrayType, &stringRule, DataTypeKeyable, _this.opts.MaxResourceIDByteLength, _this.ValidateContentsRID)
	case events.ArrayTypeCustomText:
		_this.beginArray(arrayType, &stringRule, DataTypeAnyType, _this.opts.MaxArrayByteLength, _this.ValidateContentsCustomText)
	default:
		_this.beginArray(arrayType, &arrayRule, DataTypeAnyType, _this.opts.MaxArrayByteLength, _this.ValidateNothing)
	}
}

func (_this *Context) BeginArrayKeyable(arrayType events.ArrayType) {
	_this.AssertArrayTypeKeyable(arrayType)
	_this.BeginArrayAnyType(arrayType)
}

func (_this *Context) BeginChunkAnyType(elemCount uint64, moreChunksFollow bool) {
	_this.chunkExpectedByteCount = common.ElementCountToByteCount(_this.arrayType.ElementSize(), elemCount)
	_this.markUpcomingChunkByteCount(_this.chunkExpectedByteCount)
	_this.chunkActualByteCount = 0
	_this.moreChunksFollow = moreChunksFollow
	if elemCount > 0 {
		_this.changeRule(&arrayChunkRule)
	} else {
		_this.EndChunkAnyType()
	}
}

func (_this *Context) EndChunkAnyType() {
	if _this.moreChunksFollow {
		_this.changeRule(&arrayRule)
	} else {
		_this.endArray()
	}
}

func (_this *Context) BeginArrayString(arrayType events.ArrayType) {
	_this.AssertArrayTypeString(arrayType)
	_this.beginArray(arrayType, &stringRule, DataTypeKeyable, _this.opts.MaxStringByteLength, _this.ValidateContentsString)
}

func (_this *Context) BeginArrayRID(arrayType events.ArrayType) {
	_this.AssertArrayTypeRID(arrayType)
	_this.beginArray(arrayType, &stringRule, DataTypeKeyable, _this.opts.MaxResourceIDByteLength, _this.ValidateContentsRID)
}

func (_this *Context) BeginArrayRIDReference(arrayType events.ArrayType) {
	_this.AssertArrayTypeRID(arrayType)
	_this.beginArray(arrayType, &stringRule, DataTypeAnyType, _this.opts.MaxResourceIDByteLength, _this.ValidateContentsRID)
}

func (_this *Context) BeginArrayComment(arrayType events.ArrayType) {
	_this.AssertArrayTypeString(arrayType)
	_this.beginArray(arrayType, &stringRule, DataTypeKeyable, _this.opts.MaxArrayByteLength, _this.ValidateContentsComment)
}

func (_this *Context) BeginChunkString(elemCount uint64, moreChunksFollow bool) {
	_this.chunkExpectedByteCount = elemCount
	_this.markUpcomingChunkByteCount(_this.chunkExpectedByteCount)
	_this.chunkActualByteCount = 0
	_this.moreChunksFollow = moreChunksFollow
	if elemCount > 0 {
		_this.changeRule(&stringChunkRule)
	} else {
		_this.EndChunkString()
	}
}

func (_this *Context) EndChunkString() {
	if _this.moreChunksFollow {
		_this.changeRule(&stringRule)
	} else {
		if len(_this.utf8RemainderBuffer) > 0 {
			panic(fmt.Errorf("Incomplete UTF-8 string"))
		}
		_this.endArray()
	}
}

func (_this *Context) BeginStringBuilder(arrayType events.ArrayType, completedValidatorFunc func([]byte)) {
	_this.AssertArrayTypeString(arrayType)
	_this.beginArray(arrayType, &stringBuilderRule, DataTypeKeyable, _this.opts.MaxStringByteLength, completedValidatorFunc)
}

func (_this *Context) BeginChunkStringBuilder(elemCount uint64, moreChunksFollow bool) {
	_this.chunkExpectedByteCount = elemCount
	_this.markUpcomingChunkByteCount(_this.chunkExpectedByteCount)
	_this.chunkActualByteCount = 0
	_this.moreChunksFollow = moreChunksFollow
	if elemCount > 0 {
		_this.changeRule(&stringBuilderChunkRule)
	} else {
		_this.EndChunkStringBuilder()
	}
}

func (_this *Context) EndChunkStringBuilder() {
	if _this.moreChunksFollow {
		_this.changeRule(&stringBuilderRule)
	} else {
		_this.ValidateArrayDataFunc(_this.builtArrayBuffer)
		_this.endArray()
	}
}

// Validation

func (_this *Context) ValidateFullArrayAnyType(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	switch arrayType {
	case events.ArrayTypeString:
		_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
		_this.ValidateLengthString(uint64(len(data)))
		_this.ValidateContentsString(data)
	case events.ArrayTypeResourceID:
		_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
		_this.ValidateLengthRID(uint64(len(data)))
		_this.ValidateContentsRID(data)
	case events.ArrayTypeCustomText:
		_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
		_this.ValidateLengthAnyType(uint64(len(data)))
		_this.ValidateContentsString(data)
	default:
		_this.ValidateByteCountForType(arrayType, elementCount, uint64(len(data)))
		_this.ValidateLengthAnyType(uint64(len(data)))
	}
}

func (_this *Context) ValidateFullArrayStringlike(arrayType events.ArrayType, data string) {
	switch arrayType {
	case events.ArrayTypeString:
		_this.ValidateLengthString(uint64(len(data)))
		_this.ValidateContentsStringlike(data)
	case events.ArrayTypeResourceID:
		_this.ValidateLengthRID(uint64(len(data)))
		_this.ValidateContentsRIDString(data)
	case events.ArrayTypeCustomText:
		_this.ValidateLengthAnyType(uint64(len(data)))
		_this.ValidateContentsStringlike(data)
	default:
		_this.ValidateLengthAnyType(uint64(len(data)))
	}
}

func (_this *Context) ValidateFullArrayStringlikeKeyable(arrayType events.ArrayType, data string) {
	_this.AssertArrayTypeKeyable(arrayType)
	_this.ValidateFullArrayStringlike(arrayType, data)
}

func (_this *Context) ValidateFullArrayKeyable(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.AssertArrayTypeKeyable(arrayType)
	_this.ValidateFullArrayAnyType(arrayType, elementCount, data)
}

func (_this *Context) ValidateFullArrayString(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
	_this.ValidateLengthString(uint64(len(data)))
	_this.ValidateContentsString(data)
}

func (_this *Context) ValidateFullArrayComment(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
	_this.ValidateLengthAnyType(uint64(len(data)))
	_this.ValidateContentsComment(data)
}

func (_this *Context) ValidateFullArrayCommentString(arrayType events.ArrayType, data string) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateLengthAnyType(uint64(len(data)))
	_this.ValidateContentsCommentString(data)
}

func (_this *Context) ValidateFullArrayMarkupContents(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
	_this.ValidateLengthAnyType(uint64(len(data)))
	_this.ValidateContentsComment(data)
}

func (_this *Context) ValidateFullArrayMarkupContentsString(arrayType events.ArrayType, data string) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateLengthAnyType(uint64(len(data)))
	_this.ValidateContentsCommentString(data)
}

func (_this *Context) ValidateFullArrayRID(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.AssertArrayTypeRID(arrayType)
	_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
	_this.ValidateLengthRID(uint64(len(data)))
	_this.ValidateContentsRID(data)
}

func (_this *Context) ValidateFullArrayMarkerID(arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateByteCount1BPE(elementCount, uint64(len(data)))
	_this.ValidateLengthMarkerID(uint64(len(data)))
	_this.ValidateContentsMarkerID(data)
}

func (_this *Context) ValidateFullArrayMarkerIDString(arrayType events.ArrayType, data string) {
	_this.AssertArrayTypeString(arrayType)
	_this.ValidateLengthMarkerID(uint64(len(data)))
	_this.ValidateContentsMarkerIDString(data)
}

func (_this *Context) ValidateByteCount1BPE(elementCount uint64, byteCount uint64) {
	if byteCount != elementCount {
		panic(fmt.Errorf("Expected string length of %d bytes but got %d bytes",
			elementCount, byteCount))
	}
}

func (_this *Context) ValidateByteCountForType(arrayType events.ArrayType, elementCount uint64, byteCount uint64) {
	expectedByteCount := common.ElementCountToByteCount(arrayType.ElementSize(), elementCount)
	if byteCount != expectedByteCount {
		panic(fmt.Errorf("Expected %d bytes (%d elements of %d bits) but got %d bytes",
			expectedByteCount, elementCount, arrayType.ElementSize(), byteCount))
	}
}

func (_this *Context) AssertArrayTypeKeyable(arrayType events.ArrayType) {
	if !isKeyableType(arrayType) {
		panic(fmt.Errorf("Expected a keyable array type but got %v", arrayType))
	}
}

func (_this *Context) AssertArrayTypeString(arrayType events.ArrayType) {
	if arrayType != events.ArrayTypeString {
		panic(fmt.Errorf("Expected a string array type but got %v", arrayType))
	}
}

func (_this *Context) AssertArrayTypeRID(arrayType events.ArrayType) {
	if arrayType != events.ArrayTypeResourceID {
		panic(fmt.Errorf("Expected a resource ID array type but got %v", arrayType))
	}
}

func (_this *Context) ValidateLengthAnyType(length uint64) {
	if length > _this.opts.MaxArrayByteLength && _this.opts.MaxArrayByteLength > 0 {
		panic(fmt.Errorf("Array byte length %d is greater than the maximum of %d", length, _this.opts.MaxArrayByteLength))
	}
}

func (_this *Context) ValidateLengthString(length uint64) {
	if length > _this.opts.MaxStringByteLength && _this.opts.MaxStringByteLength > 0 {
		panic(fmt.Errorf("String byte length %d is greater than the maximum of %d", length, _this.opts.MaxStringByteLength))
	}
}

func (_this *Context) ValidateLengthRID(length uint64) {
	if length > _this.opts.MaxResourceIDByteLength && _this.opts.MaxResourceIDByteLength > 0 {
		panic(fmt.Errorf("Resource ID byte length %d is greater than the maximum of %d", length, _this.opts.MaxResourceIDByteLength))
	}
}

func (_this *Context) ValidateLengthMarkerID(length uint64) {
	if length > maxMarkerIDByteCount {
		panic(fmt.Errorf("Marker ID byte length %d is greater than the maximum of %d", length, maxMarkerIDByteCount))
	}
}

func (_this *Context) ValidateContentsString(contents []byte) {
	if !utf8.Valid(contents) {
		panic(fmt.Errorf("String is not valid UTF-8: %v", string(contents)))
	}
}

func (_this *Context) ValidateContentsStringlike(contents string) {
	if !utf8.ValidString(contents) {
		panic(fmt.Errorf("String is not valid UTF-8: %v", string(contents)))
	}
}

func (_this *Context) ValidateContentsCustomText(contents []byte) {
	_this.ValidateContentsString(contents)
}

func (_this *Context) ValidateContentsComment(contents []byte) {
	// TODO: More specific validation
	_this.ValidateContentsString(contents)
}

func (_this *Context) ValidateContentsCommentString(contents string) {
	// TODO: More specific validation
	_this.ValidateContentsStringlike(contents)
}

func (_this *Context) ValidateContentsRID(contents []byte) {
	// TODO: More specific validation
	_this.ValidateContentsString(contents)
}

func (_this *Context) ValidateContentsRIDString(contents string) {
	// TODO: More specific validation
	_this.ValidateContentsStringlike(contents)
}

func (_this *Context) ValidateContentsMarkerID(contents []byte) {
	_this.ValidateContentsMarkerIDString(string(contents))
}

func (_this *Context) ValidateContentsMarkerIDString(contents string) {
	if len(contents) == 0 {
		panic(fmt.Errorf("Marker ID string cannot be empty"))
	}

	ch, _ := utf8.DecodeRuneInString(contents)
	if chars.RuneHasProperty(ch, chars.CharNeedsQuoteFirst) {
		panic(fmt.Errorf("ID [%s] first character is invalid", contents))
	}

	runeCount := 1
	for _, ch = range contents {
		runeCount++
		if chars.RuneHasProperty(ch, chars.CharNeedsQuote) {
			panic(fmt.Errorf("ID [%s] contains an invalid character", contents))
		}
	}
	if runeCount > maxMarkerIDRuneCount {
		panic(fmt.Errorf("Marker ID character length %d is greater than the maximum of %d", runeCount, maxMarkerIDRuneCount))
	}
}

func (_this *Context) ValidateNothing(_ []byte) {
	// Nothing to check
}
