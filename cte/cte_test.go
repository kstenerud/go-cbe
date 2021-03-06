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

package cte

import (
	"bytes"
	"math"
	"testing"

	"github.com/kstenerud/go-concise-encoding/internal/common"

	"github.com/kstenerud/go-concise-encoding/options"
	"github.com/kstenerud/go-concise-encoding/test"
)

// TODO: Remove this when releasing V1
func TestCTEVersion1(t *testing.T) {
	assertDecode(t, nil, "c1 1", BD(), V(ceVer), PI(1), ED())
}

func TestCTEVersion(t *testing.T) {
	// Valid
	assertDecodeEncode(t, nil, nil, "c0\n1", BD(), V(ceVer), PI(1), ED())
	assertDecode(t, nil, "\r\n\t c0 1", BD(), V(ceVer), PI(1), ED())
	assertDecode(t, nil, "c0     \r\n\t\t\t1", BD(), V(ceVer), PI(1), ED())

	// Missing whitespace
	assertDecodeFails(t, "c0{}")

	// Too big
	assertDecodeFails(t, "c000000000000000000000000000000000 ")

	// Not numeric
	for i := 0; i < 0x100; i++ {
		if i >= '0' && i <= '9' {
			continue
		}
		document := string([]byte{'c', byte(i)})
		assertDecodeFails(t, document)
	}

	// Disallowed version numbers
	for i := 0; i < 0x100; i++ {
		switch i {
		case 'c', 'C', ' ', '\n', '\r', '\t':
			continue
		default:
			document := string([]byte{byte(i)})
			assertDecodeFails(t, document)
		}
	}
}

func TestCTENA(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n@na", BD(), V(ceVer), NA(), ED())
	assertDecodeFails(t, "c0 -@na")
	assertDecodeFails(t, "c0 @na:@na:1")
	assertDecodeFails(t, "c0 @na:@na:@na")
	assertDecodeFails(t, "c0 [@na:@na:1]")
}

func TestCTEBool(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n@true", BD(), V(ceVer), TT(), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@false", BD(), V(ceVer), FF(), ED())

	assertEncode(t, nil, "c0\n@false", BD(), V(ceVer), B(false), ED())
	assertEncode(t, nil, "c0\n@true", BD(), V(ceVer), B(true), ED())

	assertDecodeFails(t, "c0 @truer")
	assertDecodeFails(t, "c0 @falser")
	assertDecodeFails(t, "c0 -@true")
	assertDecodeFails(t, "c0 -@false")
}

func TestCTEDecimalInt(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n0", BD(), V(ceVer), PI(0), ED())
	assertDecodeEncode(t, nil, nil, "c0\n123", BD(), V(ceVer), PI(123), ED())
	assertDecodeEncode(t, nil, nil, "c0\n9412504234235366", BD(), V(ceVer), PI(9412504234235366), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-49523", BD(), V(ceVer), NI(49523), ED())
	assertDecodeEncode(t, nil, nil, "c0\n10000000000000000000000000000", BD(), V(ceVer), BI(NewBigInt("10000000000000000000000000000", 10)), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-10000000000000000000000000000", BD(), V(ceVer), BI(NewBigInt("-10000000000000000000000000000", 10)), ED())
	assertDecode(t, nil, "c0 100_00_0_00000000000_00000000_000_0", BD(), V(ceVer), BI(NewBigInt("10000000000000000000000000000", 10)), ED())
	assertDecode(t, nil, "c0 -4_9_5__2___3", BD(), V(ceVer), NI(49523), ED())
	assertEncode(t, nil, "c0\n100", BD(), V(ceVer), I(100), ED())
	assertEncode(t, nil, "c0\n-100", BD(), V(ceVer), I(-100), ED())
	assertDecode(t, nil, "c0 100", BD(), V(ceVer), PI(100), ED())
	assertDecode(t, nil, "c0 -100", BD(), V(ceVer), NI(100), ED())

	assertDecodeFails(t, "c0 1f")
	assertDecodeFails(t, "c0 -1f")
}

func TestCTEBinaryInt(t *testing.T) {
	assertDecode(t, nil, "c0 0b0", BD(), V(ceVer), PI(0), ED())
	assertDecode(t, nil, "c0 0b1", BD(), V(ceVer), PI(1), ED())
	assertDecode(t, nil, "c0 0b101", BD(), V(ceVer), PI(5), ED())
	assertDecode(t, nil, "c0 0b0010100", BD(), V(ceVer), PI(20), ED())
	assertDecode(t, nil, "c0 -0b100", BD(), V(ceVer), NI(4), ED())
	assertDecode(t, nil, "c0 -0b_1_0_0", BD(), V(ceVer), NI(4), ED())

	assertDecode(t, nil, "c0 0b10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		BD(), V(ceVer), BI(NewBigInt("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 2)), ED())
	assertDecode(t, nil, "c0 -0b10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		BD(), V(ceVer), BI(NewBigInt("-10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 2)), ED())
	assertDecode(t, nil, "c0 0b100000000000000_000000000000_000000000000000000000000_000000000000000000000_0000000000000000000000000000000000000000_0",
		BD(), V(ceVer), BI(NewBigInt("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 2)), ED())

	assertDecodeFails(t, "c0 0b2")
}

func TestCTEOctalInt(t *testing.T) {
	assertDecode(t, nil, "c0 0o0", BD(), V(ceVer), PI(0), ED())
	assertDecode(t, nil, "c0 0o1", BD(), V(ceVer), PI(1), ED())
	assertDecode(t, nil, "c0 0o7", BD(), V(ceVer), PI(7), ED())
	assertDecode(t, nil, "c0 0o71", BD(), V(ceVer), PI(57), ED())
	assertDecode(t, nil, "c0 0o644", BD(), V(ceVer), PI(420), ED())
	assertDecode(t, nil, "c0 -0o777", BD(), V(ceVer), NI(511), ED())
	assertDecode(t, nil, "c0 -0o_7__7___7", BD(), V(ceVer), NI(511), ED())

	assertDecode(t, nil, "c0 0o1000000000000000000000000000000000000000000000",
		BD(), V(ceVer), BI(NewBigInt("1000000000000000000000000000000000000000000000", 8)), ED())
	assertDecode(t, nil, "c0 -0o1000000000000000000000000000000000000000000000",
		BD(), V(ceVer), BI(NewBigInt("-1000000000000000000000000000000000000000000000", 8)), ED())
	assertDecode(t, nil, "c0 0o1_0000000000000_0000000000000_000000000000000000_0",
		BD(), V(ceVer), BI(NewBigInt("1000000000000000000000000000000000000000000000", 8)), ED())

	assertDecodeFails(t, "c0 0o9")
}

func TestCTEHexInt(t *testing.T) {
	assertDecode(t, nil, "c0 0x0", BD(), V(ceVer), PI(0), ED())
	assertDecode(t, nil, "c0 0x1", BD(), V(ceVer), PI(1), ED())
	assertDecode(t, nil, "c0 0xf", BD(), V(ceVer), PI(0xf), ED())
	assertDecode(t, nil, "c0 0xfedcba9876543210", BD(), V(ceVer), PI(0xfedcba9876543210), ED())
	assertDecode(t, nil, "c0 0xFEDCBA9876543210", BD(), V(ceVer), PI(0xfedcba9876543210), ED())
	assertDecode(t, nil, "c0 -0x88", BD(), V(ceVer), NI(0x88), ED())
	assertDecode(t, nil, "c0 -0x_8_8__5_a_f__d", BD(), V(ceVer), NI(0x885afd), ED())

	assertDecode(t, nil, "c0 0x1000000000000000000000000000000000000000000000",
		BD(), V(ceVer), BI(NewBigInt("1000000000000000000000000000000000000000000000", 16)), ED())
	assertDecode(t, nil, "c0 -0x1000000000000000000000000000000000000000000000",
		BD(), V(ceVer), BI(NewBigInt("-1000000000000000000000000000000000000000000000", 16)), ED())
	assertDecode(t, nil, "c0 0x1_00000000000_00000000000_0000000000000000000000_0",
		BD(), V(ceVer), BI(NewBigInt("1000000000000000000000000000000000000000000000", 16)), ED())

	assertDecodeFails(t, "c0 0xg")
}

func TestCTEFloat(t *testing.T) {
	assertDecode(t, nil, "c0 0.0", BD(), V(ceVer), DF(NewDFloat("0")), ED())
	assertDecode(t, nil, "c0 -0.0", BD(), V(ceVer), DF(NewDFloat("-0")), ED())

	assertDecodeEncode(t, nil, nil, "c0\n1.5", BD(), V(ceVer), DF(NewDFloat("1.5")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n1.125", BD(), V(ceVer), DF(NewDFloat("1.125")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n1.125e+10", BD(), V(ceVer), DF(NewDFloat("1.125e+10")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n1.125e-10", BD(), V(ceVer), DF(NewDFloat("1.125e-10")), ED())
	assertDecode(t, nil, "c0 1.125e10", BD(), V(ceVer), DF(NewDFloat("1.125e+10")), ED())

	assertDecodeEncode(t, nil, nil, "c0\n-1.5", BD(), V(ceVer), DF(NewDFloat("-1.5")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-1.125", BD(), V(ceVer), DF(NewDFloat("-1.125")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-1.125e+10", BD(), V(ceVer), DF(NewDFloat("-1.125e+10")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-1.125e-10", BD(), V(ceVer), DF(NewDFloat("-1.125e-10")), ED())
	assertDecode(t, nil, "c0 -1.125e10", BD(), V(ceVer), DF(NewDFloat("-1.125e10")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n1.0000000000000000001", BD(), V(ceVer), BDF(NewBDF("1.0000000000000000001")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-1.0000000000000000001", BD(), V(ceVer), BDF(NewBDF("-1.0000000000000000001")), ED())

	assertDecodeEncode(t, nil, nil, "c0\n0.5", BD(), V(ceVer), DF(NewDFloat("0.5")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n0.125", BD(), V(ceVer), DF(NewDFloat("0.125")), ED())
	assertDecode(t, nil, "c0 0.125e+10", BD(), V(ceVer), DF(NewDFloat("0.125e+10")), ED())
	assertDecode(t, nil, "c0 0.125e-10", BD(), V(ceVer), DF(NewDFloat("0.125e-10")), ED())
	assertDecode(t, nil, "c0 0.125e10", BD(), V(ceVer), DF(NewDFloat("0.125e10")), ED())

	assertDecode(t, nil, "c0 -0.5", BD(), V(ceVer), DF(NewDFloat("-0.5")), ED())
	assertDecode(t, nil, "c0 -0.125", BD(), V(ceVer), DF(NewDFloat("-0.125")), ED())
	assertDecode(t, nil, "c0 -0.125e+10", BD(), V(ceVer), DF(NewDFloat("-0.125e+10")), ED())
	assertDecode(t, nil, "c0 -0.125e-10", BD(), V(ceVer), DF(NewDFloat("-0.125e-10")), ED())
	assertDecode(t, nil, "c0 -0.125e10", BD(), V(ceVer), DF(NewDFloat("-0.125e10")), ED())
	assertDecode(t, nil, "c0 -0.125E+10", BD(), V(ceVer), DF(NewDFloat("-0.125e+10")), ED())
	assertDecode(t, nil, "c0 -0.125E-10", BD(), V(ceVer), DF(NewDFloat("-0.125e-10")), ED())
	assertDecode(t, nil, "c0 -0.125E10", BD(), V(ceVer), DF(NewDFloat("-0.125e10")), ED())

	assertDecode(t, nil, "c0 -1.50000000000000000000000001E10000", BD(), V(ceVer), BDF(NewBDF("-1.50000000000000000000000001E10000")), ED())
	assertDecode(t, nil, "c0 1.50000000000000000000000001E10000", BD(), V(ceVer), BDF(NewBDF("1.50000000000000000000000001E10000")), ED())

	assertDecode(t, nil, "c0 1_._1_2_5_e+1_0", BD(), V(ceVer), DF(NewDFloat("1.125e+10")), ED())

	assertDecode(t, nil, "c0 0.1500000000000000000000000000000000000000000000000001e+10000",
		BD(), V(ceVer), BDF(NewBDF("0.1500000000000000000000000000000000000000000000000001e+10000")), ED())
	assertDecode(t, nil, "c0 -0.1500000000000000000000000000000000000000000000000001e+10000",
		BD(), V(ceVer), BDF(NewBDF("-0.1500000000000000000000000000000000000000000000000001e+10000")), ED())
	assertDecode(t, nil, "c0 0.1_50000000000_00000000000_000000000000_0000000000000000_1e+100_0_0",
		BD(), V(ceVer), BDF(NewBDF("0.1500000000000000000000000000000000000000000000000001e+10000")), ED())

	assertEncode(t, nil, "c0\n@nan", BD(), V(ceVer), F(common.QuietNan), ED())
	assertEncode(t, nil, "c0\n@snan", BD(), V(ceVer), F(common.SignalingNan), ED())

	assertEncode(t, nil, "c0\n1.1", BD(), V(ceVer), BF(NewBigFloat("1.1", 10, 2)), ED())

	assertDecodeFails(t, "c0 [-0.5.4]")
	assertDecodeFails(t, "c0 [-0,5.4]")
	assertDecodeFails(t, "c0 [0.5.4]")
	assertDecodeFails(t, "c0 [0,5.4]")
	assertDecodeFails(t, "c0 [-@blah]")
	assertDecodeFails(t, "c0 [1.1.1]")
	assertDecodeFails(t, "c0 [1,1]")
	assertDecodeFails(t, "c0 [1.1e4e5]")
	assertDecodeFails(t, "c0 [0.a]")
	assertDecodeFails(t, "c0 [0.5et]")
	assertDecodeFails(t, "c0 [0.5e99999999999999999999999]")
}

func TestCTEHexFloat(t *testing.T) {
	assertDecode(t, nil, "c0 0x0.0", BD(), V(ceVer), F(0x0.0p0), ED())
	assertDecode(t, nil, "c0 0x0.1", BD(), V(ceVer), F(0x0.1p0), ED())
	assertDecode(t, nil, "c0 0x0.1p+10", BD(), V(ceVer), F(0x0.1p+10), ED())
	assertDecode(t, nil, "c0 0x0.1p-10", BD(), V(ceVer), F(0x0.1p-10), ED())
	assertDecode(t, nil, "c0 0x0.1p10", BD(), V(ceVer), F(0x0.1p10), ED())

	assertDecode(t, nil, "c0 0x1.0", BD(), V(ceVer), F(0x1.0p0), ED())
	assertDecode(t, nil, "c0 0x1.1", BD(), V(ceVer), F(0x1.1p0), ED())
	assertDecode(t, nil, "c0 0xf.1p+10", BD(), V(ceVer), F(0xf.1p+10), ED())
	assertDecode(t, nil, "c0 0xf.1p-10", BD(), V(ceVer), F(0xf.1p-10), ED())
	assertDecode(t, nil, "c0 0xf.1p10", BD(), V(ceVer), F(0xf.1p10), ED())

	assertDecode(t, nil, "c0 -0x1.0", BD(), V(ceVer), F(-0x1.0p0), ED())
	assertDecode(t, nil, "c0 -0x1.1", BD(), V(ceVer), F(-0x1.1p0), ED())
	assertDecode(t, nil, "c0 -0xf.1p+10", BD(), V(ceVer), F(-0xf.1p+10), ED())
	assertDecode(t, nil, "c0 -0xf.1p-10", BD(), V(ceVer), F(-0xf.1p-10), ED())
	assertDecode(t, nil, "c0 -0xf.1p10", BD(), V(ceVer), F(-0xf.1p10), ED())

	assertDecode(t, nil, "c0 -0x0.0", BD(), V(ceVer), F(-0x0.0p0), ED())
	assertDecode(t, nil, "c0 -0x0.1", BD(), V(ceVer), F(-0x0.1p0), ED())
	assertDecode(t, nil, "c0 -0x0.1p+10", BD(), V(ceVer), F(-0x0.1p+10), ED())
	assertDecode(t, nil, "c0 -0x0.1p-10", BD(), V(ceVer), F(-0x0.1p-10), ED())
	assertDecode(t, nil, "c0 -0x0.1p10", BD(), V(ceVer), F(-0x0.1p10), ED())

	// Everything too big for float64
	bigExpected := NewBigFloat("-1.54fffe2ac00592375b427ap100000", 16, 22)
	assertDecode(t, nil, "c0 -0x1.54fffe2ac00592375b427ap100000", BD(), V(ceVer), BF(bigExpected), ED())
	bigExpected = bigExpected.Neg(bigExpected)
	assertDecode(t, nil, "c0 0x1.54fffe2ac00592375b427ap100000", BD(), V(ceVer), BF(bigExpected), ED())

	// Coefficient too big for float64
	bigExpected = NewBigFloat("-1.54fffe2ac00592375b427ap100", 16, 22)
	assertDecode(t, nil, "c0 -0x1.54fffe2ac00592375b427ap100", BD(), V(ceVer), BF(bigExpected), ED())
	bigExpected = bigExpected.Neg(bigExpected)
	assertDecode(t, nil, "c0 0x1.54fffe2ac00592375b427ap100", BD(), V(ceVer), BF(bigExpected), ED())
	assertDecode(t, nil, "c0 0x1.5_4fffe2ac_0059237_5b42_7ap1_00", BD(), V(ceVer), BF(bigExpected), ED())
	assertDecode(t, nil, "c0 0x1.5_4FFFE2AC_0059237_5B42_7AP1_00", BD(), V(ceVer), BF(bigExpected), ED())

	// Exponent too big for float64
	bigExpected = NewBigFloat("-1.8p100000", 16, 16)
	assertDecode(t, nil, "c0 -0x1.8p100000", BD(), V(ceVer), BF(bigExpected), ED())
	bigExpected = bigExpected.Neg(bigExpected)
	assertDecode(t, nil, "c0 0x1.8p100000", BD(), V(ceVer), BF(bigExpected), ED())

	assertDecode(t, nil, "c0 -0x_0_._1_p_1_0", BD(), V(ceVer), F(-0x0.1p10), ED())

	bigExpected = NewBigFloat("8.000000000000001p100", 16, 16)
	assertDecode(t, nil, "c0 0x8.000000000000001p100", BD(), V(ceVer), BF(bigExpected), ED())
	bigExpected = bigExpected.Neg(bigExpected)
	assertDecode(t, nil, "c0 -0x8.000000000000001p100", BD(), V(ceVer), BF(bigExpected), ED())

	assertDecodeFails(t, "[c1 -0x0.5.4]")
	assertDecodeFails(t, "[c1 -0x0,5.4]")
	assertDecodeFails(t, "[c1 0x0.5.4]")
	assertDecodeFails(t, "[c1 0x0,5.4]")
	assertDecodeFails(t, "[c1 -0x@blah]")
	assertDecodeFails(t, "[c1 0x1.1.1]")
	assertDecodeFails(t, "[c1 0x1,1]")
	assertDecodeFails(t, "[c1 0x1.1p4p5]")
	assertDecodeFails(t, "[c1 -0x0.l]")
	assertDecodeFails(t, "[c1 -0x0.5pj]")
	assertDecodeFails(t, "[c1 -0x0.5p1000000000000000000000000000]")
}

func TestCTEUUID(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n@fedcba98-7654-3210-aaaa-bbbbbbbbbbbb", BD(), V(ceVer),
		UUID([]byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb}), ED())
	assertDecode(t, nil, "c0 @FEDCBA98-7654-3210-AAAA-BBBBBBBBBBBB", BD(), V(ceVer),
		UUID([]byte{0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@00000000-0000-0000-0000-000000000000", BD(), V(ceVer),
		UUID([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}), ED())

	assertDecodeFails(t, "c0 @fedcba98-7654-3210-aaaa-bbbbbbbbbbb")
	assertDecodeFails(t, "c0 @fedcba98-7654-3210-aaaa-bbbbbbbbbbbbb")
	assertEncodeFails(t, nil, BD(), V(ceVer), UUID([]byte{0xfe, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb}), ED())
	assertEncodeFails(t, nil, BD(), V(ceVer), UUID([]byte{0xfe, 0xdc, 0xff, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10, 0xaa, 0xaa, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb}), ED())
}

func TestCTEDate(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n2000-01-01", BD(), V(ceVer), CT(test.NewDate(2000, 1, 1)), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-2000-12-31", BD(), V(ceVer), CT(test.NewDate(-2000, 12, 31)), ED())

	assertDecodeFails(t, "c0 0-01-01")
	assertDecodeFails(t, "c0 --2000-01-01")
	assertDecodeFails(t, "c0 0a-01-01")

	assertDecodeFails(t, "c0 2000-013-01")
	assertDecodeFails(t, "c0 2000-30-1")
	assertDecodeFails(t, "c0 2000-0-10")
	assertDecodeFails(t, "c0 2000-1a-10")
	assertDecodeFails(t, "c0 2000-0a-10")
	assertDecodeFails(t, "c0 2000-a-10")

	assertDecodeFails(t, "c0 2000-01-011")
	assertDecodeFails(t, "c0 2000-01-99")
	assertDecodeFails(t, "c0 2000-10-0")
	assertDecodeFails(t, "c0 2000-10-1a")
	assertDecodeFails(t, "c0 2000-10-0a")
	assertDecodeFails(t, "c0 2000-10-a")
}

func TestCTETime(t *testing.T) {
	assertDecode(t, nil, "c0 1:45:00", BD(), V(ceVer), CT(test.NewTime(1, 45, 0, 0, "")), ED())
	assertDecode(t, nil, "c0 01:45:00", BD(), V(ceVer), CT(test.NewTime(1, 45, 0, 0, "")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n23:59:59.101", BD(), V(ceVer), CT(test.NewTime(23, 59, 59, 101000000, "")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n10:00:01.93/America/Los_Angeles", BD(), V(ceVer), CT(test.NewTime(10, 0, 1, 930000000, "America/Los_Angeles")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n10:00:01.93/89.92/1.10", BD(), V(ceVer), CT(test.NewTimeLL(10, 0, 1, 930000000, 8992, 110)), ED())
	assertDecode(t, nil, "c0 10:00:01.93/89.90/1.1", BD(), V(ceVer), CT(test.NewTimeLL(10, 0, 1, 930000000, 8990, 110)), ED())
	assertDecode(t, nil, "c0 10:00:01.93/89.9/1.10", BD(), V(ceVer), CT(test.NewTimeLL(10, 0, 1, 930000000, 8990, 110)), ED())
	assertDecode(t, nil, "c0 10:00:01.93/0/0", BD(), V(ceVer), CT(test.NewTimeLL(10, 0, 1, 930000000, 0, 0)), ED())
	assertDecode(t, nil, "c0 10:00:01.93/1/1", BD(), V(ceVer), CT(test.NewTimeLL(10, 0, 1, 930000000, 100, 100)), ED())

	assertDecodeFails(t, "c0 001:45:00")
	assertDecodeFails(t, "c0 30:45:10")
	assertDecodeFails(t, "c0 1a:45:10")
	assertDecodeFails(t, "c0 0a:45:10")

	assertDecodeFails(t, "c0 1:045:00")
	assertDecodeFails(t, "c0 1:99:10")
	assertDecodeFails(t, "c0 1:1a:10")
	assertDecodeFails(t, "c0 1:0a:10")
	assertDecodeFails(t, "c0 1:a:10")

	assertDecodeFails(t, "c0 1:45:001")
	assertDecodeFails(t, "c0 1:45:99")
	assertDecodeFails(t, "c0 1:45:1e")
	assertDecodeFails(t, "c0 1:45:0e")
	assertDecodeFails(t, "c0 1:45:e")

	assertDecodeFails(t, "c0 1:45:00.3012544133")
	assertDecodeFails(t, "c0 1:45:00.301254f")
	assertDecodeFails(t, "c0 1:45:00.1f")
	assertDecodeFails(t, "c0 1:45:00.0f")
	assertDecodeFails(t, "c0 1:45:00.f")

	assertDecodeFails(t, "c0 10:00:01.93/89.92/1.10a")
	assertDecodeFails(t, "c0 10:00:01.93/89.92a/1.10")
	assertDecodeFails(t, "c0 10:00:01.93/89.92/a.10")
	assertDecodeFails(t, "c0 10:00:01.93/w89.92/1.10")
	assertDecodeFails(t, "c0 10:00:01.93/89.92/1.10a")
	assertDecodeFails(t, "c0 10:00:01.93/89.92/")
	assertDecodeFails(t, "c0 10:00:01.93/89.92/1.101")
	assertDecodeFails(t, "c0 10:00:01.93//1.10")
	assertDecodeFails(t, "c0 10:00:01.93/89.925/1.10")
}

func TestCTETimestamp(t *testing.T) {
	assertDecode(t, nil, "c0 2000-01-01/19:31:44.901554/Z", BD(), V(ceVer), CT(test.NewTS(2000, 1, 1, 19, 31, 44, 901554000, "Z")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n2020-01-15/13:41:00.000599", BD(), V(ceVer), CT(test.NewTS(2020, 1, 15, 13, 41, 0, 599000, "")), ED())
	assertDecode(t, nil, "c0 2020-01-15/13:41:00.000599", BD(), V(ceVer), CT(test.NewTS(2020, 1, 15, 13, 41, 0, 599000, "")), ED())
	assertDecodeEncode(t, nil, nil, "c0\n2020-01-15/10:00:01.93/89.92/1.10", BD(), V(ceVer), CT(test.NewTSLL(2020, 1, 15, 10, 0, 1, 930000000, 8992, 110)), ED())

	assertDecodeFails(t, "c0 0-01-01/19:31:44.901554")
	assertDecodeFails(t, "c0 1a-01-01/19:31:44.901554")

	assertDecodeFails(t, "c0 2000-0-01/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-001-01/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-100-01/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-1a-01/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-0a-01/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-a-01/19:31:44.901554")

	assertDecodeFails(t, "c0 2000-01-0/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-001/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-100/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-1a/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-0a/19:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-a/19:31:44.901554")

	assertDecodeFails(t, "c0 2000-01-01/019:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/25:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/1a:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/0a:31:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/a:31:44.901554")

	assertDecodeFails(t, "c0 2000-01-01/19:031:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:310:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:1a:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:0a:44.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:a:44.901554")

	assertDecodeFails(t, "c0 2000-01-01/19:31:044.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:31:440.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:31:1a.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:31:0a.901554")
	assertDecodeFails(t, "c0 2000-01-01/19:31:a.901554")

	assertDecodeFails(t, "c0 2000-01-01/19:31:44.9015544348")
	assertDecodeFails(t, "c0 2000-01-01/19:31:44.1a")
	assertDecodeFails(t, "c0 2000-01-01/19:31:44.0a")
	assertDecodeFails(t, "c0 2000-01-01/19:31:44.a")

	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92/1.10a")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92a/1.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92/a1.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/a89.92/1.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92/a")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/a/1.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92/")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93//1.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92/1999.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/8965.92/1.10")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.92/1.a")
	assertDecodeFails(t, "c0 2020-01-15/10:00:01.93/89.a/1.10")

	ts := NewTS(2020, 1, 15, 10, 0, 1, 930000000, "")
	gotime, err := ts.AsGoTime()
	if err != nil {
		panic(err)
	}
	assertEncode(t, nil, "c0\n2020-01-15/10:00:01.93", BD(), V(ceVer), GT(gotime), ED())
}

func TestCTEConstant(t *testing.T) {
	// TODO: Update test support to allow passing rules opts
	// assertDecodeEncode(t, nil, nil, "c0\n#someconst", BD(), V(ceVer), CONST("someconst", false), ED())
	// assertDecodeEncode(t, nil, nil, `c0
	// [
	//    #c
	//    1
	// ]`, BD(), V(ceVer), L(), CONST("c", false), PI(1), E(), ED())
	// assertDecodeEncode(t, nil, nil, `c0
	// {
	//    #c = 1
	// }`, BD(), V(ceVer), M(), CONST("c", false), PI(1), E(), ED())

	assertDecodeEncode(t, nil, nil, "c0\n#someconst:xyz", BD(), V(ceVer), CONST("someconst", true), S("xyz"), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    #c:xyz
    1
]`, BD(), V(ceVer), L(), CONST("c", true), S("xyz"), PI(1), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    #c:xyz = 1
}`, BD(), V(ceVer), M(), CONST("c", true), S("xyz"), PI(1), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    #c:123 = 1
}`, BD(), V(ceVer), M(), CONST("c", true), PI(123), PI(1), E(), ED())
}

func TestCTEQuotedString(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
"test string"`, BD(), V(ceVer), S("test string"), ED())
	assertDecode(t, nil, `c0 "test\nstring"`, BD(), V(ceVer), S("test\nstring"), ED())
	assertDecode(t, nil, `c0 "test\rstring"`, BD(), V(ceVer), S("test\rstring"), ED())
	assertDecode(t, nil, `c0 "test\tstring"`, BD(), V(ceVer), S("test\tstring"), ED())
	assertDecodeEncode(t, nil, nil, `c0
"test\"string"`, BD(), V(ceVer), S("test\"string"), ED())
	assertDecode(t, nil, `c0 "test\*string"`, BD(), V(ceVer), S("test*string"), ED())
	assertDecode(t, nil, `c0 "test\/string"`, BD(), V(ceVer), S("test/string"), ED())
	assertDecodeEncode(t, nil, nil, `c0
"test\\string"`, BD(), V(ceVer), S("test\\string"), ED())
	assertDecodeEncode(t, nil, nil, `c0
"test\11string"`, BD(), V(ceVer), S("test\u0001string"), ED())
	assertDecodeEncode(t, nil, nil, `c0
"test\4206dstring"`, BD(), V(ceVer), S("test\u206dstring"), ED())
	assertDecode(t, nil, `c0 "test\4206Dstring"`, BD(), V(ceVer), S("test\u206dstring"), ED())
	assertDecode(t, nil, `c0 "test\
string"`, BD(), V(ceVer), S("teststring"), ED())
	assertDecode(t, nil, "c0 \"test\\\r\nstring\"", BD(), V(ceVer), S("teststring"), ED())

	assertDecodeFails(t, `c0 "test\x"`)
	assertDecodeFails(t, `c0 "\1g"`)
}

func TestCTECustomBinary(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n|cb 12 34 56 78|", BD(), V(ceVer), CUB([]byte{0x12, 0x34, 0x56, 0x78}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n|cb ab cd|", BD(), V(ceVer), CUB([]byte{0xab, 0xcd}), ED())
	assertDecode(t, nil, "c0 |cb AB CD|", BD(), V(ceVer), CUB([]byte{0xab, 0xcd}), ED())
	assertDecodeFails(t, "c0 |cb qwer|")
}

func TestCTECustomText(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n|ct something(123)|", BD(), V(ceVer), CUT("something(123)"), ED())
	assertDecodeEncode(t, nil, nil, `c0
|ct some\\thing("123")|`, BD(), V(ceVer), CUT("some\\thing(\"123\")"), ED())
	assertDecodeEncode(t, nil, nil, `c0
|ct some\nthing\11(123)|`, BD(), V(ceVer), CUT("some\nthing\u0001(123)"), ED())
	assertDecodeEncode(t, nil, nil, `c0
|ct something('123\r\n\t')|`, BD(), V(ceVer), CUT("something('123\r\n\t')"), ED())

	assertDecodeFails(t, `c0 |ct something('123\r\n\t\x')|`)
}

func TestCTEUnquotedString(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\na", BD(), V(ceVer), S("a"), ED())
	assertDecodeEncode(t, nil, nil, "c0\nabcd", BD(), V(ceVer), S("abcd"), ED())
	assertDecodeEncode(t, nil, nil, "c0\n_-123aF", BD(), V(ceVer), S("_-123aF"), ED())
	assertDecodeEncode(t, nil, nil, "c0\n新しい", BD(), V(ceVer), S("新しい"), ED())
}

func TestCTEInvalidString(t *testing.T) {
	assertDecodeFails(t, "c0 [a|b]")
	assertDecodeFails(t, "c0 [a*b]")
}

func TestCTEVerbatimString(t *testing.T) {
	assertDecodeFails(t, `c0 "\."`)
	assertDecodeFails(t, `c0 "\.A"`)
	assertDecodeFails(t, `c0 "\.A "`)
	assertDecodeFails(t, `c0 "\.A xyz"`)
	assertDecode(t, nil, `c0 "\.A \n\n\n\n\n\n\n\n\n\nA"`, BD(), V(ceVer), S(`\n\n\n\n\n\n\n\n\n\n`), ED())
	assertDecode(t, nil, `c0 "\.A aA"`, BD(), V(ceVer), S("a"), ED())
	assertDecode(t, nil, "c0 \"\\.A\taA\"", BD(), V(ceVer), S("a"), ED())
	assertDecode(t, nil, "c0 \"\\.A\naA\"", BD(), V(ceVer), S("a"), ED())
	assertDecode(t, nil, "c0 \"\\.A\r\naA\"", BD(), V(ceVer), S("a"), ED())
	assertDecode(t, nil, `c0 "\.#ENDOFSTRING a test\nwith \.stuff#ENDOFSTRING"`, BD(), V(ceVer), S(`a test\nwith \.stuff`), ED())
}

func TestCTERID(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
|r http://example.com|`, BD(), V(ceVer), RID("http://example.com"), ED())
	assertDecodeEncode(t, nil, nil, `c0
|r http://x.com/\||`, BD(), V(ceVer), RID(`http://x.com/|`), ED())
}

func TestCTEArrayBoolean(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n|b|", BD(), V(ceVer), AB(0, []byte{}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n|b 0|", BD(), V(ceVer), AB(1, []byte{0x00}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n|b 1|", BD(), V(ceVer), AB(1, []byte{0x01}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n|b 1011001|", BD(), V(ceVer), AB(7, []byte{0b1001101}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n|b 10110011|", BD(), V(ceVer), AB(8, []byte{0b11001101}), ED())
	assertDecodeEncode(t, nil, nil, "c0\n|b 101100111|", BD(), V(ceVer), AB(9, []byte{0b11001101, 0b1}), ED())

	assertEncode(t, nil, "c0\n|b|", BD(), V(ceVer), ABB(), AC(0, false), ED())
	assertEncode(t, nil, "c0\n|b 0|", BD(), V(ceVer), ABB(), AC(1, false), AD([]byte{0x00}), ED())
	assertEncode(t, nil, "c0\n|b 1|", BD(), V(ceVer), ABB(), AC(1, false), AD([]byte{0x01}), ED())
	assertEncode(t, nil, "c0\n|b 1011001|", BD(), V(ceVer), ABB(), AC(7, false), AD([]byte{0b1001101}), ED())
	assertEncode(t, nil, "c0\n|b 10110011|", BD(), V(ceVer), ABB(), AC(8, false), AD([]byte{0b11001101}), ED())
	assertEncode(t, nil, "c0\n|b 101100111|", BD(), V(ceVer), ABB(), AC(9, false), AD([]byte{0b11001101, 0b1}), ED())

	assertDecode(t, nil, "c0\n|b |", BD(), V(ceVer), AB(0, []byte{}), ED())
	assertDecode(t, nil, "c0\n|b 0 |", BD(), V(ceVer), AB(1, []byte{0x00}), ED())
	assertDecode(t, nil, "c0\n|b 1 |", BD(), V(ceVer), AB(1, []byte{0x01}), ED())
	assertDecode(t, nil, "c0\n|b 1 01 1 001 |", BD(), V(ceVer), AB(7, []byte{0b1001101}), ED())
	assertDecode(t, nil, "c0\n|b 1 0 1 1 0 0 1 1 |", BD(), V(ceVer), AB(8, []byte{0b11001101}), ED())
	assertDecode(t, nil, "c0\n|b  10  110 0 1 1   1    |", BD(), V(ceVer), AB(9, []byte{0b11001101, 0b1}), ED())
}

func TestCTEArrayUintX(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n|u8x f1 93|", BD(), V(ceVer), AU8([]byte{0xf1, 0x93}), ED())
	assertDecode(t, nil, "c0\n|u8x f 93 |", BD(), V(ceVer), AU8([]byte{0xf, 0x93}), ED())
	assertDecodeFails(t, "c0\n|u8x f14 93|")
	assertDecodeFails(t, "c0\n|u8x f1o 93|")

	assertDecodeEncode(t, nil, nil, "c0\n|u16x f122 9385|", BD(), V(ceVer), AU16([]uint16{0xf122, 0x9385}), ED())
	assertDecode(t, nil, "c0\n|u16x f12 95|", BD(), V(ceVer), AU16([]uint16{0xf12, 0x95}), ED())
	assertDecodeFails(t, "c0\n|u16x f129e 95|")
	assertDecodeFails(t, "c0\n|u16x f12j 95|")

	assertDecodeEncode(t, nil, nil, "c0\n|u32x 7ddf8134 93cd7aac|", BD(), V(ceVer), AU32([]uint32{0x7ddf8134, 0x93cd7aac}), ED())
	assertDecode(t, nil, "c0\n|u32x 7ddf834 93aac|", BD(), V(ceVer), AU32([]uint32{0x7ddf834, 0x93aac}), ED())
	assertDecodeFails(t, "c0\n|u32x 7ddf8134e 93cd7aac|")
	assertDecodeFails(t, "c0\n|u32x 7ddf81x 93cd7aac|")

	assertDecodeEncode(t, nil, nil, "c0\n|u64x 83ff9ac2445aace7 94ff7ac3219465c1|", BD(), V(ceVer), AU64([]uint64{0x83ff9ac2445aace7, 0x94ff7ac3219465c1}), ED())
	assertDecode(t, nil, "c0\n|u64x 83ff9ac245aace7 94ff79465c1|", BD(), V(ceVer), AU64([]uint64{0x83ff9ac245aace7, 0x94ff79465c1}), ED())
	assertDecodeFails(t, "c0\n|u64x 83ff9ac2445aace72 94ff7ac3219465c1|")
	assertDecodeFails(t, "c0\n|u64x 83ff9ac2l 94ff7ac3219465c1|")
}

func TestCTEArrayInt8(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Int8 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|i8b|", BD(), V(ceVer), AI8([]int8{}), ED())
	assertDecode(t, nil, "c0\n|i8b |", BD(), V(ceVer), AI8([]int8{}), ED())
	assertDecodeEncode(t, nil, eOpts, "c0\n|i8b 0 1 -10 101 1111111 -10000000|", BD(), V(ceVer), AI8([]int8{0, 1, -2, 5, 0x7f, -0x80}), ED())
	assertEncode(t, eOpts, "c0\n|i8b|", BD(), V(ceVer), AI8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8b 1|", BD(), V(ceVer), AI8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|i8b 1|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8b 1 0|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	eOpts.DefaultFormats.Array.Int8 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i8o 0 -10 50 -127|", BD(), V(ceVer), AI8([]int8{0o0, -0o10, 0o50, -0o127}), ED())
	assertEncode(t, eOpts, "c0\n|i8o|", BD(), V(ceVer), AI8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8o 1|", BD(), V(ceVer), AI8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|i8o 1|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8o 1 0|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	eOpts.DefaultFormats.Array.Int8 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|i8 0 10 -50 127 -128|", BD(), V(ceVer), AI8([]int8{0, 10, -50, 127, -128}), ED())
	assertEncode(t, eOpts, "c0\n|i8|", BD(), V(ceVer), AI8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8 1|", BD(), V(ceVer), AI8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|i8 1|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8 1 0|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	eOpts.DefaultFormats.Array.Int8 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i8x 0 1 -50 7f -80|", BD(), V(ceVer), AI8([]int8{0x00, 0x01, -0x50, 0x7f, -0x80}), ED())
	assertEncode(t, eOpts, "c0\n|i8x|", BD(), V(ceVer), AI8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8x 1|", BD(), V(ceVer), AI8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|i8x 1|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i8x 1 0|", BD(), V(ceVer), AI8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	assertDecode(t, nil, "c0 |i8 00 01 -01 0b101 -0b110 0B101 -0B110 0o10 -0o11 0O10 -0O11 0x7f -0x80 0X7f -0X80|",
		BD(), V(ceVer), AI8([]int8{0, 1, -1, 5, -6, 5, -6, 8, -9, 8, -9, 127, -128, 127, -128}), ED())

	assertDecodeFails(t, "c0 |i8b 10000000|")
	assertDecodeFails(t, "c0 |i8b -10000001|")
	assertDecodeFails(t, "c0 |i8o 178|")
	assertDecodeFails(t, "c0 |i8o -179|")
	assertDecodeFails(t, "c0 |i8 128|")
	assertDecodeFails(t, "c0 |i8 -129|")
	assertDecodeFails(t, "c0 |i8x 80|")
	assertDecodeFails(t, "c0 |i8x -81|")
}

func TestCTEArrayUint8(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Uint8 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|u8b 0 1 10 101 1111111 10000000 11111111|", BD(), V(ceVer), AU8([]uint8{0, 1, 2, 5, 0x7f, 0x80, 0xff}), ED())
	assertEncode(t, eOpts, "c0\n|u8b|", BD(), V(ceVer), AU8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8b 1|", BD(), V(ceVer), AU8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|u8b 1|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8b 1 0|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	eOpts.DefaultFormats.Array.Uint8 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|u8o 0 10 50 127 254 377|", BD(), V(ceVer), AU8([]uint8{0o0, 0o10, 0o50, 0o127, 0o254, 0o377}), ED())
	assertEncode(t, eOpts, "c0\n|u8o|", BD(), V(ceVer), AU8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8o 1|", BD(), V(ceVer), AU8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|u8o 1|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8o 1 0|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	eOpts.DefaultFormats.Array.Uint8 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|u8 0 10 50 128 254 255|", BD(), V(ceVer), AU8([]uint8{0, 10, 50, 128, 254, 255}), ED())
	assertEncode(t, eOpts, "c0\n|u8|", BD(), V(ceVer), AU8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8 1|", BD(), V(ceVer), AU8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|u8 1|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8 1 0|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	eOpts.DefaultFormats.Array.Uint8 = options.CTEEncodingFormatHexadecimalZeroFilled
	assertDecodeEncode(t, nil, eOpts, "c0\n|u8x 00 01 50 7f 80 ff|", BD(), V(ceVer), AU8([]uint8{0x00, 0x01, 0x50, 0x7f, 0x80, 0xff}), ED())
	assertEncode(t, eOpts, "c0\n|u8x|", BD(), V(ceVer), AU8B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8x 01|", BD(), V(ceVer), AU8B(), AC(1, false), AD([]uint8{1}), ED())
	assertEncode(t, eOpts, "c0\n|u8x 01|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|u8x 01 00|", BD(), V(ceVer), AU8B(), AC(1, true), AD([]uint8{1}), AC(1, false), AD([]uint8{0}), ED())

	assertDecode(t, nil, "c0 |u8 00 01 01 0b101 0b110 0B101 0B110 0o10 0o11 0O10 0O11 0x7f 0x80 0X7f 0X80 0xff 0Xff|",
		BD(), V(ceVer), AU8([]uint8{0, 1, 1, 5, 6, 5, 6, 8, 9, 8, 9, 127, 128, 127, 128, 255, 255}), ED())

	assertDecodeFails(t, "c0 |u8b 100000000|")
	assertDecodeFails(t, "c0 |u8o 400|")
	assertDecodeFails(t, "c0 |u8 256|")
	assertDecodeFails(t, "c0 |u8x 100|")
}

func TestCTEArrayInt16(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Int16 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|i16b 0 1 -10 101 111111111111111 -1000000000000000|", BD(), V(ceVer), AI16([]int16{0, 1, -2, 5, 0x7fff, -0x8000}), ED())
	assertEncode(t, eOpts, "c0\n|i16b|", BD(), V(ceVer), AI16B(), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i16b 1|", BD(), V(ceVer), AI16B(), AC(1, false), AD([]uint8{1, 0}), ED())
	assertEncode(t, eOpts, "c0\n|i16b 1|", BD(), V(ceVer), AI16B(), AC(1, true), AD([]uint8{1, 0}), AC(0, false), ED())
	assertEncode(t, eOpts, "c0\n|i16b 1 0|", BD(), V(ceVer), AI16B(), AC(1, true), AD([]uint8{1, 0}), AC(1, false), AD([]uint8{0, 0}), ED())

	eOpts.DefaultFormats.Array.Int16 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i16o 0 -10 50 -77777|", BD(), V(ceVer), AI16([]int16{0o0, -0o10, 0o50, -0o77777}), ED())

	eOpts.DefaultFormats.Array.Int16 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|i16 0 10 -50 32767 -32768|", BD(), V(ceVer), AI16([]int16{0, 10, -50, 32767, -32768}), ED())

	eOpts.DefaultFormats.Array.Int16 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i16x 0 1 -50 7fff -8000|", BD(), V(ceVer), AI16([]int16{0x00, 0x01, -0x50, 0x7fff, -0x8000}), ED())

	assertDecode(t, nil, "c0 |i16 00 01 -01 0b101 -0b110 0B101 -0B110 0o10 -0o11 0O10 -0O11 0x7f -0x80 0X7fff -0X8000|",
		BD(), V(ceVer), AI16([]int16{0, 1, -1, 5, -6, 5, -6, 8, -9, 8, -9, 127, -128, 32767, -32768}), ED())

	assertDecodeFails(t, "c0 |i16b 1000000000000000|")
	assertDecodeFails(t, "c0 |i16b -1000000000000001|")
	assertDecodeFails(t, "c0 |i16o 100000|")
	assertDecodeFails(t, "c0 |i16o -100001|")
	assertDecodeFails(t, "c0 |i16 32768|")
	assertDecodeFails(t, "c0 |i16 -32769|")
	assertDecodeFails(t, "c0 |i16x 8000|")
	assertDecodeFails(t, "c0 |i16x -8001|")
}

func TestCTEArrayUint16(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Uint16 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|u16b 0 1 10 101 111111111111111 1000000000000000 1111111111111111|",
		BD(), V(ceVer), AU16([]uint16{0, 1, 2, 5, 0x7fff, 0x8000, 0xffff}), ED())

	eOpts.DefaultFormats.Array.Uint16 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|u16o 0 10 50 127 254 377 177777|",
		BD(), V(ceVer), AU16([]uint16{0o0, 0o10, 0o50, 0o127, 0o254, 0o377, 0o177777}), ED())

	eOpts.DefaultFormats.Array.Uint16 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|u16 0 10 50 128 254 255 65535|",
		BD(), V(ceVer), AU16([]uint16{0, 10, 50, 128, 254, 255, 65535}), ED())

	eOpts.DefaultFormats.Array.Uint16 = options.CTEEncodingFormatHexadecimalZeroFilled
	assertDecodeEncode(t, nil, eOpts, "c0\n|u16x 0000 0001 0050 007f 0080 00ff fffe|",
		BD(), V(ceVer), AU16([]uint16{0x00, 0x01, 0x50, 0x7f, 0x80, 0xff, 0xfffe}), ED())

	assertDecode(t, nil, "c0 |u16 00 01 01 0b101 0b110 0B101 0B110 0o10 0o11 0O10 0O11 0x7f 0x80 0X7f 0X80 0xff 0Xff|",
		BD(), V(ceVer), AU16([]uint16{0, 1, 1, 5, 6, 5, 6, 8, 9, 8, 9, 127, 128, 127, 128, 255, 255}), ED())

	assertDecodeFails(t, "c0 |u16b 10000000000000000|")
	assertDecodeFails(t, "c0 |u16o 200000|")
	assertDecodeFails(t, "c0 |u16 65536|")
	assertDecodeFails(t, "c0 |u16x 10000|")
}

func TestCTEArrayInt32(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Int32 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|i32b 0 1 -10 101 1111111111111111111111111111111 -10000000000000000000000000000000|",
		BD(), V(ceVer), AI32([]int32{0, 1, -2, 5, 0x7fffffff, -0x80000000}), ED())

	eOpts.DefaultFormats.Array.Int32 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i32o 0 -10 50 -17777777777|", BD(), V(ceVer), AI32([]int32{0o0, -0o10, 0o50, -0o17777777777}), ED())

	eOpts.DefaultFormats.Array.Int32 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|i32 0 10 -50 2147483647 -2147483648|", BD(), V(ceVer), AI32([]int32{0, 10, -50, 2147483647, -2147483648}), ED())

	eOpts.DefaultFormats.Array.Int32 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i32x 0 1 -50 7fffffff -80000000 7f6f5f4f|", BD(), V(ceVer), AI32([]int32{0x00, 0x01, -0x50, 0x7fffffff, -0x80000000, 0x7f6f5f4f}), ED())

	assertDecode(t, nil, "c0 |i32 00 01 -01 0b101 -0b110 0B101 -0B110 0o10 -0o11 0O10 -0O11 0x7f -0x80 0X7fffffff -0X80000000|",
		BD(), V(ceVer), AI32([]int32{0, 1, -1, 5, -6, 5, -6, 8, -9, 8, -9, 127, -128, 0x7fffffff, -0x80000000}), ED())

	assertDecodeFails(t, "c0 |i32b 100000000000000000000000000000000|")
	assertDecodeFails(t, "c0 |i32b -100000000000000000000000000000001|")
	assertDecodeFails(t, "c0 |i32o 20000000000|")
	assertDecodeFails(t, "c0 |i32o -20000000001|")
	assertDecodeFails(t, "c0 |i32 2147483648|")
	assertDecodeFails(t, "c0 |i32 -2147483649|")
	assertDecodeFails(t, "c0 |i32x 80000000|")
	assertDecodeFails(t, "c0 |i32x -80000001|")
}

func TestCTEArrayUint32(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Uint32 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|u32b 0 1 10 101 1111111111111111111111111111111 10000000000000000000000000000000 11111111111111111111111111111111|",
		BD(), V(ceVer), AU32([]uint32{0, 1, 2, 5, 0x7fffffff, 0x80000000, 0xffffffff}), ED())

	eOpts.DefaultFormats.Array.Uint32 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|u32o 0 10 50 127 254 377 177777 37777777776|",
		BD(), V(ceVer), AU32([]uint32{0o0, 0o10, 0o50, 0o127, 0o254, 0o377, 0o177777, 0o37777777776}), ED())

	eOpts.DefaultFormats.Array.Uint32 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|u32 0 10 50 128 254 255 65535 4294967294|",
		BD(), V(ceVer), AU32([]uint32{0, 10, 50, 128, 254, 255, 65535, 4294967294}), ED())

	eOpts.DefaultFormats.Array.Uint32 = options.CTEEncodingFormatHexadecimalZeroFilled
	assertDecodeEncode(t, nil, eOpts, "c0\n|u32x 00000000 00000001 00000050 0000007f 00000080 000000ff 0000ffff fffcfdfe|",
		BD(), V(ceVer), AU32([]uint32{0x00, 0x01, 0x50, 0x7f, 0x80, 0xff, 0xffff, 0xfffcfdfe}), ED())

	assertDecode(t, nil, "c0 |u32 00 01 01 0b101 0b110 0B101 0B110 0o10 0o11 0O10 0O11 0x7f 0x80 0X7f 0X80 0xff 0Xff 100000000|",
		BD(), V(ceVer), AU32([]uint32{0, 1, 1, 5, 6, 5, 6, 8, 9, 8, 9, 127, 128, 127, 128, 255, 255, 100000000}), ED())

	assertDecodeFails(t, "c0 |u32b 100000000000000000000000000000000|")
	assertDecodeFails(t, "c0 |u32o 40000000000|")
	assertDecodeFails(t, "c0 |u32 4294967296|")
	assertDecodeFails(t, "c0 |u32x 100000000|")
}

func TestCTEArrayInt64(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Int64 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|i64b 0 1 -10 101 111111111111111111111111111111111111111111111111111111111111111 -1000000000000000000000000000000000000000000000000000000000000000|",
		BD(), V(ceVer), AI64([]int64{0, 1, -2, 5, 0x7fffffffffffffff, -0x8000000000000000}), ED())

	eOpts.DefaultFormats.Array.Int64 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i64o 0 -10 50 -777777777777777777777|",
		BD(), V(ceVer), AI64([]int64{0o0, -0o10, 0o50, -0o777777777777777777777}), ED())

	eOpts.DefaultFormats.Array.Int64 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|i64 0 10 -50 9223372036854775807 -9223372036854775808|",
		BD(), V(ceVer), AI64([]int64{0, 10, -50, 9223372036854775807, -9223372036854775808}), ED())

	eOpts.DefaultFormats.Array.Int64 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|i64x 0 1 -50 7fffffffffffffff -8000000000000000 7f6f5f4f3f2f1f0f|",
		BD(), V(ceVer), AI64([]int64{0x00, 0x01, -0x50, 0x7fffffffffffffff, -0x8000000000000000, 0x7f6f5f4f3f2f1f0f}), ED())

	assertDecode(t, nil, "c0 |i64 00 01 -01 0b101 -0b110 0B101 -0B110 0o10 -0o11 0O10 -0O11 0x7f -0x80 0X7fffffffffffffff -0X8000000000000000|",
		BD(), V(ceVer), AI64([]int64{0, 1, -1, 5, -6, 5, -6, 8, -9, 8, -9, 127, -128, 0x7fffffffffffffff, -0x8000000000000000}), ED())

	assertDecodeFails(t, "c0 |i64b 1000000000000000000000000000000000000000000000000000000000000000|")
	assertDecodeFails(t, "c0 |i64b -1000000000000000000000000000000000000000000000000000000000000001|")
	assertDecodeFails(t, "c0 |i64o 1000000000000000000000|")
	assertDecodeFails(t, "c0 |i64o -1000000000000000000001|")
	assertDecodeFails(t, "c0 |i64 9223372036854775808|")
	assertDecodeFails(t, "c0 |i64 -9223372036854775809|")
	assertDecodeFails(t, "c0 |i64x 8000000000000000|")
	assertDecodeFails(t, "c0 |i64x -8000000000000001|")
}

func TestCTEArrayUint64(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Uint64 = options.CTEEncodingFormatBinary
	assertDecodeEncode(t, nil, eOpts, "c0\n|u64b 0 1 10 101 1111111111111111111111111111111 10000000000000000000000000000000 11111111111111111111111111111111|",
		BD(), V(ceVer), AU64([]uint64{0, 1, 2, 5, 0x7fffffff, 0x80000000, 0xffffffff}), ED())

	eOpts.DefaultFormats.Array.Uint64 = options.CTEEncodingFormatOctal
	assertDecodeEncode(t, nil, eOpts, "c0\n|u64o 0 10 50 127 254 377 177777 1777777777777777777777|",
		BD(), V(ceVer), AU64([]uint64{0o0, 0o10, 0o50, 0o127, 0o254, 0o377, 0o177777, 0o1777777777777777777777}), ED())

	eOpts.DefaultFormats.Array.Uint64 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|u64 0 10 50 128 254 255 65535 4294967294 18446744073709551615|",
		BD(), V(ceVer), AU64([]uint64{0, 10, 50, 128, 254, 255, 65535, 4294967294, 18446744073709551615}), ED())

	eOpts.DefaultFormats.Array.Uint64 = options.CTEEncodingFormatHexadecimalZeroFilled
	assertDecodeEncode(t, nil, eOpts, "c0\n|u64x 0000000000000000 0000000000000001 0000000000000050 000000000000007f 0000000000000080 00000000000000ff 000000000000ffff 00000000fffcfdfe|",
		BD(), V(ceVer), AU64([]uint64{0x00, 0x01, 0x50, 0x7f, 0x80, 0xff, 0xffff, 0xfffcfdfe}), ED())

	assertDecode(t, nil, "c0 |u64 00 01 01 0b101 0b110 0B101 0B110 0o10 0o11 0O10 0O11 0x7f 0x80 0X7f 0X80 0xff 0Xff 100000000|",
		BD(), V(ceVer), AU64([]uint64{0, 1, 1, 5, 6, 5, 6, 8, 9, 8, 9, 127, 128, 127, 128, 255, 255, 100000000}), ED())

	assertDecodeFails(t, "c0 |u64b 10000000000000000000000000000000000000000000000000000000000000000|")
	assertDecodeFails(t, "c0 |u64o 2000000000000000000000|")
	assertDecodeFails(t, "c0 |u64 18446744073709551616|")
	assertDecodeFails(t, "c0 |u64x 10000000000000000|")
}

func TestCTEArrayFloat16(t *testing.T) {
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Float16 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|f16x 1.fep+10 -1.3p-40 1.18p+127 1.18p-126|",
		BD(), V(ceVer), AF16([]uint8{0xff, 0x44, 0x98, 0xab, 0x0c, 0x7f, 0x8c, 0x00}), ED())

	eOpts.DefaultFormats.Array.Float16 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|f16 250 -0.25|",
		BD(), V(ceVer), AF16([]uint8{0x7a, 0x43, 0x80, 0xbe}), ED())

	assertDecode(t, nil, "c0 |f16 0.25 0x4.dp-30|",
		BD(), V(ceVer), AF16([]uint8{0x80, 0x3e, 0x9a, 0x31}), ED())

	assertDecodeFails(t, "c0 |f16 0x1.fep+128|")
	assertDecodeFails(t, "c0 |f16 0x1.fep-127|")
	assertDecodeFails(t, "c0 |f16 0x1.fffffffffffffffffffffffff|")
	assertDecodeFails(t, "c0 |f16 -0x1.fffffffffffffffffffffffff|")
}

func TestCTEArrayFloat32(t *testing.T) {
	// 24 sig bits, 8 exp bits
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Float32 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|f32x 1.fep+10 -1.3p-40 1.111112p+127 1.111112p-126|",
		BD(), V(ceVer), AF32([]float32{0x1.fep+10, -0x1.3p-40, 0x1.111112p+127, 0x1.111112p-126}), ED())

	eOpts.DefaultFormats.Array.Float32 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|f32 1.5e+10 -5.9012e-30|",
		BD(), V(ceVer), AF32([]float32{1.5e+10, -5.9012e-30}), ED())

	assertDecode(t, nil, "c0 |f32 5.5e+10 -0xe.89p+50|",
		BD(), V(ceVer), AF32([]float32{5.5e+10, -0xe.89p+50}), ED())

	assertDecodeFails(t, "c0 |f32 0x1.fep+128|")
	assertDecodeFails(t, "c0 |f32 0x1.fep-127|")
	assertDecodeFails(t, "c0 |f32 0x1.fffffffffffffffffffffffff|")
	assertDecodeFails(t, "c0 |f32 -0x1.fffffffffffffffffffffffff|")
}

func TestCTEArrayFloat64(t *testing.T) {
	// 53 sig bits, 11 exp bits
	eOpts := options.DefaultCTEEncoderOptions()

	eOpts.DefaultFormats.Array.Float64 = options.CTEEncodingFormatHexadecimal
	assertDecodeEncode(t, nil, eOpts, "c0\n|f64x 1.fep+10 -1.3p-40 1.111112p+1023 1.111112p-1022|",
		BD(), V(ceVer), AF64([]float64{0x1.fep+10, -0x1.3p-40, 0x1.111112p+1023, 0x1.111112p-1022}), ED())

	eOpts.DefaultFormats.Array.Float64 = options.CTEEncodingFormatUnset
	assertDecodeEncode(t, nil, eOpts, "c0\n|f64 1.5e+308 1.5e-308|",
		BD(), V(ceVer), AF64([]float64{1.5e+308, 1.5e-308}), ED())

	assertDecodeEncode(t, nil, eOpts, "c0\n|f64 1.5e+10 -5.9012e-30|",
		BD(), V(ceVer), AF64([]float64{1.5e+10, -5.9012e-30}), ED())

	assertDecode(t, nil, "c0 |f64 5.5e+10 -0xe.89p+50|",
		BD(), V(ceVer), AF64([]float64{5.5e+10, -0xe.89p+50}), ED())

	assertDecodeFails(t, "c0 |f64 0x1.fep+1024|")
	assertDecodeFails(t, "c0 |f64 0x1.fep-1023|")
	assertDecodeFails(t, "c0 |f64 0x1.fffffffffffffffffffffffff|")
	assertDecodeFails(t, "c0 |f64 -0x1.fffffffffffffffffffffffff|")
}

func TestCTEArrayUUID(t *testing.T) {
	// TODO: TestCTEArrayUUID
}

func TestCTEArrayBool(t *testing.T) {
	// TODO: TestCTEArrayBool
}

func TestCTEChunked(t *testing.T) {
	assertChunkedStringlike := func(encoded string, startEvent *test.TEvent) {
		assertEncode(t, nil, encoded, BD(), V(ceVer), startEvent, AC(8, false), AD([]byte("abcdefgh")), ED())
		assertEncode(t, nil, encoded, BD(), V(ceVer), startEvent,
			AC(1, true), AD([]byte("a")),
			AC(2, true), AD([]byte("bc")),
			AC(3, true), AD([]byte("def")),
			AC(2, false), AD([]byte("gh")),
			ED())

		assertEncode(t, nil, encoded, BD(), V(ceVer), startEvent,
			AC(1, true), AD([]byte("a")),
			AC(2, true), AD([]byte("bc")),
			AC(3, true), AD([]byte("def")),
			AC(2, true), AD([]byte("gh")),
			AC(0, false), ED())
	}

	assertChunkedByteslike := func(encoded string, startEvent *test.TEvent) {
		assertEncode(t, nil, encoded, BD(), V(ceVer), startEvent, AC(5, false), AD([]byte{0x12, 0x34, 0x56, 0x78, 0x9a}), ED())
		assertEncode(t, nil, encoded, BD(), V(ceVer), startEvent,
			AC(1, true), AD([]byte{0x12}),
			AC(2, true), AD([]byte{0x34, 0x56}),
			AC(2, false), AD([]byte{0x78, 0x9a}),
			ED())
	}

	assertChunkedStringlike("c0\nabcdefgh", SB())
	//TODO: assertChunkedStringlike("c0 `# abcdefgh#", VB())
	assertChunkedStringlike("c0\n|r abcdefgh|", RB())
	assertChunkedStringlike("c0\n|ct abcdefgh|", CTB())
	assertChunkedByteslike("c0\n|cb 12 34 56 78 9a|", CBB())
	assertChunkedByteslike("c0\n|u8x 12 34 56 78 9a|", AU8B())
}

func TestCTEList(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
[]`, BD(), V(ceVer), L(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    123
]`, BD(), V(ceVer), L(), PI(123), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    test
]`, BD(), V(ceVer), L(), S("test"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    -1
    a
    2
    test
    -3
]`, BD(), V(ceVer), L(), NI(1), S("a"), PI(2), S("test"), NI(3), E(), ED())
}

func TestCTEDuplicateEmptySliceInSlice(t *testing.T) {
	sl := []interface{}{}
	v := []interface{}{sl, sl, sl}
	assertMarshalUnmarshal(t, v, `c0
[
    []
    []
    []
]`)
}

func TestCTEMap(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
{}`, BD(), V(ceVer), M(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = 2
}`, BD(), V(ceVer), M(), PI(1), PI(2), E(), ED())
	assertDecode(t, nil, "c0 {  1 = 2 3=4 \t}", BD(), V(ceVer), M(), PI(1), PI(2), PI(3), PI(4), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    na = @na
    1.5 = 1000
}`)

	assertDecode(t, nil, `c0 {email = |r mailto:me@somewhere.com| 1.5 = "a string"}`, BD(), V(ceVer), M(),
		S("email"), RID("mailto:me@somewhere.com"),
		DF(NewDFloat("1.5")), S("a string"),
		E(), ED())

	assertDecodeEncode(t, nil, nil, `c0
{
    a = @inf
    b = 1
}`)
	assertDecodeEncode(t, nil, nil, `c0
{
    a = -@inf
    b = 1
}`)
}

func TestCTEMapBadKVSeparator(t *testing.T) {
	assertDecodeFails(t, "c0 {a:b}")
}

func TestCTEListList(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
[
    []
]`, BD(), V(ceVer), L(), L(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    []
]`, BD(), V(ceVer), L(), PI(1), L(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    []
    1
]`, BD(), V(ceVer), L(), PI(1), L(), E(), PI(1), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    [
        2
    ]
    1
]`, BD(), V(ceVer), L(), PI(1), L(), PI(2), E(), PI(1), E(), ED())
}

func TestCTEListMap(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
[
    {}
]`, BD(), V(ceVer), L(), M(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    {}
]`, BD(), V(ceVer), L(), PI(1), M(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    {}
    1
]`, BD(), V(ceVer), L(), PI(1), M(), E(), PI(1), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    {
        2 = 3
    }
    1
]`, BD(), V(ceVer), L(), PI(1), M(), PI(2), PI(3), E(), PI(1), E(), ED())
}

func TestCTEMapList(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = []
}`, BD(), V(ceVer), M(), PI(1), L(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = [
        2
    ]
    test = [
        1
        2
        3
    ]
}`, BD(), V(ceVer), M(), PI(1), L(), PI(2), E(), S("test"), L(), PI(1), PI(2), PI(3), E(), E(), ED())
}

func TestCTEMapMap(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = {}
}`, BD(), V(ceVer), M(), PI(1), M(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = {
        a = b
    }
    test = {}
}`, BD(), V(ceVer), M(), PI(1), M(), S("a"), S("b"), E(), S("test"), M(), E(), E(), ED())
}

func TestCTEMetadata(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
()a`, BD(), V(ceVer), META(), E(), S("a"), ED())
	assertDecodeEncode(t, nil, nil, `c0
(
    1 = 2
)a`, BD(), V(ceVer), META(), PI(1), PI(2), E(), S("a"), ED())
	assertDecode(t, nil, "c0 (  1 = 2 3=4 \t)a", BD(), V(ceVer), META(), PI(1), PI(2), PI(3), PI(4), E(), S("a"), ED())
}

func TestCTEMarkup(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
<a>`, BD(), V(ceVer), MUP(), S("a"), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a 1=2 3=4>`, BD(), V(ceVer), MUP(), S("a"), PI(1), PI(2), PI(3), PI(4), E(), E(), ED())
	assertDecode(t, nil, `c0
<a,>`, BD(), V(ceVer), MUP(), S("a"), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    a
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("a"), E(), ED())
	assertDecode(t, nil, `c0 <a,a string >`, BD(), V(ceVer), MUP(), S("a"), E(), S("a string"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    <a>
>`, BD(), V(ceVer), MUP(), S("a"), E(), MUP(), S("a"), E(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    a<a>
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("a"), MUP(), S("a"), E(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    <a>
>`, BD(), V(ceVer), MUP(), S("a"), E(), MUP(), S("a"), E(), E(), E(), ED())
	assertDecode(t, nil, `c0 <a 1=2 ,>`, BD(), V(ceVer), MUP(), S("a"), PI(1), PI(2), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a 1=2,
    a
>`, BD(), V(ceVer), MUP(), S("a"), PI(1), PI(2), E(), S("a"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a 1=2,
    <a>
>`, BD(), V(ceVer), MUP(), S("a"), PI(1), PI(2), E(), MUP(), S("a"), E(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a 1=2,
    a <a>
>`, BD(), V(ceVer), MUP(), S("a"), PI(1), PI(2), E(), S("a "), MUP(), S("a"), E(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    ***
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("***"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    /x
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("/x"), E(), ED())

	assertDecodeEncode(t, nil, nil, `c0
<a,
    \\
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("\\"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    \210
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("\u0010"), E(), ED())

	assertDecodeEncode(t, nil, nil, `c0
<a,
    \\
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("\\"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    \<
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("<"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
<a,
    \>
>`, BD(), V(ceVer), MUP(), S("a"), E(), S(">"), E(), ED())
	assertDecode(t, nil, `c0 <a,\*>`, BD(), V(ceVer), MUP(), S("a"), E(), S("*"), E(), ED())
	assertDecode(t, nil, `c0 <a,\/>`, BD(), V(ceVer), MUP(), S("a"), E(), S("/"), E(), ED())

	assertDecodeFails(t, `c0 <a,\y>`)
}

func TestCTEMarkupVerbatimString(t *testing.T) {
	assertDecode(t, nil, `c0 <s, \.## <d></d>##>`)
	assertDecode(t, nil, `c0 <s, \.## /d##>`)
}

func TestCTEMarkupMarkup(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
<a,
    <a>
>`, BD(), V(ceVer), MUP(), S("a"), E(), MUP(), S("a"), E(), E(), E(), ED())
}

func TestCTEMarkupComment(t *testing.T) {
	assertDecode(t, nil, "c0 <a,//blah\n>", BD(), V(ceVer), MUP(), S("a"), E(), CMT(), S("blah"), E(), E(), ED())
	assertDecode(t, nil, "c0 <a,//blah\n a>", BD(), V(ceVer), MUP(), S("a"), E(), CMT(), S("blah"), E(), S("a"), E(), ED())
	assertDecode(t, nil, "c0 <a,a//blah\n a>", BD(), V(ceVer), MUP(), S("a"), E(), S("a"), CMT(), S("blah"), E(), S("a"), E(), ED())

	assertDecode(t, nil, "c0 <a,/*blah*/>", BD(), V(ceVer), MUP(), S("a"), E(), CMT(), S("blah"), E(), E(), ED())
	assertDecode(t, nil, "c0 <a,a/*blah*/>", BD(), V(ceVer), MUP(), S("a"), E(), S("a"), CMT(), S("blah"), E(), E(), ED())
	assertDecode(t, nil, "c0 <a,/*blah*/a>", BD(), V(ceVer), MUP(), S("a"), E(), CMT(), S("blah"), E(), S("a"), E(), ED())

	assertDecode(t, nil, "c0 <a,/*/*blah*/*/>", BD(), V(ceVer), MUP(), S("a"), E(), CMT(), CMT(), S("blah"), E(), E(), E(), ED())
	assertDecode(t, nil, "c0 <a,a/*/*blah*/*/>", BD(), V(ceVer), MUP(), S("a"), E(), S("a"), CMT(), CMT(), S("blah"), E(), E(), E(), ED())
	assertDecode(t, nil, "c0 <a,/*/*blah*/*/a>", BD(), V(ceVer), MUP(), S("a"), E(), CMT(), CMT(), S("blah"), E(), E(), S("a"), E(), ED())

	// TODO: Should it be picking up the extra space between the x and comment?
	assertDecode(t, nil, "c0 <a,x /*blah*/ x>", BD(), V(ceVer), MUP(), S("a"), E(), S("x "), CMT(), S("blah"), E(), S("x"), E(), ED())
}

func TestCTEMapMetadata(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
[
    1
    ()a
]`, BD(), V(ceVer), L(), PI(1), META(), E(), S("a"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = ()a
}`, BD(), V(ceVer), M(), PI(1), META(), E(), S("a"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = {}
}`, BD(), V(ceVer), M(), PI(1), M(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    1 = (){}
}`, BD(), V(ceVer), M(), PI(1), META(), E(), M(), E(), E(), ED())

	assertDecodeEncode(t, nil, nil, `c0
{
    ()()1 = ()()a
}`, BD(), V(ceVer), M(), META(), E(), META(), E(), PI(1), META(), E(), META(), E(), S("a"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    ()()1 = ()(){}
}`, BD(), V(ceVer), M(), META(), E(), META(), E(), PI(1), META(), E(), META(), E(), M(), E(), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    ()()1 = ()()[]
}`, BD(), V(ceVer), M(), META(), E(), META(), E(), PI(1), META(), E(), META(), E(), L(), E(), E(), ED())

	assertDecodeEncode(t, nil, nil, `c0
(
    x = y
){
    (
        x = y
    )1 = (
        x = y
    )(
        x = y
    ){
        a = b
    }
}`, BD(), V(ceVer),
		META(), S("x"), S("y"), E(), M(),
		META(), S("x"), S("y"), E(), PI(1),
		META(), S("x"), S("y"), E(),
		META(), S("x"), S("y"), E(),
		M(), S("a"), S("b"), E(), E(), ED())
}

func TestCTENamed(t *testing.T) {
	assertDecodeEncode(t, nil, nil, "c0\n@na", BD(), V(ceVer), NA(), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@nan", BD(), V(ceVer), NAN(), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@snan", BD(), V(ceVer), SNAN(), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@inf", BD(), V(ceVer), F(math.Inf(1)), ED())
	assertDecodeEncode(t, nil, nil, "c0\n-@inf", BD(), V(ceVer), F(math.Inf(-1)), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@false", BD(), V(ceVer), FF(), ED())
	assertDecodeEncode(t, nil, nil, "c0\n@true", BD(), V(ceVer), TT(), ED())
}

func TestCTEMarker(t *testing.T) {
	assertDecodeFails(t, `c0 &2`)
	assertDecode(t, nil, `c0 &1:string`, BD(), V(ceVer), MARK(), PI(1), S("string"), ED())
	assertDecode(t, nil, `c0 &a:string`, BD(), V(ceVer), MARK(), S("a"), S("string"), ED())
	assertDecodeFails(t, `c0 & 1:string`)
	assertDecodeFails(t, `c0 &1 string`)
	assertDecodeFails(t, `c0 &1string`)
	assertDecodeFails(t, `c0 &rgnsekfrnsekrgfnskergnslekrgnslergselrgblserfbserfbvsekrskfrvbskerfbksefbskerbfserbfrbksuerfbsekjrfbdjfgbsdjfgbsdfgbsdjkhfg`)
	assertDecodeFails(t, `c0 &100000000000000000000000000000000000000000000000`)
}

func TestCTEReference(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
[
    &2:aaaaa
    $2
]`, BD(), V(ceVer), L(), MARK(), PI(2), S("aaaaa"), REF(), PI(2), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    &a:aaaaa
    $a
]`, BD(), V(ceVer), L(), MARK(), S("a"), S("aaaaa"), REF(), S("a"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
[
    &a:aaaaa
    a
]`, BD(), V(ceVer), L(), MARK(), S("a"), S("aaaaa"), S("a"), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
$|r http://x.y|`, BD(), V(ceVer), REF(), RID("http://x.y"), ED())
	assertDecodeFails(t, `c0 $ 1`)
}

func TestCTEMarkerReference(t *testing.T) {
	assertDecode(t, nil, `c0 [&2:testing $2]`, BD(), V(ceVer), L(), MARK(), PI(2), S("testing"), REF(), PI(2), E(), ED())
	assertDecodeEncode(t, nil, nil, `c0
{
    first = &1:1000
    second = $1
}`)
}

func TestCTEComment(t *testing.T) {
	// TODO: Better comment formatting
	assertDecodeEncode(t, nil, nil, `c0
{
    a = @inf
    /* test */
    b = 1
}`)
}

func TestCTECommentSingleLine(t *testing.T) {
	assertDecodeFails(t, "c0 //")
	assertDecode(t, nil, "c0 //\n1", BD(), V(ceVer), CMT(), E(), PI(1), ED())
	assertDecode(t, nil, "c0 //\r\n1", BD(), V(ceVer), CMT(), E(), PI(1), ED())
	assertDecodeFails(t, "c0 // ")
	assertDecode(t, nil, "c0 // \n1", BD(), V(ceVer), CMT(), E(), PI(1), ED())
	assertDecode(t, nil, "c0 // \r\n1", BD(), V(ceVer), CMT(), E(), PI(1), ED())
	assertDecodeFails(t, "c0 //a")
	assertDecode(t, nil, "c0 //a\n1", BD(), V(ceVer), CMT(), S("a"), E(), PI(1), ED())
	assertDecode(t, nil, "c0 //a\r\n1", BD(), V(ceVer), CMT(), S("a"), E(), PI(1), ED())
	assertDecode(t, nil, "c0 // This is a comment\n1", BD(), V(ceVer), CMT(), S("This is a comment"), E(), PI(1), ED())
	assertDecodeFails(t, "c0 /-\n")
}

func TestCTECommentMultiline(t *testing.T) {
	assertDecode(t, nil, "c0 /**/1", BD(), V(ceVer), CMT(), E(), PI(1), ED())
	assertDecode(t, nil, "c0 /**/1", BD(), V(ceVer), CMT(), E(), PI(1), ED())
	assertDecode(t, nil, "c0 /* This is a comment */1", BD(), V(ceVer), CMT(), S("This is a comment"), E(), PI(1), ED())
	assertDecode(t, nil, "c0 /*This is a comment*/1", BD(), V(ceVer), CMT(), S("This is a comment"), E(), PI(1), ED())
}

func TestCTECommentMultilineNested(t *testing.T) {
	assertDecode(t, nil, "c0 /*/**/*/1", BD(), V(ceVer), CMT(), CMT(), E(), E(), PI(1), ED())
	assertDecode(t, nil, "c0 /*/**/*/1", BD(), V(ceVer), CMT(), CMT(), E(), E(), PI(1), ED())
	assertDecode(t, nil, "c0 /* /**/ */1", BD(), V(ceVer), CMT(), CMT(), E(), E(), PI(1), ED())
	assertDecode(t, nil, "c0  /* before/* mid */ after*/1  ", BD(), V(ceVer), CMT(), S("before"), CMT(), S("mid"), E(), S("after"), E(), PI(1), ED())
}

func TestCTECommentAfterValue(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
[
    a
    /**/
]`, BD(), V(ceVer), L(), S("a"), CMT(), E(), E(), ED())
}

func TestCTEComplexComment(t *testing.T) {
	document := []byte(`c0
/**/ ( /**/ a= /**/ b /**/ ) /**/
<a,
    /**/
    <b>
>`)

	expected := `c0
/**/
(
    /**/
    a = /**/
    b
    /**/
)/**/
<a,
    /**/
    <b>
>`

	encoded := &bytes.Buffer{}
	encOpts := options.DefaultCTEEncoderOptions()
	encOpts.Indent = "    "
	encoder := NewEncoder(encOpts)
	encoder.PrepareToEncode(encoded)
	decoder := NewDecoder(nil)
	err := decoder.Decode(bytes.NewBuffer(document), encoder)
	if err != nil {
		t.Error(err)
		return
	}

	actual := string(encoded.Bytes())
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestCTECommentFollowing(t *testing.T) {
	assertDecode(t, nil, "c0 {a=b/**/}", BD(), V(ceVer), M(), S("a"), S("b"), CMT(), E(), E(), ED())
	assertDecode(t, nil, "c0 {a=2/**/}", BD(), V(ceVer), M(), S("a"), PI(2), CMT(), E(), E(), ED())
	assertDecode(t, nil, "c0 {a=-2/**/}", BD(), V(ceVer), M(), S("a"), NI(2), CMT(), E(), E(), ED())
	// TODO: All other bare values: float, date/time, etc
	assertDecode(t, nil, "c0 {a=1.5/**/}", BD(), V(ceVer), M(), S("a"), DF(NewDFloat("1.5")), CMT(), E(), E(), ED())
	// TODO: Also test for //
}

func TestCTECommentPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
{
    a = b
    /**/
}`, BD(), V(ceVer), M(), S("a"), S("b"), CMT(), E(), E(), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
/**/
1`, BD(), V(ceVer), CMT(), E(), PI(1), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
/* a */
1`, BD(), V(ceVer), CMT(), S("a"), E(), PI(1), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
/* /**/ */
1`, BD(), V(ceVer), CMT(), CMT(), E(), E(), PI(1), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
/* /* a */ */
1`, BD(), V(ceVer), CMT(), CMT(), S("a"), E(), E(), PI(1), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
/**/
a`, BD(), V(ceVer), CMT(), E(), S("a"), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[
    /* xyz */
    a
]`, BD(), V(ceVer), L(), CMT(), S("xyz"), E(), S("a"), E(), ED())
}

func TestCTEMarkupPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
<a>`, BD(), V(ceVer), MUP(), S("a"), E(), E(), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
<a x=1>`, BD(), V(ceVer), MUP(), S("a"), S("x"), PI(1), E(), E(), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
<a,
    aaa
>`, BD(), V(ceVer), MUP(), S("a"), E(), S("aaa"), E(), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
<a x=y,
    aaa
>`, BD(), V(ceVer), MUP(), S("a"), S("x"), S("y"), E(), S("aaa"), E(), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
<a x=y z=1,
    aaa
>`, BD(), V(ceVer), MUP(), S("a"), S("x"), S("y"), S("z"), PI(1), E(), S("aaa"), E(), ED())
}

func TestCTEPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, "c0\n1", BD(), V(ceVer), PI(1), ED())
}

func TestCTEListPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	// Empty 1 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[]`, BD(), V(ceVer), L(), E(), ED())

	// Empty 2 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[
    []
]`, BD(), V(ceVer), L(), L(), E(), E(), ED())
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[
    []
    []
]`, BD(), V(ceVer), L(), L(), E(), L(), E(), E(), ED())

	// 1 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[
    1
    2
]`, BD(), V(ceVer), L(), PI(1), PI(2), E(), ED())

	// 2 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[
    [
        1
        2
    ]
    [
        3
        4
    ]
]`, BD(), V(ceVer), L(), L(), PI(1), PI(2), E(), L(), PI(3), PI(4), E(), E(), ED())
}

func TestCTEMapPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	// // Empty 1 level
	assertDecodeEncode(t, nil, opts, `c0
{}`, BD(), V(ceVer), M(), E(), ED())

	// Empty 2 level
	assertDecodeEncode(t, nil, opts, `c0
{
    a = {}
}`, BD(), V(ceVer), M(), S("a"), M(), E(), E(), ED())
	assertDecodeEncode(t, nil, opts, `c0
{
    a = {}
    b = {}
}`, BD(), V(ceVer), M(), S("a"), M(), E(), S("b"), M(), E(), E(), ED())

	// 1 level
	assertDecodeEncode(t, nil, opts, `c0
{
    1 = 2
}`, BD(), V(ceVer), M(), PI(1), PI(2), E(), ED())

	// 2 level
	assertDecodeEncode(t, nil, opts, `c0
{
    a = {
        1 = 2
        3 = 4
    }
    b = {
        5 = 6
        7 = 8
    }
}`, BD(), V(ceVer), M(), S("a"), M(), PI(1), PI(2), PI(3), PI(4), E(), S("b"), M(), PI(5), PI(6), PI(7), PI(8), E(), E(), ED())
}

func TestCTEMetadataPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	// Empty 1 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
()aa`, BD(), V(ceVer), META(), E(), S("aa"), ED())

	// Empty 2 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
(
    a = ()1
    b = ()2
)aa`, BD(), V(ceVer), META(), S("a"), META(), E(), PI(1), S("b"), META(), E(), PI(2), E(), S("aa"), ED())

	// 1 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
(
    1 = 2
)aa`, BD(), V(ceVer), META(), PI(1), PI(2), E(), S("aa"), ED())

	// 2 level
	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
(
    a = (
        1 = 2
        3 = 4
    )x
    b = (
        5 = 6
        7 = 8
    )y
)aa`, BD(), V(ceVer), META(),
		S("a"), META(), PI(1), PI(2), PI(3), PI(4), E(), S("x"),
		S("b"), META(), PI(5), PI(6), PI(7), PI(8), E(), S("y"),
		E(), S("aa"), ED())
}

func TestCTEArrayPretty(t *testing.T) {
	opts := options.DefaultCTEEncoderOptions()

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
|u8x 22 33|`, BD(), V(ceVer), AU8([]uint8{0x22, 0x33}), ED())

	opts.Indent = "    "
	assertDecodeEncode(t, nil, opts, `c0
[
    |u8x 22 33|
    |u8x 66 77|
]`, BD(), V(ceVer), L(), AU8([]uint8{0x22, 0x33}), AU8([]uint8{0x66, 0x77}), E(), ED())
}

func TestCTEMarkupVerbatimPretty(t *testing.T) {
	assertDecode(t, nil, `c0 <blah, \.# aaa #>`,
		BD(), V(ceVer), MUP(), S("blah"), E(), S("aaa"), E(), ED())
}

func TestCTEBufferEdge(t *testing.T) {
	assertDecode(t, nil, `c0
{
     1  = <a,
            <b,
               <c, `+"`"+`##                       ##>
                         >
                       >
}
`)
}

func TestCTEBufferEdge2(t *testing.T) {
	assertDecode(t, nil, `c0
{
    x  = <a,
                     <b,
                             <c, `+"`"+`##                     ##>
                           >
                       >
}
`)
}

func TestCTEComplexExample(t *testing.T) {
	assertDecode(t, nil, `c0
// Metadata: _ct is the creation time
(_ct = 2019-9-1/22:14:01)
{
    /* /* Nested comments are allowed */ */
    // There are no commas in maps and lists
    (info = "something interesting about a_list")
    a_list           = [1 2 "a string"]
    map              = {2=two 3=3000 1=one}
    string           = "A string value"
    boolean          = @true
    "binary int"     = -0b10001011
    "octal int"      = 0o644
    "regular int"    = -10000000
    "hex int"        = 0xfffe0001
    "decimal float"  = -14.125
    "hex float"      = 0x5.1ec4p20
    uuid             = @f1ce4567-e89b-12d3-a456-426655440000
    date             = 2019-7-1
    time             = 18:04:00.940231541/E/Prague
    timestamp        = 2010-7-15/13:28:15.415942344/Z
    na               = @na
    bytes            = |u8x 10 ff 38 9a dd 00 4f 4f 91|
    url              = |r https://example.com/|
    email            = |r mailto:me@somewhere.com|
    1.5              = "Keys don't have to be strings"
    long-string      = "\.ZZZ
A backtick induces verbatim processing, which in this case will continue
until three Z characters are encountered, similar to how here documents in
bash work.
You can put anything in here, including double-quote ("), or even more
backticks (`+"`"+`). Verbatim processing stops at the end sequence, which in this
case is three Z characters, specified earlier as a sentinel.ZZZ"
    marked_object    = &tag1:{
                                description = "This map will be referenced later using $tag1"
                                value = -@inf
                                child_elements = @na
                                recursive = $tag1
                            }
    ref1             = $tag1
    ref2             = $tag1
    outside_ref      = $|r https://somewhere.else.com/path/to/document.cte#some_tag|
    // The markup type is good for presentation data
    html_compatible  = <html xmlns=|r http://www.w3.org/1999/xhtml| "xml:lang"=en ,
                         <body,
                           Please choose from the following widgets:
                           <div id=parent style=normal ref-id=1 ,
                             /* Here we use a backtick to induce verbatim processing.
                              * In this case, "##" is chosen as the ending sequence
                              */
                             <script, \.##
                               document.getElementById('parent').insertAdjacentHTML('beforeend',
                                  '<div id="idChild"> content </div>'),
                             ##>
                           >
                         >
                       >
}
`)
}

func TestCTEEncodeDecodeExample(t *testing.T) {
	document := `c0
/* _ct is the creation time, in this case referring to the entire document */
(
    _ct = 2019-09-01/22:14:01
){
    /* Comments look very C-like, except: /* Nested comments are allowed! */ */
    /* Notice that there are no commas in maps and lists */
    (
        info = "something interesting about a_list"
    )a_list = [
        1
        2
        "a string"
    ]
    map = {
        2 = two
        3 = 3000
        1 = one
    }
    string = "A string value"
    boolean = @true
    "regular int" = -10000000
    "decimal float" = -14.125
    uuid = @f1ce4567-e89b-12d3-a456-426655440000
    date = 2019-07-01
    time = 18:04:00.940231541/E/Prague
    timestamp = 2010-07-15/13:28:15.415942344/Z
    na = @na
    bytes = |u8x 10 ff 38 9a dd 00 4f 4f 91|
    url = |r https://example.com/|
    email = |r mailto:me@somewhere.com|
    1.5 = "Keys don't have to be strings"
    marked_object = &tag1:{
        description = "This map will be referenced later using $tag1"
        value = -@inf
        child_elements = @na
        recursive = $tag1
    }
    ref1 = $tag1
    ref2 = $tag1
    outside_ref = $|r https://somewhere.else.com/path/to/document.cte#some_tag|
    // The markup type is good for presentation data
    html_compatible = <html xmlns=|r http://www.w3.org/1999/xhtml| "xml:lang"=en,
        <body,
            Please choose from the following widgets:
            <div id=parent style=normal ref-id=1,
                /* Here we use a backtick to induce verbatim processing.
                 * In this case, "#" is chosen as the ending sequence */
            >
        >
    >
}`

	expected := `c0
/* _ct is the creation time, in this case referring to the entire document */
(
    _ct = 2019-09-01/22:14:01
){
    /* Comments look very C-like, except: /* Nested comments are allowed! */ */
    /* Notice that there are no commas in maps and lists */
    (
        info = "something interesting about a_list"
    )a_list = [
        1
        2
        "a string"
    ]
    map = {
        2 = two
        3 = 3000
        1 = one
    }
    string = "A string value"
    boolean = @true
    "regular int" = -10000000
    "decimal float" = -14.125
    uuid = @f1ce4567-e89b-12d3-a456-426655440000
    date = 2019-07-01
    time = 18:04:00.940231541/Europe/Prague
    timestamp = 2010-07-15/13:28:15.415942344
    na = @na
    bytes = |u8x 10 ff 38 9a dd 00 4f 4f 91|
    url = |r https://example.com/|
    email = |r mailto:me@somewhere.com|
    1.5 = "Keys don't have to be strings"
    marked_object = &tag1:{
        description = "This map will be referenced later using $tag1"
        value = -@inf
        child_elements = @na
        recursive = $tag1
    }
    ref1 = $tag1
    ref2 = $tag1
    outside_ref = $|r https://somewhere.else.com/path/to/document.cte#some_tag|
    /* The markup type is good for presentation data */
    html_compatible = <html xmlns=|r http://www.w3.org/1999/xhtml| "xml:lang"=en,
        <body,
            Please choose from the following widgets: <div id=parent style=normal ref-id=1,
                /* Here we use a backtick to induce verbatim processing.
                 * In this case, "#" is chosen as the ending sequence */
            >
        >
    >
}`

	encoded := &bytes.Buffer{}
	encOpts := options.DefaultCTEEncoderOptions()
	encOpts.Indent = "    "
	encoder := NewEncoder(encOpts)
	encoder.PrepareToEncode(encoded)
	decoder := NewDecoder(nil)
	err := decoder.Decode(bytes.NewBuffer([]byte(document)), encoder)
	if err != nil {
		t.Error(err)
		return
	}

	actual := string(encoded.Bytes())
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestMapValueComment(t *testing.T) {
	assertEncode(t, nil, `c0
{
    1 = /**/
    1
}`, BD(), V(ceVer), M(), PI(1), CMT(), E(), PI(1), E(), ED())
}

func TestEmptyDocument(t *testing.T) {
	assertDecodeEncode(t, nil, nil, `c0
@na`, BD(), V(ceVer), NA(), ED())
}
