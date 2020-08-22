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

package cbe

import (
	"bytes"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/kstenerud/go-concise-encoding/debug"
	"github.com/kstenerud/go-concise-encoding/events"
	"github.com/kstenerud/go-concise-encoding/rules"
	"github.com/kstenerud/go-concise-encoding/test"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
	"github.com/kstenerud/go-describe"
	"github.com/kstenerud/go-equivalence"
)

const (
	version = 1

	typeDecimal        = 0x65
	typePosInt         = 0x66
	typeNegInt         = 0x67
	typePosInt8        = 0x68
	typeNegInt8        = 0x69
	typePosInt16       = 0x6a
	typeNegInt16       = 0x6b
	typePosInt32       = 0x6c
	typeNegInt32       = 0x6d
	typePosInt64       = 0x6e
	typeNegInt64       = 0x6f
	typeFloat16        = 0x70
	typeFloat32        = 0x72
	typeFloat64        = 0x73
	typeUUID           = 0x74
	typeReserved74     = 0x74
	typeReserved75     = 0x75
	typeComment        = 0x76
	typeMetadata       = 0x77
	typeMarkup         = 0x78
	typeMap            = 0x79
	typeList           = 0x7a
	typeEndContainer   = 0x7b
	typeFalse          = 0x7c
	typeTrue           = 0x7d
	typeNil            = 0x7e
	typePadding        = 0x7f
	typeString0        = 0x80
	typeString1        = 0x81
	typeString2        = 0x82
	typeString3        = 0x83
	typeString4        = 0x84
	typeString5        = 0x85
	typeString6        = 0x86
	typeString7        = 0x87
	typeString8        = 0x88
	typeString9        = 0x89
	typeString10       = 0x8a
	typeString11       = 0x8b
	typeString12       = 0x8c
	typeString13       = 0x8d
	typeString14       = 0x8e
	typeString15       = 0x8f
	typeString         = 0x90
	typeVerbatimString = 0x91
	typeURI            = 0x92
	typeCustomBinary   = 0x93
	typeCustomText     = 0x94
	typeArray          = 0x95
	typeReserved96     = 0x96
	typeMarker         = 0x97
	typeReference      = 0x98
	typeDate           = 0x99
	typeTime           = 0x9a
	typeTimestamp      = 0x9b
)

func NewBigInt(str string) *big.Int {
	return test.NewBigInt(str)
}

func NewBigFloat(str string, significantDigits int) *big.Float {
	return test.NewBigFloat(str, significantDigits)
}

func NewDFloat(str string) compact_float.DFloat {
	return test.NewDFloat(str)
}

func NewBDF(str string) *apd.Decimal {
	return test.NewBDF(str)
}

func NewURI(uriString string) *url.URL {
	return test.NewURI(uriString)
}

func TT() *test.TEvent                       { return test.TT() }
func FF() *test.TEvent                       { return test.FF() }
func I(v int64) *test.TEvent                 { return test.I(v) }
func F(v float64) *test.TEvent               { return test.F(v) }
func BF(v *big.Float) *test.TEvent           { return test.BF(v) }
func DF(v compact_float.DFloat) *test.TEvent { return test.DF(v) }
func BDF(v *apd.Decimal) *test.TEvent        { return test.BDF(v) }
func V(v uint64) *test.TEvent                { return test.V(v) }
func N() *test.TEvent                        { return test.N() }
func PAD(v int) *test.TEvent                 { return test.PAD(v) }
func B(v bool) *test.TEvent                  { return test.B(v) }
func PI(v uint64) *test.TEvent               { return test.PI(v) }
func NI(v uint64) *test.TEvent               { return test.NI(v) }
func BI(v *big.Int) *test.TEvent             { return test.BI(v) }
func NAN() *test.TEvent                      { return test.NAN() }
func SNAN() *test.TEvent                     { return test.SNAN() }
func UUID(v []byte) *test.TEvent             { return test.UUID(v) }
func GT(v time.Time) *test.TEvent            { return test.GT(v) }
func CT(v *compact_time.Time) *test.TEvent   { return test.CT(v) }
func BIN(v []byte) *test.TEvent              { return test.BIN(v) }
func S(v string) *test.TEvent                { return test.S(v) }
func VS(v string) *test.TEvent               { return test.VS(v) }
func URI(v string) *test.TEvent              { return test.URI(v) }
func CUB(v []byte) *test.TEvent              { return test.CUB(v) }
func CUT(v string) *test.TEvent              { return test.CUT(v) }
func BB() *test.TEvent                       { return test.BB() }
func SB() *test.TEvent                       { return test.SB() }
func VB() *test.TEvent                       { return test.VB() }
func UB() *test.TEvent                       { return test.UB() }
func CBB() *test.TEvent                      { return test.CBB() }
func CTB() *test.TEvent                      { return test.CTB() }
func AC(l uint64, more bool) *test.TEvent    { return test.AC(l, more) }
func AD(v []byte) *test.TEvent               { return test.AD(v) }
func L() *test.TEvent                        { return test.L() }
func M() *test.TEvent                        { return test.M() }
func MUP() *test.TEvent                      { return test.MUP() }
func META() *test.TEvent                     { return test.META() }
func CMT() *test.TEvent                      { return test.CMT() }
func E() *test.TEvent                        { return test.E() }
func MARK() *test.TEvent                     { return test.MARK() }
func REF() *test.TEvent                      { return test.REF() }
func BD() *test.TEvent                       { return test.BD() }
func ED() *test.TEvent                       { return test.ED() }

var DebugPrintEvents = false

func decodeToEvents(document []byte, withRules bool) (evts []*test.TEvent, err error) {
	var topLevelReceiver events.DataEventReceiver
	ter := test.NewTER()
	topLevelReceiver = ter
	if withRules {
		topLevelReceiver = rules.NewRules(topLevelReceiver, nil)
	}
	if DebugPrintEvents {
		topLevelReceiver = test.NewStdoutTEventPrinter(topLevelReceiver)
	}
	err = NewDecoder(nil).Decode(bytes.NewBuffer(document), topLevelReceiver)
	evts = ter.Events
	return
}

func encodeEvents(events ...*test.TEvent) []byte {
	buffer := &bytes.Buffer{}
	encoder := NewEncoder(nil)
	encoder.PrepareToEncode(buffer)
	test.InvokeEvents(encoder, events...)
	return buffer.Bytes()
}

func assertDecode(t *testing.T, document []byte, expectedEvents ...*test.TEvent) (successful bool, events []*test.TEvent) {
	actualEvents, err := decodeToEvents(document, false)
	if err != nil {
		t.Error(err)
		return
	}

	if len(expectedEvents) > 0 {
		if !equivalence.IsEquivalent(actualEvents, expectedEvents) {
			t.Errorf("Expected events %v but got %v", expectedEvents, actualEvents)
			return
		}
	}
	events = actualEvents
	successful = true
	return
}

func assertDecodeWithRules(t *testing.T, document []byte, expectedEvents ...*test.TEvent) (successful bool, events []*test.TEvent) {
	actualEvents, err := decodeToEvents(document, true)
	if err != nil {
		t.Error(err)
		return
	}

	if len(expectedEvents) > 0 {
		if !equivalence.IsEquivalent(actualEvents, expectedEvents) {
			t.Errorf("Expected events %v but got %v", expectedEvents, actualEvents)
			return
		}
	}
	events = actualEvents
	successful = true
	return
}

func assertDecodeFails(t *testing.T, document []byte) {
	_, err := decodeToEvents(document, false)
	if err == nil {
		t.Errorf("Expected decode to fail")
	}
}

func assertEncode(t *testing.T, expectedDocument []byte, events ...*test.TEvent) (successful bool) {
	actualDocument := encodeEvents(events...)
	if !equivalence.IsEquivalent(actualDocument, expectedDocument) {
		t.Errorf("Expected document %v but got %v", describe.D(expectedDocument), describe.D(actualDocument))
		return
	}
	successful = true
	return
}

func assertDecodeEncode(t *testing.T, document []byte, expectedEvents ...*test.TEvent) (successful bool) {
	debug.DebugOptions.PassThroughPanics = true
	defer func() { debug.DebugOptions.PassThroughPanics = false }()

	successful, actualEvents := assertDecode(t, document, expectedEvents...)
	if !successful {
		return
	}
	return assertEncode(t, document, actualEvents...)
}

func assertMarshal(t *testing.T, value interface{}, expectedDocument []byte) (successful bool) {
	document, err := NewMarshaler(nil).MarshalToDocument(value)
	if err != nil {
		t.Error(err)
		return
	}
	if !equivalence.IsEquivalent(document, expectedDocument) {
		t.Errorf("Expected document %v but got %v", describe.D(expectedDocument), describe.D(document))
		return
	}
	successful = true
	return
}

func assertUnmarshal(t *testing.T, expectedValue interface{}, document []byte) (successful bool) {
	actualValue, err := NewUnmarshaler(nil).UnmarshalFromDocument([]byte(document), expectedValue)
	if err != nil {
		t.Error(err)
		return
	}

	if !equivalence.IsEquivalent(actualValue, expectedValue) {
		t.Errorf("Expected unmarshaled [%v] but got [%v]", describe.D(expectedValue), describe.D(actualValue))
		return
	}
	successful = true
	return
}

func assertMarshalUnmarshal(t *testing.T, expectedValue interface{}, expectedDocument []byte) (successful bool) {
	if !assertMarshal(t, expectedValue, expectedDocument) {
		return
	}
	return assertUnmarshal(t, expectedValue, expectedDocument)
}