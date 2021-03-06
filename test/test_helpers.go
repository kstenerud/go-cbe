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

// Test helper code.
package test

import (
	"fmt"
	"math"
	"math/big"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/kstenerud/go-concise-encoding/version"

	"github.com/kstenerud/go-concise-encoding/internal/arrays"

	"github.com/kstenerud/go-concise-encoding/debug"

	"github.com/kstenerud/go-concise-encoding/conversions"
	"github.com/kstenerud/go-concise-encoding/events"
	"github.com/kstenerud/go-concise-encoding/internal/common"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
	"github.com/kstenerud/go-equivalence"
)

// ----------------------------------------------------------------------------
// Pass through panics
// ----------------------------------------------------------------------------

// Causes the library to pass through all panics for the duration of the current
// function instead of converting them to error objects. This can be useful for
// tracking down the ultimate cause.
//
// Usage: defer test.PassThroughPanics(true)()
func PassThroughPanics(shouldPassThrough bool) func() {
	oldValue := debug.DebugOptions.PassThroughPanics
	debug.DebugOptions.PassThroughPanics = shouldPassThrough
	return func() {
		debug.DebugOptions.PassThroughPanics = oldValue
	}
}

// ----------------------------------------------------------------------------
// Constructors for common data types
// ----------------------------------------------------------------------------

func NewBigInt(str string, base int) *big.Int {
	bi := new(big.Int)
	_, success := bi.SetString(str, base)
	if !success {
		panic(fmt.Errorf("cannot convert %v to big.Int", str))
	}
	return bi
}

func NewBigFloat(str string, base int, significantDigits int) *big.Float {
	bits := uint(0)
	switch base {
	case 10:
		bits = uint(conversions.DecimalDigitsToBits(significantDigits))
	case 16:
		bits = uint(conversions.HexDigitsToBits(significantDigits))
	default:
		panic(fmt.Errorf("%v: Unhandled base", base))
	}
	f, _, err := big.ParseFloat(str, base, bits, big.ToNearestEven)
	if err != nil {
		panic(err)
	}
	return f
}

func NewDFloat(str string) compact_float.DFloat {
	df, err := compact_float.DFloatFromString(str)
	if err != nil {
		panic(err)
	}
	return df
}

func NewBDF(str string) *apd.Decimal {
	v, _, err := apd.NewFromString(str)
	if err != nil {
		panic(err)
	}
	return v
}

func NewRID(RIDString string) *url.URL {
	rid, err := url.Parse(RIDString)
	if err != nil {
		panic(fmt.Errorf("TEST CODE BUG: Bad URL (%v): %w", RIDString, err))
	}
	return rid
}

func NewDate(year, month, day int) compact_time.Time {
	t, err := compact_time.NewDate(year, month, day)
	if err != nil {
		panic(err)
	}
	return t
}

func NewTime(hour, minute, second, nanosecond int, areaLocation string) compact_time.Time {
	t, err := compact_time.NewTime(hour, minute, second, nanosecond, areaLocation)
	if err != nil {
		panic(err)
	}
	return t
}

func NewTimeLL(hour, minute, second, nanosecond, latitudeHundredths, longitudeHundredths int) compact_time.Time {
	t, err := compact_time.NewTimeLatLong(hour, minute, second, nanosecond, latitudeHundredths, longitudeHundredths)
	if err != nil {
		panic(err)
	}
	return t
}

func NewTS(year, month, day, hour, minute, second, nanosecond int, areaLocation string) compact_time.Time {
	t, err := compact_time.NewTimestamp(year, month, day, hour, minute, second, nanosecond, areaLocation)
	if err != nil {
		panic(err)
	}
	return t
}

func NewTSLL(year, month, day, hour, minute, second, nanosecond, latitudeHundredths, longitudeHundredths int) compact_time.Time {
	t, err := compact_time.NewTimestampLatLong(year, month, day, hour, minute, second, nanosecond, latitudeHundredths, longitudeHundredths)
	if err != nil {
		panic(err)
	}
	return t
}

func AsGoTime(t compact_time.Time) time.Time {
	gt, err := t.AsGoTime()
	if err != nil {
		panic(err)
	}
	return gt
}

func AsCompactTime(t time.Time) compact_time.Time {
	ct, err := compact_time.AsCompactTime(t)
	if err != nil {
		panic(err)
	}
	return ct
}

// ----------------------------------------------------------------------------
// Panics
// ----------------------------------------------------------------------------

func ReportPanic(function func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	function()
	return
}

func AssertNoPanic(t *testing.T, name interface{}, function func()) {
	if debug.DebugOptions.PassThroughPanics {
		function()
	} else {
		if err := ReportPanic(function); err != nil {
			t.Errorf("Unexpected error in %v: %v", name, err)
		}
	}
}

func AssertPanics(t *testing.T, name interface{}, function func()) bool {
	if err := ReportPanic(function); err == nil {
		t.Errorf("Expected an error in %v", name)
		return false
	}
	return true
}

// ----------------------------------------------------------------------------
// Generators
// ----------------------------------------------------------------------------

func GenerateString(charCount int, startIndex int) string {
	charRange := int('z' - 'a')
	var object strings.Builder
	for i := 0; i < charCount; i++ {
		ch := 'a' + (i+charCount+startIndex)%charRange
		object.WriteByte(byte(ch))
	}
	return object.String()
}

func GenerateBytes(length int, startIndex int) []byte {
	return []byte(GenerateString(length, startIndex))
}

func InvokeEvents(receiver events.DataEventReceiver, events ...*TEvent) {
	for _, event := range events {
		event.Invoke(receiver)
	}
}

func CloneBytes(bytes []byte) []byte {
	bytesCopy := make([]byte, len(bytes), len(bytes))
	copy(bytesCopy, bytes)
	return bytesCopy
}

// ----------------------------------------------------------------------------
// Events
// ----------------------------------------------------------------------------

var (
	EvBD     = BD()
	EvED     = ED()
	EvV      = V(version.ConciseEncodingVersion)
	EvPAD    = PAD(1)
	EvNA     = NA()
	EvB      = B(true)
	EvTT     = TT()
	EvFF     = FF()
	EvPI     = PI(1)
	EvNI     = NI(1)
	EvI      = I(0)
	EvBI     = BI(NewBigInt("1", 10))
	EvBINil  = BI(nil)
	EvF      = F(0.1)
	EvFNAN   = F(math.NaN())
	EvBF     = BF(NewBigFloat("0.1", 10, 1))
	EvBFNil  = BF(nil)
	EvDF     = DF(NewDFloat("0.1"))
	EvDFNAN  = DF(NewDFloat("nan"))
	EvBDF    = BDF(NewBDF("0.1"))
	EvBDFNil = BDF(nil)
	EvBDFNAN = BDF(NewBDF("nan"))
	EvNAN    = NAN()
	EvUUID   = UUID([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	EvGT     = GT(time.Date(2020, time.Month(1), 1, 1, 1, 1, 1, time.UTC))
	EvCT     = CT(NewDate(2020, 1, 1))
	EvL      = L()
	EvM      = M()
	EvMUP    = MUP()
	EvMETA   = META()
	EvCMT    = CMT()
	EvE      = E()
	EvMARK   = MARK()
	EvREF    = REF()
	EvCAT    = CAT()
	EvAC     = AC(1, false)
	EvAD     = AD([]byte{1})
	EvS      = S("a")
	EvSB     = SB()
	EvRID    = RID("http://z.com")
	EvRB     = RB()
	EvCUB    = CUB([]byte{1})
	EvCBB    = CBB()
	EvCUT    = CUT("a")
	EvCTB    = CTB()
	EvAB     = AB(1, []byte{1})
	EvABB    = ABB()
	EvAU8    = AU8([]uint8{1})
	EvAU8B   = AU8B()
	EvAU16   = AU16([]uint16{1})
	EvAU16B  = AU16B()
	EvAU32   = AU32([]uint32{1})
	EvAU32B  = AU32B()
	EvAU64   = AU64([]uint64{1})
	EvAU64B  = AU64B()
	EvAI8    = AI8([]int8{1})
	EvAI8B   = AI8B()
	EvAI16   = AI16([]int16{1})
	EvAI16B  = AI16B()
	EvAI32   = AI32([]int32{1})
	EvAI32B  = AI32B()
	EvAI64   = AI64([]int64{1})
	EvAI64B  = AI64B()
	EvAF16   = AF16([]byte{1, 2})
	EvAF16B  = AF16B()
	EvAF32   = AF32([]float32{1})
	EvAF32B  = AF32B()
	EvAF64   = AF64([]float64{1})
	EvAF64B  = AF64B()
	EvAUU    = AUU([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	EvAUUB   = AUUB()
)

var allEvents = []*TEvent{
	EvBD, EvED, EvV, EvPAD, EvNA, EvB, EvTT, EvFF, EvPI, EvNI, EvI, EvBI,
	EvBINil, EvF, EvFNAN, EvBF, EvBFNil, EvDF, EvDFNAN, EvBDF, EvBDFNil,
	EvBDFNAN, EvNAN, EvUUID, EvGT, EvCT, EvL, EvM, EvMUP, EvMETA,
	EvCMT, EvE, EvMARK, EvREF, EvCAT, EvAC, EvAD, EvS, EvSB, EvRID, EvRB,
	EvCUB, EvCBB, EvCUT, EvCTB, EvAB, EvABB, EvAU8, EvAU8B, EvAU16, EvAU16B,
	EvAU32, EvAU32B, EvAU64, EvAU64B, EvAI8, EvAI8B, EvAI16, EvAI16B, EvAI32,
	EvAI32B, EvAI64, EvAI64B, EvAF16, EvAF16B, EvAF32, EvAF32B, EvAF64,
	EvAF64B, EvAUU, EvAUUB,
}

var Completions = map[*TEvent][]*TEvent{
	EvL:     []*TEvent{EvE},
	EvM:     []*TEvent{EvE},
	EvMETA:  []*TEvent{EvE, S("a")},
	EvCMT:   []*TEvent{EvE, S("a")},
	EvMUP:   []*TEvent{S("a"), EvE, EvE},
	EvMARK:  []*TEvent{S("a"), S("m")},
	EvREF:   []*TEvent{S("a")},
	EvCAT:   []*TEvent{S("a")},
	EvPAD:   []*TEvent{S("a")},
	EvSB:    []*TEvent{AC(0, false)},
	EvRB:    []*TEvent{AC(0, false)},
	EvCBB:   []*TEvent{AC(0, false)},
	EvCTB:   []*TEvent{AC(0, false)},
	EvABB:   []*TEvent{AC(0, false)},
	EvAU8B:  []*TEvent{AC(0, false)},
	EvAU16B: []*TEvent{AC(0, false)},
	EvAU32B: []*TEvent{AC(0, false)},
	EvAU64B: []*TEvent{AC(0, false)},
	EvAI8B:  []*TEvent{AC(0, false)},
	EvAI16B: []*TEvent{AC(0, false)},
	EvAI32B: []*TEvent{AC(0, false)},
	EvAI64B: []*TEvent{AC(0, false)},
	EvAF16B: []*TEvent{AC(0, false)},
	EvAF32B: []*TEvent{AC(0, false)},
	EvAF64B: []*TEvent{AC(0, false)},
	EvAUUB:  []*TEvent{AC(0, false)},
}

func isEffectivelyNA(event *TEvent) bool {
	return event.Type == TEventNA ||
		event == EvBINil ||
		event == EvBFNil ||
		event == EvBDFNil
}

func FilterAllEvents(events []*TEvent, filter func(*TEvent) []*TEvent) []*TEvent {
	filtered := []*TEvent{}
	for _, event := range events {
		filtered = append(filtered, filter(event)...)
	}
	return filtered
}

func FilterCTE(event *TEvent) []*TEvent {
	switch event.Type {
	case TEventPadding:
		return []*TEvent{}
	case TEventArrayBooleanBegin, TEventArrayFloat16Begin,
		TEventArrayFloat32Begin, TEventArrayFloat64Begin,
		TEventArrayInt8Begin, TEventArrayInt16Begin, TEventArrayInt32Begin,
		TEventArrayInt64Begin, TEventArrayUint8Begin, TEventArrayUint16Begin,
		TEventArrayUint32Begin, TEventArrayUint64Begin, TEventArrayUUIDBegin,
		TEventCustomBinaryBegin, TEventCustomTextBegin, TEventResourceIDBegin,
		TEventStringBegin:
		return []*TEvent{S("x")}
	case TEventArrayChunk, TEventArrayData:
		return []*TEvent{}
	default:
		return []*TEvent{event}
	}
}

func FilterContainer(event *TEvent) []*TEvent {
	switch event.Type {
	case TEventEnd:
		return []*TEvent{}
	default:
		return []*TEvent{event}
	}
}

func FilterKey(event *TEvent) []*TEvent {
	switch event.Type {
	case TEventEnd, TEventReference:
		return []*TEvent{}
	default:
		return []*TEvent{event}
	}
}

func FilterEventsSwitchToRIDRefs(events []*TEvent) []*TEvent {
	filtered := []*TEvent{}
	var lastEvent *TEvent = EvBD
	for _, event := range events {
		if lastEvent.Type == TEventReference && (event.Type == TEventInt || event.Type == TEventPInt || event.Type == TEventBigInt || event.Type == TEventString) {
			filtered = append(filtered, EvRID)
		} else {
			filtered = append(filtered, event)
		}
		lastEvent = event
	}

	return filtered
}

func FilterEventsForCTE(events []*TEvent) []*TEvent {
	return FilterAllEvents(events, FilterCTE)
}

func FilterEventsForContainer(events []*TEvent) []*TEvent {
	return FilterAllEvents(events, FilterContainer)
}

func FilterEventsForKey(events []*TEvent) []*TEvent {
	return FilterAllEvents(events, FilterKey)
}

func ComplementaryEvents(events []*TEvent) []*TEvent {
	complementary := make([]*TEvent, 0, len(allEvents)/2)
	for _, event := range allEvents {
		for _, compareEvent := range events {
			if event == compareEvent {
				goto Skip
			}
		}
		complementary = append(complementary, event)
	Skip:
	}
	return complementary
}

var (
	MarkerIDTypes        = []*TEvent{EvPAD, EvPI, EvI, EvBI, EvS, EvSB}
	InvalidMarkerIDTypes = ComplementaryEvents(MarkerIDTypes)

	ReferenceIDTypes        = []*TEvent{EvPAD, EvPI, EvI, EvBI, EvS, EvSB, EvRID, EvRB}
	InvalidReferenceIDTypes = ComplementaryEvents(ReferenceIDTypes)

	KeyableReferenceIDTypes        = []*TEvent{EvPAD, EvPI, EvI, EvBI, EvS, EvSB}
	InvalidKeyableReferenceIDTypes = ComplementaryEvents(KeyableReferenceIDTypes)

	ArrayBeginTypes = []*TEvent{
		EvSB, EvRB, EvCBB, EvCTB, EvABB, EvAU8B, EvAU16B, EvAU32B, EvAU64B,
		EvAI8B, EvAI16B, EvAI32B, EvAI64B, EvAF16B, EvAF32B, EvAF64B, EvAUUB,
	}

	ValidTLOValues   = ComplementaryEvents(InvalidTLOValues)
	InvalidTLOValues = []*TEvent{EvBD, EvED, EvV, EvE, EvAC, EvAD, EvCAT}

	ValidMapKeys = []*TEvent{
		EvPAD, EvB, EvTT, EvFF, EvPI, EvNI, EvI, EvBI, EvF, EvBF, EvDF, EvBDF,
		EvUUID, EvGT, EvCT, EvMARK, EvS, EvSB, EvRID, EvRB, EvREF, EvMETA, EvCMT, EvE,
	}
	InvalidMapKeys = ComplementaryEvents(ValidMapKeys)

	ValidMapValues   = ComplementaryEvents(InvalidMapValues)
	InvalidMapValues = []*TEvent{EvBD, EvED, EvV, EvE, EvAC, EvAD, EvCAT}

	ValidListValues   = ComplementaryEvents(InvalidListValues)
	InvalidListValues = []*TEvent{EvBD, EvED, EvV, EvAC, EvAD, EvCAT}

	ValidCommentValues   = []*TEvent{EvCMT, EvE, EvS, EvSB, EvPAD}
	InvalidCommentValues = ComplementaryEvents(ValidCommentValues)

	ValidMarkupNames = []*TEvent{
		EvPAD, EvB, EvTT, EvFF, EvPI, EvNI, EvI, EvBI, EvF, EvBF, EvDF, EvBDF,
		EvUUID, EvGT, EvCT, EvMARK, EvREF, EvS, EvSB, EvRID, EvRB,
	}
	InvalidMarkupNames = ComplementaryEvents(ValidMarkupNames)

	ValidMarkupContents   = []*TEvent{EvPAD, EvS, EvSB, EvMUP, EvCMT, EvE}
	InvalidMarkupContents = ComplementaryEvents(ValidMarkupContents)

	ValidAfterArrayBegin   = []*TEvent{EvAC}
	InvalidAfterArrayBegin = ComplementaryEvents(ValidAfterArrayBegin)

	ValidAfterArrayChunk   = []*TEvent{EvAD}
	InvalidAfterArrayChunk = ComplementaryEvents(ValidAfterArrayChunk)

	ValidMarkerIDs   = []*TEvent{EvPAD, EvS, EvSB, EvPI, EvI, EvBI}
	InvalidMarkerIDs = ComplementaryEvents(ValidMarkerIDs)

	ValidMarkerValues   = ComplementaryEvents(InvalidMarkerValues)
	InvalidMarkerValues = []*TEvent{EvBD, EvED, EvV, EvE, EvAC, EvAD, EvMETA, EvCMT, EvMARK, EvCAT}

	ValidReferenceIDs   = []*TEvent{EvPAD, EvS, EvSB, EvPI, EvI, EvBI, EvRID, EvRB}
	InvalidReferenceIDs = ComplementaryEvents(ValidReferenceIDs)
)

// ----------------------------------------------------------------------------
// Event types and pretty-print names
// ----------------------------------------------------------------------------

type TEventType int

const (
	TEventBeginDocument TEventType = iota
	TEventVersion
	TEventPadding
	TEventNA
	TEventBool
	TEventTrue
	TEventFalse
	TEventPInt
	TEventNInt
	TEventInt
	TEventBigInt
	TEventFloat
	TEventBigFloat
	TEventDecimalFloat
	TEventBigDecimalFloat
	TEventNan
	TEventSNan
	TEventUUID
	TEventTime
	TEventCompactTime
	TEventString
	TEventResourceID
	TEventCustomBinary
	TEventCustomText
	TEventArrayBoolean
	TEventArrayInt8
	TEventArrayInt16
	TEventArrayInt32
	TEventArrayInt64
	TEventArrayUint8
	TEventArrayUint16
	TEventArrayUint32
	TEventArrayUint64
	TEventArrayFloat16
	TEventArrayFloat32
	TEventArrayFloat64
	TEventArrayUUID
	TEventStringBegin
	TEventResourceIDBegin
	TEventCustomBinaryBegin
	TEventCustomTextBegin
	TEventArrayBooleanBegin
	TEventArrayInt8Begin
	TEventArrayInt16Begin
	TEventArrayInt32Begin
	TEventArrayInt64Begin
	TEventArrayUint8Begin
	TEventArrayUint16Begin
	TEventArrayUint32Begin
	TEventArrayUint64Begin
	TEventArrayFloat16Begin
	TEventArrayFloat32Begin
	TEventArrayFloat64Begin
	TEventArrayUUIDBegin
	TEventArrayChunk
	TEventArrayData
	TEventList
	TEventMap
	TEventMarkup
	TEventMetadata
	TEventComment
	TEventEnd
	TEventMarker
	TEventReference
	TEventConcatenate
	TEventConstant
	TEventEndDocument
)

var TEventNames = []string{
	TEventBeginDocument:     "BD",
	TEventVersion:           "V",
	TEventPadding:           "PAD",
	TEventNA:                "NA",
	TEventBool:              "B",
	TEventTrue:              "TT",
	TEventFalse:             "FF",
	TEventPInt:              "PI",
	TEventNInt:              "NI",
	TEventInt:               "I",
	TEventBigInt:            "BI",
	TEventFloat:             "F",
	TEventBigFloat:          "BF",
	TEventDecimalFloat:      "DF",
	TEventBigDecimalFloat:   "BDF",
	TEventNan:               "NAN",
	TEventSNan:              "SNAN",
	TEventUUID:              "UUID",
	TEventTime:              "GT",
	TEventCompactTime:       "CT",
	TEventString:            "S",
	TEventResourceID:        "RID",
	TEventCustomBinary:      "CUB",
	TEventCustomText:        "CUT",
	TEventArrayBoolean:      "AB",
	TEventArrayInt8:         "AI8",
	TEventArrayInt16:        "AI16",
	TEventArrayInt32:        "AI32",
	TEventArrayInt64:        "AI64",
	TEventArrayUint8:        "AU8",
	TEventArrayUint16:       "AU16",
	TEventArrayUint32:       "AU32",
	TEventArrayUint64:       "AU64",
	TEventArrayFloat16:      "AF16",
	TEventArrayFloat32:      "AF32",
	TEventArrayFloat64:      "AF64",
	TEventArrayUUID:         "AUU",
	TEventStringBegin:       "SB",
	TEventResourceIDBegin:   "RB",
	TEventCustomBinaryBegin: "CBB",
	TEventCustomTextBegin:   "CTB",
	TEventArrayBooleanBegin: "ABB",
	TEventArrayInt8Begin:    "AI8B",
	TEventArrayInt16Begin:   "AI16B",
	TEventArrayInt32Begin:   "AI32B",
	TEventArrayInt64Begin:   "AI64B",
	TEventArrayUint8Begin:   "AU8B",
	TEventArrayUint16Begin:  "AU16B",
	TEventArrayUint32Begin:  "AU32B",
	TEventArrayUint64Begin:  "AU64B",
	TEventArrayFloat16Begin: "AF16B",
	TEventArrayFloat32Begin: "AF32B",
	TEventArrayFloat64Begin: "AF64B",
	TEventArrayUUIDBegin:    "AUUB",
	TEventArrayChunk:        "AC",
	TEventArrayData:         "AD",
	TEventList:              "L",
	TEventMap:               "M",
	TEventMarkup:            "MUP",
	TEventMetadata:          "META",
	TEventComment:           "CMT",
	TEventEnd:               "E",
	TEventMarker:            "MARK",
	TEventReference:         "REF",
	TEventConcatenate:       "CAT",
	TEventConstant:          "CONST",
	TEventEndDocument:       "ED",
}

func (_this TEventType) String() string {
	return TEventNames[_this]
}

func (_this TEventType) IsBoolean() bool {
	switch _this {
	case TEventTrue, TEventFalse, TEventBool:
		return true
	default:
		return false
	}
}

func (_this TEventType) IsNumeric() bool {
	switch _this {
	case TEventPInt, TEventNInt, TEventInt, TEventBigInt, TEventFloat,
		TEventBigFloat, TEventDecimalFloat, TEventBigDecimalFloat, TEventNan,
		TEventSNan:
		return true
	default:
		return false
	}
}

// ----------------------------------------------------------------------------
// Stored events
// ----------------------------------------------------------------------------

type TEvent struct {
	Type TEventType
	V1   interface{}
	V2   interface{}
}

func newTEvent(eventType TEventType, v1 interface{}, v2 interface{}) *TEvent {
	return &TEvent{
		Type: eventType,
		V1:   v1,
		V2:   v2,
	}
}

func (_this *TEvent) String() string {
	str := _this.Type.String()
	if _this.V1 != nil {
		// TODO: Stringify bytes as hex
		if _this.V2 != nil {
			return fmt.Sprintf("%v(%v,%v)", str, _this.V1, _this.V2)
		}
		return fmt.Sprintf("%v(%v)", str, _this.V1)
	}
	return str
}

func (_this *TEvent) isNan() bool {
	switch _this.Type {
	case TEventNan:
		return true
	case TEventFloat:
		f64 := _this.V1.(float64)
		return math.IsNaN(f64) && !common.IsSignalingNan(f64)
	case TEventDecimalFloat:
		return _this.V1.(compact_float.DFloat).IsNan()
	case TEventBigDecimalFloat:
		return _this.V1.(*apd.Decimal).Form == apd.NaN
	default:
		return false
	}
}

func (_this *TEvent) isSignalingNan() bool {
	switch _this.Type {
	case TEventSNan:
		return true
	case TEventFloat:
		f64 := _this.V1.(float64)
		return math.IsNaN(f64) && common.IsSignalingNan(f64)
	case TEventDecimalFloat:
		return _this.V1.(compact_float.DFloat).IsSignalingNan()
	case TEventBigDecimalFloat:
		return _this.V1.(*apd.Decimal).Form == apd.NaNSignaling
	default:
		return false
	}
}

func isZeroTZ(tz string) bool {
	switch tz {
	case "", "Z", "Zero", "Etc/GMT", "Etc/GMT+0", "Etc/GMT-0", "Etc/GMT0", "Etc/Greenwich",
		"Etc/UCT", "Etc/Universal", "Etc/UTC", "Etc/Zulu", "Factory", "GMT",
		"GMT+0", "GMT-0", "GMT0", "Greenwich", "UCT", "Universal", "UTC", "Zulu":
		return true
	default:
		return false
	}
}

func (_this *TEvent) isEquivalentTo(that *TEvent) bool {
	if equivalence.IsEquivalent(_this, that) {
		return true
	}

	if _this.Type.IsNumeric() && that.Type.IsNumeric() {
		if _this.isNan() && that.isNan() {
			return true
		}
		if _this.isSignalingNan() && that.isSignalingNan() {
			return true
		}

		var a string
		if _this.Type == TEventNInt {
			a = fmt.Sprintf("-%v", _this.V1)
		} else {
			a = fmt.Sprintf("%v", _this.V1)
		}

		var b string
		if that.Type == TEventNInt {
			b = fmt.Sprintf("-%v", that.V1)
		} else {
			b = fmt.Sprintf("%v", that.V1)
		}

		return a == b
	}

	if _this.Type.IsBoolean() && that.Type.IsBoolean() {
		a := fmt.Sprintf("%v", _this.V1)
		b := fmt.Sprintf("%v", that.V1)
		return a == b ||
			a == "true" && that.Type == TEventTrue ||
			b == "true" && _this.Type == TEventTrue ||
			a == "false" && that.Type == TEventFalse ||
			b == "false" && _this.Type == TEventFalse
	}

	if _this.Type == TEventTime || _this.Type == TEventCompactTime {
		var err error
		var a compact_time.Time
		var b compact_time.Time

		switch _this.Type {
		case TEventCompactTime:
			a = _this.V1.(compact_time.Time)
		default:
			a, err = compact_time.AsCompactTime(_this.V1.(time.Time))
			if err != nil {
				panic(err)
			}
		}

		switch that.Type {
		case TEventCompactTime:
			b = that.V1.(compact_time.Time)
		case TEventTime:
			b, err = compact_time.AsCompactTime(that.V1.(time.Time))
			if err != nil {
				panic(err)
			}
		default:
			return false
		}

		if isZeroTZ(a.LongAreaLocation) && isZeroTZ(b.LongAreaLocation) {
			return a.Year == b.Year &&
				a.Month == b.Month &&
				a.Day == b.Day &&
				a.Hour == b.Hour &&
				a.Minute == b.Minute &&
				a.Second == b.Second &&
				a.Nanosecond == b.Nanosecond &&
				a.TimeType == b.TimeType
		}
		return a == b
	}

	if isEffectivelyNA(_this) && isEffectivelyNA(that) {
		return true
	}

	return false
}

func AreAllEventsEqual(a []*TEvent, b []*TEvent) bool {
	if len(a) != len(b) {
		return false
	}

	for i, ev := range a {
		if !ev.isEquivalentTo(b[i]) {
			return false
		}
	}
	return true
}

// Invoking a stored event generates the appropriate data event message to
// the receiver.
func (_this *TEvent) Invoke(receiver events.DataEventReceiver) {
	switch _this.Type {
	case TEventBeginDocument:
		receiver.OnBeginDocument()
	case TEventVersion:
		receiver.OnVersion(_this.V1.(uint64))
	case TEventPadding:
		receiver.OnPadding(_this.V1.(int))
	case TEventNA:
		receiver.OnNA()
	case TEventBool:
		receiver.OnBool(_this.V1.(bool))
	case TEventTrue:
		receiver.OnTrue()
	case TEventFalse:
		receiver.OnFalse()
	case TEventPInt:
		receiver.OnPositiveInt(_this.V1.(uint64))
	case TEventNInt:
		receiver.OnNegativeInt(_this.V1.(uint64))
	case TEventInt:
		receiver.OnInt(_this.V1.(int64))
	case TEventBigInt:
		receiver.OnBigInt(_this.V1.(*big.Int))
	case TEventFloat:
		receiver.OnFloat(_this.V1.(float64))
	case TEventBigFloat:
		receiver.OnBigFloat(_this.V1.(*big.Float))
	case TEventDecimalFloat:
		receiver.OnDecimalFloat(_this.V1.(compact_float.DFloat))
	case TEventBigDecimalFloat:
		receiver.OnBigDecimalFloat(_this.V1.(*apd.Decimal))
	case TEventNan:
		receiver.OnNan(false)
	case TEventSNan:
		receiver.OnNan(true)
	case TEventUUID:
		receiver.OnUUID(_this.V1.([]byte))
	case TEventTime:
		receiver.OnTime(_this.V1.(time.Time))
	case TEventCompactTime:
		receiver.OnCompactTime(_this.V1.(compact_time.Time))
	case TEventString:
		receiver.OnStringlikeArray(events.ArrayTypeString, _this.V1.(string))
	case TEventResourceID:
		receiver.OnStringlikeArray(events.ArrayTypeResourceID, _this.V1.(string))
	case TEventCustomBinary:
		bytes := _this.V1.([]byte)
		receiver.OnArray(events.ArrayTypeCustomBinary, uint64(len(bytes)), bytes)
	case TEventCustomText:
		receiver.OnStringlikeArray(events.ArrayTypeCustomText, _this.V1.(string))
	case TEventArrayBoolean:
		bitCount := _this.V1.(uint64)
		bytes := _this.V2.([]byte)
		receiver.OnArray(events.ArrayTypeBoolean, bitCount, bytes)
	case TEventArrayInt8:
		bytes := arrays.Int8SliceAsBytes(_this.V1.([]int8))
		receiver.OnArray(events.ArrayTypeInt8, uint64(len(bytes)), bytes)
	case TEventArrayInt16:
		bytes := arrays.Int16SliceAsBytes(_this.V1.([]int16))
		receiver.OnArray(events.ArrayTypeInt16, uint64(len(bytes)/2), bytes)
	case TEventArrayInt32:
		bytes := arrays.Int32SliceAsBytes(_this.V1.([]int32))
		receiver.OnArray(events.ArrayTypeInt32, uint64(len(bytes)/4), bytes)
	case TEventArrayInt64:
		bytes := arrays.Int64SliceAsBytes(_this.V1.([]int64))
		receiver.OnArray(events.ArrayTypeInt64, uint64(len(bytes)/8), bytes)
	case TEventArrayUint8:
		bytes := _this.V1.([]byte)
		receiver.OnArray(events.ArrayTypeUint8, uint64(len(bytes)), bytes)
	case TEventArrayUint16:
		bytes := arrays.Uint16SliceAsBytes(_this.V1.([]uint16))
		receiver.OnArray(events.ArrayTypeUint16, uint64(len(bytes)/2), bytes)
	case TEventArrayUint32:
		bytes := arrays.Uint32SliceAsBytes(_this.V1.([]uint32))
		receiver.OnArray(events.ArrayTypeUint32, uint64(len(bytes)/4), bytes)
	case TEventArrayUint64:
		bytes := arrays.Uint64SliceAsBytes(_this.V1.([]uint64))
		receiver.OnArray(events.ArrayTypeUint64, uint64(len(bytes)/8), bytes)
	case TEventArrayFloat16:
		// TODO: How to handle float16 in go code?
		bytes := _this.V1.([]byte)
		receiver.OnArray(events.ArrayTypeFloat16, uint64(len(bytes)/2), bytes)
	case TEventArrayFloat32:
		bytes := arrays.Float32SliceAsBytes(_this.V1.([]float32))
		receiver.OnArray(events.ArrayTypeFloat32, uint64(len(bytes)/4), bytes)
	case TEventArrayFloat64:
		bytes := arrays.Float64SliceAsBytes(_this.V1.([]float64))
		receiver.OnArray(events.ArrayTypeFloat64, uint64(len(bytes)/8), bytes)
	case TEventArrayUUID:
		bytes := _this.V1.([]byte)
		receiver.OnArray(events.ArrayTypeUUID, uint64(len(bytes)/16), bytes)
	case TEventStringBegin:
		receiver.OnArrayBegin(events.ArrayTypeString)
	case TEventResourceIDBegin:
		receiver.OnArrayBegin(events.ArrayTypeResourceID)
	case TEventCustomBinaryBegin:
		receiver.OnArrayBegin(events.ArrayTypeCustomBinary)
	case TEventCustomTextBegin:
		receiver.OnArrayBegin(events.ArrayTypeCustomText)
	case TEventArrayBooleanBegin:
		receiver.OnArrayBegin(events.ArrayTypeBoolean)
	case TEventArrayInt8Begin:
		receiver.OnArrayBegin(events.ArrayTypeInt8)
	case TEventArrayInt16Begin:
		receiver.OnArrayBegin(events.ArrayTypeInt16)
	case TEventArrayInt32Begin:
		receiver.OnArrayBegin(events.ArrayTypeInt32)
	case TEventArrayInt64Begin:
		receiver.OnArrayBegin(events.ArrayTypeInt64)
	case TEventArrayUint8Begin:
		receiver.OnArrayBegin(events.ArrayTypeUint8)
	case TEventArrayUint16Begin:
		receiver.OnArrayBegin(events.ArrayTypeUint16)
	case TEventArrayUint32Begin:
		receiver.OnArrayBegin(events.ArrayTypeUint32)
	case TEventArrayUint64Begin:
		receiver.OnArrayBegin(events.ArrayTypeUint64)
	case TEventArrayFloat16Begin:
		receiver.OnArrayBegin(events.ArrayTypeFloat16)
	case TEventArrayFloat32Begin:
		receiver.OnArrayBegin(events.ArrayTypeFloat32)
	case TEventArrayFloat64Begin:
		receiver.OnArrayBegin(events.ArrayTypeFloat64)
	case TEventArrayUUIDBegin:
		receiver.OnArrayBegin(events.ArrayTypeUUID)
	case TEventArrayChunk:
		receiver.OnArrayChunk(_this.V1.(uint64), _this.V2.(bool))
	case TEventArrayData:
		receiver.OnArrayData(_this.V1.([]byte))
	case TEventList:
		receiver.OnList()
	case TEventMap:
		receiver.OnMap()
	case TEventMarkup:
		receiver.OnMarkup()
	case TEventMetadata:
		receiver.OnMetadata()
	case TEventComment:
		receiver.OnComment()
	case TEventEnd:
		receiver.OnEnd()
	case TEventMarker:
		receiver.OnMarker()
	case TEventReference:
		receiver.OnReference()
	case TEventConcatenate:
		receiver.OnConcatenate()
	case TEventConstant:
		receiver.OnConstant([]byte(_this.V1.(string)), _this.V2.(bool))
	case TEventEndDocument:
		receiver.OnEndDocument()
	default:
		panic(fmt.Errorf("%v: Unhandled event type", _this.Type))
	}
}

func EventOrNil(eventType TEventType, value interface{}) *TEvent {
	if value == nil {
		eventType = TEventNA
	}
	return newTEvent(eventType, value, nil)
}

// ----------------------------------------------------------------------------
// Stored event convenience constructors
// ----------------------------------------------------------------------------

func TT() *TEvent                       { return newTEvent(TEventTrue, nil, nil) }
func FF() *TEvent                       { return newTEvent(TEventFalse, nil, nil) }
func I(v int64) *TEvent                 { return newTEvent(TEventInt, v, nil) }
func F(v float64) *TEvent               { return newTEvent(TEventFloat, v, nil) }
func BF(v *big.Float) *TEvent           { return EventOrNil(TEventBigFloat, v) }
func DF(v compact_float.DFloat) *TEvent { return newTEvent(TEventDecimalFloat, v, nil) }
func BDF(v *apd.Decimal) *TEvent        { return EventOrNil(TEventBigDecimalFloat, v) }
func V(v uint64) *TEvent                { return newTEvent(TEventVersion, v, nil) }
func NA() *TEvent                       { return newTEvent(TEventNA, nil, nil) }
func PAD(v int) *TEvent                 { return newTEvent(TEventPadding, v, nil) }
func B(v bool) *TEvent                  { return newTEvent(TEventBool, v, nil) }
func PI(v uint64) *TEvent               { return newTEvent(TEventPInt, v, nil) }
func NI(v uint64) *TEvent               { return newTEvent(TEventNInt, v, nil) }
func BI(v *big.Int) *TEvent             { return EventOrNil(TEventBigInt, v) }
func NAN() *TEvent                      { return newTEvent(TEventNan, nil, nil) }
func SNAN() *TEvent                     { return newTEvent(TEventSNan, nil, nil) }
func UUID(v []byte) *TEvent             { return newTEvent(TEventUUID, v, nil) }
func GT(v time.Time) *TEvent            { return newTEvent(TEventTime, v, nil) }
func CT(v compact_time.Time) *TEvent    { return EventOrNil(TEventCompactTime, v) }
func S(v string) *TEvent                { return newTEvent(TEventString, v, nil) }
func RID(v string) *TEvent              { return newTEvent(TEventResourceID, v, nil) }
func CUB(v []byte) *TEvent              { return newTEvent(TEventCustomBinary, v, nil) }
func CUT(v string) *TEvent              { return newTEvent(TEventCustomText, v, nil) }
func AB(l uint64, v []byte) *TEvent     { return newTEvent(TEventArrayBoolean, l, v) }
func AI8(v []int8) *TEvent              { return newTEvent(TEventArrayInt8, v, nil) }
func AI16(v []int16) *TEvent            { return newTEvent(TEventArrayInt16, v, nil) }
func AI32(v []int32) *TEvent            { return newTEvent(TEventArrayInt32, v, nil) }
func AI64(v []int64) *TEvent            { return newTEvent(TEventArrayInt64, v, nil) }
func AU8(v []byte) *TEvent              { return newTEvent(TEventArrayUint8, v, nil) }
func AU16(v []uint16) *TEvent           { return newTEvent(TEventArrayUint16, v, nil) }
func AU32(v []uint32) *TEvent           { return newTEvent(TEventArrayUint32, v, nil) }
func AU64(v []uint64) *TEvent           { return newTEvent(TEventArrayUint64, v, nil) }
func AF16(v []byte) *TEvent             { return newTEvent(TEventArrayFloat16, v, nil) }
func AF32(v []float32) *TEvent          { return newTEvent(TEventArrayFloat32, v, nil) }
func AF64(v []float64) *TEvent          { return newTEvent(TEventArrayFloat64, v, nil) }
func AUU(v []byte) *TEvent              { return newTEvent(TEventArrayUUID, v, nil) }
func SB() *TEvent                       { return newTEvent(TEventStringBegin, nil, nil) }
func RB() *TEvent                       { return newTEvent(TEventResourceIDBegin, nil, nil) }
func CBB() *TEvent                      { return newTEvent(TEventCustomBinaryBegin, nil, nil) }
func CTB() *TEvent                      { return newTEvent(TEventCustomTextBegin, nil, nil) }
func ABB() *TEvent                      { return newTEvent(TEventArrayBooleanBegin, nil, nil) }
func AI8B() *TEvent                     { return newTEvent(TEventArrayInt8Begin, nil, nil) }
func AI16B() *TEvent                    { return newTEvent(TEventArrayInt16Begin, nil, nil) }
func AI32B() *TEvent                    { return newTEvent(TEventArrayInt32Begin, nil, nil) }
func AI64B() *TEvent                    { return newTEvent(TEventArrayInt64Begin, nil, nil) }
func AU8B() *TEvent                     { return newTEvent(TEventArrayUint8Begin, nil, nil) }
func AU16B() *TEvent                    { return newTEvent(TEventArrayUint16Begin, nil, nil) }
func AU32B() *TEvent                    { return newTEvent(TEventArrayUint32Begin, nil, nil) }
func AU64B() *TEvent                    { return newTEvent(TEventArrayUint64Begin, nil, nil) }
func AF16B() *TEvent                    { return newTEvent(TEventArrayFloat16Begin, nil, nil) }
func AF32B() *TEvent                    { return newTEvent(TEventArrayFloat32Begin, nil, nil) }
func AF64B() *TEvent                    { return newTEvent(TEventArrayFloat64Begin, nil, nil) }
func AUUB() *TEvent                     { return newTEvent(TEventArrayUUIDBegin, nil, nil) }
func AC(l uint64, more bool) *TEvent    { return newTEvent(TEventArrayChunk, l, more) }
func AD(v []byte) *TEvent               { return newTEvent(TEventArrayData, v, nil) }
func L() *TEvent                        { return newTEvent(TEventList, nil, nil) }
func M() *TEvent                        { return newTEvent(TEventMap, nil, nil) }
func MUP() *TEvent                      { return newTEvent(TEventMarkup, nil, nil) }
func META() *TEvent                     { return newTEvent(TEventMetadata, nil, nil) }
func CMT() *TEvent                      { return newTEvent(TEventComment, nil, nil) }
func E() *TEvent                        { return newTEvent(TEventEnd, nil, nil) }
func MARK() *TEvent                     { return newTEvent(TEventMarker, nil, nil) }
func REF() *TEvent                      { return newTEvent(TEventReference, nil, nil) }
func CAT() *TEvent                      { return newTEvent(TEventConcatenate, nil, nil) }
func CONST(n string, e bool) *TEvent    { return newTEvent(TEventConstant, n, e) }
func BD() *TEvent                       { return newTEvent(TEventBeginDocument, nil, nil) }
func ED() *TEvent                       { return newTEvent(TEventEndDocument, nil, nil) }

// Converts a go value into a stored event
func EventForValue(value interface{}) *TEvent {
	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return NA()
	}
	switch rv.Kind() {
	case reflect.Bool:
		return B(rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return I(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return PI(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return F(rv.Float())
	case reflect.String:
		return S(rv.String())
	case reflect.Slice:
		switch rv.Type().Elem().Kind() {
		case reflect.Uint8:
			return AU8(rv.Bytes())
		}
	case reflect.Ptr:
		if rv.IsNil() {
			return NA()
		}
		switch rv.Type() {
		case common.TypePBigDecimalFloat:
			return BDF(rv.Interface().(*apd.Decimal))
		case common.TypePBigFloat:
			return BF(rv.Interface().(*big.Float))
		case common.TypePBigInt:
			return BI(rv.Interface().(*big.Int))
		case common.TypePURL:
			return RID(rv.Interface().(*url.URL).String())
		}
		return EventForValue(rv.Elem().Interface())
	case reflect.Struct:
		switch rv.Type() {
		case common.TypeBigDecimalFloat:
			v := rv.Interface().(apd.Decimal)
			return BDF(&v)
		case common.TypeBigFloat:
			v := rv.Interface().(big.Float)
			return BF(&v)
		case common.TypeBigInt:
			v := rv.Interface().(big.Int)
			return BI(&v)
		case common.TypeCompactTime:
			v := rv.Interface().(compact_time.Time)
			return CT(v)
		case common.TypeDFloat:
			v := rv.Interface().(compact_float.DFloat)
			return DF(v)
		case common.TypeTime:
			v := rv.Interface().(time.Time)
			return GT(v)
		case common.TypeURL:
			v := rv.Interface().(url.URL)
			return RID(v.String())
		}
	}
	panic(fmt.Errorf("TEST CODE BUG: Unhandled kind: %v", rv.Kind()))
}

// ----------------------------------------------------------------------------
// Event printer
// ----------------------------------------------------------------------------

type TEventPrinter struct {
	Next  events.DataEventReceiver
	Print func(event *TEvent)
}

// Create an event receiver that prints the event to stdout.
func NewStdoutTEventPrinter(next events.DataEventReceiver) *TEventPrinter {
	return &TEventPrinter{
		Next: next,
		Print: func(event *TEvent) {
			fmt.Printf("EVENT %v\n", event)
		},
	}
}

func (h *TEventPrinter) OnBeginDocument() {
	h.Print(BD())
	h.Next.OnBeginDocument()
}
func (h *TEventPrinter) OnVersion(version uint64) {
	h.Print(V(version))
	h.Next.OnVersion(version)
}
func (h *TEventPrinter) OnPadding(count int) {
	h.Print(PAD(count))
	h.Next.OnPadding(count)
}
func (h *TEventPrinter) OnNA() {
	h.Print(NA())
	h.Next.OnNA()
}
func (h *TEventPrinter) OnBool(value bool) {
	h.Print(B(value))
	h.Next.OnBool(value)
}
func (h *TEventPrinter) OnTrue() {
	h.Print(TT())
	h.Next.OnTrue()
}
func (h *TEventPrinter) OnFalse() {
	h.Print(FF())
	h.Next.OnFalse()
}
func (h *TEventPrinter) OnPositiveInt(value uint64) {
	h.Print(PI(value))
	h.Next.OnPositiveInt(value)
}
func (h *TEventPrinter) OnNegativeInt(value uint64) {
	h.Print(NI(value))
	h.Next.OnNegativeInt(value)
}
func (h *TEventPrinter) OnInt(value int64) {
	h.Print(I(value))
	h.Next.OnInt(value)
}
func (h *TEventPrinter) OnBigInt(value *big.Int) {
	h.Print(BI(value))
	h.Next.OnBigInt(value)
}
func (h *TEventPrinter) OnFloat(value float64) {
	if math.IsNaN(value) {
		if common.IsSignalingNan(value) {
			h.Print(SNAN())
		} else {
			h.Print(NAN())
		}
	} else {
		h.Print(F(value))
	}
	h.Next.OnFloat(value)
}
func (h *TEventPrinter) OnBigFloat(value *big.Float) {
	h.Print(newTEvent(TEventBigFloat, value, nil))
	h.Next.OnBigFloat(value)
}
func (h *TEventPrinter) OnDecimalFloat(value compact_float.DFloat) {
	if value.IsNan() {
		if value.IsSignalingNan() {
			h.Print(SNAN())
		} else {
			h.Print(NAN())
		}
	} else {
		h.Print(DF(value))
	}
	h.Next.OnDecimalFloat(value)
}
func (h *TEventPrinter) OnBigDecimalFloat(value *apd.Decimal) {
	switch value.Form {
	case apd.NaN:
		h.Print(NAN())
	case apd.NaNSignaling:
		h.Print(SNAN())
	default:
		h.Print(BDF(value))
	}
	h.Next.OnBigDecimalFloat(value)
}
func (h *TEventPrinter) OnUUID(value []byte) {
	h.Print(UUID(value))
	h.Next.OnUUID(value)
}
func (h *TEventPrinter) OnTime(value time.Time) {
	h.Print(GT(value))
	h.Next.OnTime(value)
}
func (h *TEventPrinter) OnCompactTime(value compact_time.Time) {
	h.Print(CT(value))
	h.Next.OnCompactTime(value)
}
func (h *TEventPrinter) OnArray(arrayType events.ArrayType, elementCount uint64, value []byte) {
	switch arrayType {
	case events.ArrayTypeString:
		h.Print(S(string(value)))
	case events.ArrayTypeResourceID:
		h.Print(RID(string(value)))
	case events.ArrayTypeCustomBinary:
		h.Print(CUB(value))
	case events.ArrayTypeCustomText:
		h.Print(CUT(string(value)))
	case events.ArrayTypeBoolean:
		h.Print(AB(elementCount, value))
	case events.ArrayTypeInt8:
		h.Print(AI8(arrays.BytesToInt8Slice(value)))
	case events.ArrayTypeInt16:
		h.Print(AI16(arrays.BytesToInt16Slice(value)))
	case events.ArrayTypeInt32:
		h.Print(AI32(arrays.BytesToInt32Slice(value)))
	case events.ArrayTypeInt64:
		h.Print(AI64(arrays.BytesToInt64Slice(value)))
	case events.ArrayTypeUint8:
		h.Print(AU8(value))
	case events.ArrayTypeUint16:
		h.Print(AU16(arrays.BytesToUint16Slice(value)))
	case events.ArrayTypeUint32:
		h.Print(AU32(arrays.BytesToUint32Slice(value)))
	case events.ArrayTypeUint64:
		h.Print(AU64(arrays.BytesToUint64Slice(value)))
	case events.ArrayTypeFloat16:
		h.Print(AF16(value))
	case events.ArrayTypeFloat32:
		h.Print(AF32(arrays.BytesToFloat32Slice(value)))
	case events.ArrayTypeFloat64:
		h.Print(AF64(arrays.BytesToFloat64Slice(value)))
	case events.ArrayTypeUUID:
		h.Print(AUU(value))
	default:
		panic(fmt.Errorf("TODO: Typed array support for %v", arrayType))
	}
	h.Next.OnArray(arrayType, elementCount, value)
}
func (h *TEventPrinter) OnStringlikeArray(arrayType events.ArrayType, value string) {
	switch arrayType {
	case events.ArrayTypeString:
		h.Print(S(value))
	case events.ArrayTypeResourceID:
		h.Print(RID(value))
	case events.ArrayTypeCustomText:
		h.Print(CUT(value))
	default:
		panic(fmt.Errorf("BUG: Array type %v is not stringlike", arrayType))
	}
	h.Next.OnStringlikeArray(arrayType, value)
}
func (h *TEventPrinter) OnArrayBegin(arrayType events.ArrayType) {
	switch arrayType {
	case events.ArrayTypeString:
		h.Print(SB())
	case events.ArrayTypeResourceID:
		h.Print(RB())
	case events.ArrayTypeCustomBinary:
		h.Print(CBB())
	case events.ArrayTypeCustomText:
		h.Print(CTB())
	case events.ArrayTypeBoolean:
		h.Print(ABB())
	case events.ArrayTypeInt8:
		h.Print(AI8B())
	case events.ArrayTypeInt16:
		h.Print(AI16B())
	case events.ArrayTypeInt32:
		h.Print(AI32B())
	case events.ArrayTypeInt64:
		h.Print(AI64B())
	case events.ArrayTypeUint8:
		h.Print(AU8B())
	case events.ArrayTypeUint16:
		h.Print(AU16B())
	case events.ArrayTypeUint32:
		h.Print(AU32B())
	case events.ArrayTypeUint64:
		h.Print(AU64B())
	case events.ArrayTypeFloat16:
		h.Print(AF16B())
	case events.ArrayTypeFloat32:
		h.Print(AF32B())
	case events.ArrayTypeFloat64:
		h.Print(AF64B())
	case events.ArrayTypeUUID:
		h.Print(AUUB())
	default:
		panic(fmt.Errorf("TODO: Typed array support for %v", arrayType))
	}
	h.Next.OnArrayBegin(arrayType)
}
func (h *TEventPrinter) OnArrayChunk(l uint64, moreChunksFollow bool) {
	h.Print(AC(l, moreChunksFollow))
	h.Next.OnArrayChunk(l, moreChunksFollow)
}
func (h *TEventPrinter) OnArrayData(data []byte) {
	h.Print(AD(data))
	h.Next.OnArrayData(data)
}
func (h *TEventPrinter) OnList() {
	h.Print(L())
	h.Next.OnList()
}
func (h *TEventPrinter) OnMap() {
	h.Print(M())
	h.Next.OnMap()
}
func (h *TEventPrinter) OnMarkup() {
	h.Print(MUP())
	h.Next.OnMarkup()
}
func (h *TEventPrinter) OnMetadata() {
	h.Print(META())
	h.Next.OnMetadata()
}
func (h *TEventPrinter) OnComment() {
	h.Print(CMT())
	h.Next.OnComment()
}
func (h *TEventPrinter) OnEnd() {
	h.Print(E())
	h.Next.OnEnd()
}
func (h *TEventPrinter) OnMarker() {
	h.Print(MARK())
	h.Next.OnMarker()
}
func (h *TEventPrinter) OnReference() {
	h.Print(REF())
	h.Next.OnReference()
}
func (h *TEventPrinter) OnConcatenate() {
	h.Print(CAT())
	h.Next.OnConcatenate()
}
func (h *TEventPrinter) OnConstant(name []byte, explicitValue bool) {
	h.Print(CONST(string(name), explicitValue))
	h.Next.OnConstant(name, explicitValue)
}
func (h *TEventPrinter) OnEndDocument() {
	h.Print(ED())
	h.Next.OnEndDocument()
}
func (h *TEventPrinter) OnNan(signaling bool) {
	if signaling {
		h.Print(SNAN())
	} else {
		h.Print(NAN())
	}
	h.Next.OnNan(signaling)
}

// ----------------------------------------------------------------------------
// Event receiver
// ----------------------------------------------------------------------------

// Event receiver receives data events and stores them to an array which can be
// inspected, printed, or played back.
type TEventStore struct {
	Events []*TEvent
}

func NewTEventStore() *TEventStore {
	return &TEventStore{
		Events: make([]*TEvent, 1024),
	}
}
func (h *TEventStore) add(event *TEvent) {
	h.Events = append(h.Events, event)
}
func (h *TEventStore) OnVersion(version uint64)                  { h.add(V(version)) }
func (h *TEventStore) OnPadding(count int)                       { h.add(PAD(count)) }
func (h *TEventStore) OnNA()                                     { h.add(NA()) }
func (h *TEventStore) OnBool(value bool)                         { h.add(B(value)) }
func (h *TEventStore) OnTrue()                                   { h.add(TT()) }
func (h *TEventStore) OnFalse()                                  { h.add(FF()) }
func (h *TEventStore) OnPositiveInt(value uint64)                { h.add(PI(value)) }
func (h *TEventStore) OnNegativeInt(value uint64)                { h.add(NI(value)) }
func (h *TEventStore) OnInt(value int64)                         { h.add(I(value)) }
func (h *TEventStore) OnBigInt(value *big.Int)                   { h.add(BI(value)) }
func (h *TEventStore) OnFloat(value float64)                     { h.add(F(value)) }
func (h *TEventStore) OnBigFloat(value *big.Float)               { h.add(newTEvent(TEventBigFloat, value, nil)) }
func (h *TEventStore) OnDecimalFloat(value compact_float.DFloat) { h.add(DF(value)) }
func (h *TEventStore) OnBigDecimalFloat(value *apd.Decimal)      { h.add(BDF(value)) }
func (h *TEventStore) OnUUID(value []byte)                       { h.add(UUID(CloneBytes(value))) }
func (h *TEventStore) OnTime(value time.Time)                    { h.add(GT(value)) }
func (h *TEventStore) OnCompactTime(value compact_time.Time)     { h.add(CT(value)) }
func (h *TEventStore) OnArray(arrayType events.ArrayType, elementCount uint64, value []byte) {
	switch arrayType {
	case events.ArrayTypeString:
		h.add(S(string(value)))
	case events.ArrayTypeResourceID:
		h.add(RID(string(value)))
	case events.ArrayTypeCustomBinary:
		h.add(CUB(CloneBytes(value)))
	case events.ArrayTypeCustomText:
		h.add(CUT(string(value)))
	case events.ArrayTypeBoolean:
		h.add(AB(elementCount, CloneBytes(value)))
	case events.ArrayTypeInt8:
		h.add(AI8(arrays.BytesToInt8Slice(value)))
	case events.ArrayTypeInt16:
		h.add(AI16(arrays.BytesToInt16Slice(value)))
	case events.ArrayTypeInt32:
		h.add(AI32(arrays.BytesToInt32Slice(value)))
	case events.ArrayTypeInt64:
		h.add(AI64(arrays.BytesToInt64Slice(value)))
	case events.ArrayTypeUint8:
		h.add(AU8(CloneBytes(value)))
	case events.ArrayTypeUint16:
		h.add(AU16(arrays.BytesToUint16Slice(value)))
	case events.ArrayTypeUint32:
		h.add(AU32(arrays.BytesToUint32Slice(value)))
	case events.ArrayTypeUint64:
		h.add(AU64(arrays.BytesToUint64Slice(value)))
	case events.ArrayTypeFloat16:
		h.add(AF16(CloneBytes(value)))
	case events.ArrayTypeFloat32:
		h.add(AF32(arrays.BytesToFloat32Slice(value)))
	case events.ArrayTypeFloat64:
		h.add(AF64(arrays.BytesToFloat64Slice(value)))
	case events.ArrayTypeUUID:
		h.add(AUU(CloneBytes(value)))
	default:
		panic(fmt.Errorf("TODO: Typed array support for %v", arrayType))
	}
}
func (h *TEventStore) OnStringlikeArray(arrayType events.ArrayType, value string) {
	switch arrayType {
	case events.ArrayTypeString:
		h.add(S(value))
	case events.ArrayTypeResourceID:
		h.add(RID(value))
	case events.ArrayTypeCustomText:
		h.add(CUT(value))
	default:
		panic(fmt.Errorf("BUG: Array type %v is not stringlike", arrayType))
	}
}
func (h *TEventStore) OnArrayBegin(arrayType events.ArrayType) {
	switch arrayType {
	case events.ArrayTypeString:
		h.add(SB())
	case events.ArrayTypeResourceID:
		h.add(RB())
	case events.ArrayTypeCustomBinary:
		h.add(CBB())
	case events.ArrayTypeCustomText:
		h.add(CTB())
	case events.ArrayTypeBoolean:
		h.add(ABB())
	case events.ArrayTypeInt8:
		h.add(AI8B())
	case events.ArrayTypeInt16:
		h.add(AI16B())
	case events.ArrayTypeInt32:
		h.add(AI32B())
	case events.ArrayTypeInt64:
		h.add(AI64B())
	case events.ArrayTypeUint8:
		h.add(AU8B())
	case events.ArrayTypeUint16:
		h.add(AU16B())
	case events.ArrayTypeUint32:
		h.add(AU32B())
	case events.ArrayTypeUint64:
		h.add(AU64B())
	case events.ArrayTypeFloat16:
		h.add(AF16B())
	case events.ArrayTypeFloat32:
		h.add(AF32B())
	case events.ArrayTypeFloat64:
		h.add(AF64B())
	case events.ArrayTypeUUID:
		h.add(AUUB())
	default:
		panic(fmt.Errorf("TODO: Typed array support for %v", arrayType))
	}
}
func (h *TEventStore) OnArrayChunk(l uint64, moreChunks bool) { h.add(AC(l, moreChunks)) }
func (h *TEventStore) OnArrayData(data []byte)                { h.add(AD(CloneBytes(data))) }
func (h *TEventStore) OnList()                                { h.add(L()) }
func (h *TEventStore) OnMap()                                 { h.add(M()) }
func (h *TEventStore) OnMarkup()                              { h.add(MUP()) }
func (h *TEventStore) OnMetadata()                            { h.add(META()) }
func (h *TEventStore) OnComment()                             { h.add(CMT()) }
func (h *TEventStore) OnEnd()                                 { h.add(E()) }
func (h *TEventStore) OnMarker()                              { h.add(MARK()) }
func (h *TEventStore) OnReference()                           { h.add(REF()) }
func (h *TEventStore) OnConcatenate()                         { h.add(CAT()) }
func (h *TEventStore) OnConstant(n []byte, e bool)            { h.add(CONST(string(n), e)) }
func (h *TEventStore) OnBeginDocument() {
	h.Events = h.Events[:0]
	h.add(BD())
}
func (h *TEventStore) OnEndDocument() { h.add(ED()) }
func (h *TEventStore) OnNan(signaling bool) {
	if signaling {
		h.add(SNAN())
	} else {
		h.add(NAN())
	}
}

// ----------------------------------------------------------------------------
// Testing structures
// ----------------------------------------------------------------------------

// These are just complex structures used by a lot of the subsystem tests.

type TestingInnerStruct struct {
	Inner int
}

type TestingOuterStruct struct {
	Bo     bool
	PBo    *bool
	By     byte
	PBy    *byte
	I      int
	PI     *int
	I8     int8
	PI8    *int8
	I16    int16
	PI16   *int16
	I32    int32
	PI32   *int32
	I64    int64
	PI64   *int64
	U      uint
	PU     *uint
	U8     uint8
	PU8    *uint8
	U16    uint16
	PU16   *uint16
	U32    uint32
	PU32   *uint32
	U64    uint64
	PU64   *uint64
	BI     big.Int
	PBI    *big.Int
	F32    float32
	PF32   *float32
	F64    float64
	PF64   *float64
	BF     big.Float
	PBF    *big.Float
	DF     compact_float.DFloat
	BDF    apd.Decimal
	PBDF   *apd.Decimal
	St     string
	Au8    [4]byte
	Su8    []byte
	Sl     []interface{}
	M      map[interface{}]interface{}
	IS     TestingInnerStruct
	PIS    *TestingInnerStruct
	Time   time.Time
	PTime  *time.Time
	CTime  compact_time.Time
	PCTime compact_time.Time
	PURL   *url.URL
	URL    url.URL
}

func (_this *TestingOuterStruct) GetRepresentativeEvents(includeFakes bool) (events []*TEvent) {
	ade := func(e ...*TEvent) {
		events = append(events, e...)
	}
	adv := func(value interface{}) {
		ade(EventForValue(value))
	}
	anv := func(name string, value interface{}) {
		adv(name)
		adv(value)
	}
	ane := func(name string, e ...*TEvent) {
		adv(name)
		ade(e...)
	}

	ade(M())

	anv("Bo", _this.Bo)
	anv("PBo", _this.PBo)
	anv("By", _this.By)
	anv("PBy", _this.PBy)
	anv("I", _this.I)
	anv("PI", _this.PI)
	anv("I8", _this.I8)
	anv("PI8", _this.PI8)
	anv("I16", _this.I16)
	anv("PI16", _this.PI16)
	anv("I32", _this.I32)
	anv("PI32", _this.PI32)
	anv("I64", _this.I64)
	anv("PI64", _this.PI64)
	anv("U", _this.U)
	anv("PU", _this.PU)
	anv("U8", _this.U8)
	anv("PU8", _this.PU8)
	anv("U16", _this.U16)
	anv("PU16", _this.PU16)
	anv("U32", _this.U32)
	anv("PU32", _this.PU32)
	anv("U64", _this.U64)
	anv("PU64", _this.PU64)
	anv("BI", _this.BI)
	anv("PBI", _this.PBI)
	anv("F32", _this.F32)
	anv("PF32", _this.PF32)
	anv("F64", _this.F64)
	anv("PF64", _this.PF64)
	anv("BF", _this.BF)
	anv("PBF", _this.PBF)
	anv("DF", _this.DF)
	anv("BDF", _this.BDF)
	anv("PBDF", _this.PBDF)
	anv("St", _this.St)
	ane("Au8", AU8(_this.Au8[:]))
	anv("Su8", _this.Su8)

	ane("Sl", L())
	for _, v := range _this.Sl {
		adv(v)
	}
	ade(E())

	ane("M", M())
	for k, v := range _this.M {
		adv(k)
		adv(v)
	}
	ade(E())

	ane("IS", M())
	anv("Inner", _this.IS.Inner)
	ade(E())

	if _this.PIS != nil {
		ane("PIS", M())
		anv("Inner", _this.PIS.Inner)
		ade(E())
	}

	anv("Time", _this.Time)
	anv("PTime", _this.PTime)
	anv("CTime", _this.CTime)
	anv("PCTime", _this.PCTime)
	anv("PURL", _this.PURL)

	if includeFakes {
		ane("F1", B(true))
		ane("F2", B(false))
		ane("F3", I(1))
		ane("F4", I(-1))
		ane("F5", F(1.1))
		ane("F6", BF(NewBigFloat("1.1", 10, 2)))
		ane("F7", DF(NewDFloat("1.1")))
		ane("F8", BDF(NewBDF("1.1")))
		ane("F9", NA())
		ane("F10", BI(NewBigInt("1000", 10)))
		ane("F12", NAN())
		ane("F13", SNAN())
		ane("F14", UUID([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
		ane("F15", GT(_this.Time))
		ane("F16", CT(_this.CTime))
		ane("F17", AU8([]byte{1}))
		ane("F18", S("xyz"))
		ane("F19", RID("http://example.com"))
		// ane("F20", cust([]byte{1}))
		ane("FakeList", L(), E())
		ane("FakeMap", M(), E())
		ane("FakeDeep", L(), M(), S("A"), L(),
			B(true),
			B(false),
			I(1),
			I(-1),
			F(1.1),
			BF(NewBigFloat("1.1", 10, 2)),
			DF(NewDFloat("1.1")),
			BDF(NewBDF("1.1")),
			NA(),
			BI(NewBigInt("1000", 10)),
			NAN(),
			SNAN(),
			UUID([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
			GT(_this.Time),
			CT(_this.CTime),
			AU8([]byte{1}),
			S("xyz"),
			RID("http://example.com"),
			// cust([]byte{1}),
			E(), E(), E())
	}

	ade(E())
	return
}

func NewTestingOuterStruct(baseValue int) *TestingOuterStruct {
	_this := new(TestingOuterStruct)
	_this.Init(baseValue)
	return _this
}

func NewBlankTestingOuterStruct() *TestingOuterStruct {
	_this := new(TestingOuterStruct)
	_this.CTime.Year = 1
	_this.CTime.Month = 1
	_this.CTime.Day = 1
	return _this
}

func (_this *TestingOuterStruct) Init(baseValue int) {
	_this.Bo = baseValue&1 == 1
	_this.PBo = &_this.Bo
	_this.By = byte(baseValue + int(unsafe.Offsetof(_this.By)))
	_this.PBy = &_this.By
	_this.I = 100000 + baseValue + int(unsafe.Offsetof(_this.I))
	_this.PI = &_this.I
	_this.I8 = int8(baseValue + int(unsafe.Offsetof(_this.I8)))
	_this.PI8 = &_this.I8
	_this.I16 = int16(-1000 - baseValue - int(unsafe.Offsetof(_this.I16)))
	_this.PI16 = &_this.I16
	_this.I32 = int32(1000000000 + baseValue + int(unsafe.Offsetof(_this.I32)))
	_this.PI32 = &_this.I32
	_this.I64 = int64(1000000000000000) + int64(baseValue+int(unsafe.Offsetof(_this.I64)))
	_this.PI64 = &_this.I64
	_this.U = uint(1000000 + baseValue + int(unsafe.Offsetof(_this.U)))
	_this.PU = &_this.U
	_this.U8 = uint8(baseValue + int(unsafe.Offsetof(_this.U8)))
	_this.PU8 = &_this.U8
	_this.U16 = uint16(10000 + baseValue + int(unsafe.Offsetof(_this.U16)))
	_this.PU16 = &_this.U16
	_this.U32 = uint32(100000000 + baseValue + int(unsafe.Offsetof(_this.U32)))
	_this.PU32 = &_this.U32
	_this.U64 = uint64(1000000000000) + uint64(baseValue+int(unsafe.Offsetof(_this.U64)))
	_this.PU64 = &_this.U64
	_this.PBI = NewBigInt(fmt.Sprintf("-10000000000000000000000000000000000000%v", unsafe.Offsetof(_this.PBI)), 10)
	_this.BI = *_this.PBI
	_this.F32 = float32(1000000+baseValue+int(unsafe.Offsetof(_this.F32))) + 0.5
	_this.PF32 = &_this.F32
	_this.F64 = float64(1000000000000) + float64(baseValue+int(unsafe.Offsetof(_this.F64))) + 0.5
	_this.PF64 = &_this.F64
	_this.PBF = NewBigFloat("12345678901234567890123.1234567", 10, 30)
	_this.BF = *_this.PBF
	_this.DF = NewDFloat(fmt.Sprintf("-100000000000000%ve-1000000", unsafe.Offsetof(_this.DF)))
	_this.PBDF = NewBDF("-1.234567890123456789777777777777777777771234e-10000")
	_this.BDF = *_this.PBDF
	_this.St = GenerateString(baseValue+5, baseValue)
	_this.Au8[0] = byte(baseValue + int(unsafe.Offsetof(_this.Au8)))
	_this.Au8[1] = byte(baseValue + int(unsafe.Offsetof(_this.Au8)+1))
	_this.Au8[2] = byte(baseValue + int(unsafe.Offsetof(_this.Au8)+2))
	_this.Au8[3] = byte(baseValue + int(unsafe.Offsetof(_this.Au8)+3))
	_this.Su8 = GenerateBytes(baseValue+1, baseValue)
	_this.M = make(map[interface{}]interface{})
	for i := 0; i < baseValue+2; i++ {
		_this.Sl = append(_this.Sl, i)
		_this.M[fmt.Sprintf("key%v", i)] = i
	}
	_this.IS.Inner = baseValue + 15
	_this.PIS = new(TestingInnerStruct)
	_this.PIS.Inner = baseValue + 16
	testTime := time.Date(30000+baseValue, time.Month(1), 1, 1, 1, 1, 0, time.UTC)
	_this.PTime = &testTime
	_this.CTime = NewTS(-1000, 1, 1, 1, 1, 1, 1, "Europe/Berlin")
	_this.PURL, _ = url.Parse(fmt.Sprintf("http://example.com/%v", baseValue))
}
