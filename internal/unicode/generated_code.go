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

// Generated by github.com/kstenerud/go-concise-encoding/codegen
  // DO NOT EDIT
  // IF THIS LINE SHOWS UP IN THE GIT DIFF, THIS FILE HAS BEEN EDITED

package unicode

type CharProperty uint16

const (
	NumeralOrLookalike CharProperty = 1 << iota
	LowSymbolOrLookalike
	Whitespace
	Control
	TabReturnNewline
	QuotedTextDelimiter
	ArrayDelimiter
	MarkupDelimiter
	MarkerIDSafe
	NoProperties CharProperty = 0
)

var charProperties = map[rune]CharProperty{
	/*     */ 0x00: Control,
	/*     */ 0x01: Control,
	/*     */ 0x02: Control,
	/*     */ 0x03: Control,
	/*     */ 0x04: Control,
	/*     */ 0x05: Control,
	/*     */ 0x06: Control,
	/*     */ 0x07: Control,
	/*     */ 0x08: Control,
	/* \t  */ 0x09: Whitespace | TabReturnNewline,
	/* \n  */ 0x0a: Whitespace | TabReturnNewline,
	/*     */ 0x0b: Control,
	/*     */ 0x0c: Control,
	/* \r  */ 0x0d: Whitespace | TabReturnNewline,
	/*     */ 0x0e: Control,
	/*     */ 0x0f: Control,
	/*     */ 0x10: Control,
	/*     */ 0x11: Control,
	/*     */ 0x12: Control,
	/*     */ 0x13: Control,
	/*     */ 0x14: Control,
	/*     */ 0x15: Control,
	/*     */ 0x16: Control,
	/*     */ 0x17: Control,
	/*     */ 0x18: Control,
	/*     */ 0x19: Control,
	/*     */ 0x1a: Control,
	/*     */ 0x1b: Control,
	/*     */ 0x1c: Control,
	/*     */ 0x1d: Control,
	/*     */ 0x1e: Control,
	/*     */ 0x1f: Control,
	/* [ ] */ 0x20: Whitespace,
	/* [!] */ 0x21: LowSymbolOrLookalike,
	/* ["] */ 0x22: LowSymbolOrLookalike | QuotedTextDelimiter,
	/* [#] */ 0x23: LowSymbolOrLookalike,
	/* [$] */ 0x24: LowSymbolOrLookalike,
	/* [%] */ 0x25: LowSymbolOrLookalike,
	/* [&] */ 0x26: LowSymbolOrLookalike,
	/* ['] */ 0x27: LowSymbolOrLookalike,
	/* [(] */ 0x28: LowSymbolOrLookalike,
	/* [)] */ 0x29: LowSymbolOrLookalike,
	/* [*] */ 0x2a: LowSymbolOrLookalike,
	/* [+] */ 0x2b: LowSymbolOrLookalike,
	/* [,] */ 0x2c: LowSymbolOrLookalike,
	/* [-] */ 0x2d: NoProperties,
	/* [.] */ 0x2e: NoProperties,
	/* [0] */ 0x30: NumeralOrLookalike | MarkerIDSafe,
	/* [1] */ 0x31: NumeralOrLookalike | MarkerIDSafe,
	/* [2] */ 0x32: NumeralOrLookalike | MarkerIDSafe,
	/* [3] */ 0x33: NumeralOrLookalike | MarkerIDSafe,
	/* [4] */ 0x34: NumeralOrLookalike | MarkerIDSafe,
	/* [5] */ 0x35: NumeralOrLookalike | MarkerIDSafe,
	/* [6] */ 0x36: NumeralOrLookalike | MarkerIDSafe,
	/* [7] */ 0x37: NumeralOrLookalike | MarkerIDSafe,
	/* [8] */ 0x38: NumeralOrLookalike | MarkerIDSafe,
	/* [9] */ 0x39: NumeralOrLookalike | MarkerIDSafe,
	/* [:] */ 0x3a: NoProperties,
	/* [;] */ 0x3b: LowSymbolOrLookalike,
	/* [<] */ 0x3c: LowSymbolOrLookalike | MarkupDelimiter,
	/* [=] */ 0x3d: LowSymbolOrLookalike,
	/* [>] */ 0x3e: LowSymbolOrLookalike | MarkupDelimiter,
	/* [?] */ 0x3f: LowSymbolOrLookalike,
	/* [@] */ 0x40: LowSymbolOrLookalike,
	/* [A] */ 0x41: MarkerIDSafe,
	/* [B] */ 0x42: MarkerIDSafe,
	/* [C] */ 0x43: MarkerIDSafe,
	/* [D] */ 0x44: MarkerIDSafe,
	/* [E] */ 0x45: MarkerIDSafe,
	/* [F] */ 0x46: MarkerIDSafe,
	/* [G] */ 0x47: MarkerIDSafe,
	/* [H] */ 0x48: MarkerIDSafe,
	/* [I] */ 0x49: MarkerIDSafe,
	/* [J] */ 0x4a: MarkerIDSafe,
	/* [K] */ 0x4b: MarkerIDSafe,
	/* [L] */ 0x4c: MarkerIDSafe,
	/* [M] */ 0x4d: MarkerIDSafe,
	/* [N] */ 0x4e: MarkerIDSafe,
	/* [O] */ 0x4f: MarkerIDSafe,
	/* [P] */ 0x50: MarkerIDSafe,
	/* [Q] */ 0x51: MarkerIDSafe,
	/* [R] */ 0x52: MarkerIDSafe,
	/* [S] */ 0x53: MarkerIDSafe,
	/* [T] */ 0x54: MarkerIDSafe,
	/* [U] */ 0x55: MarkerIDSafe,
	/* [V] */ 0x56: MarkerIDSafe,
	/* [W] */ 0x57: MarkerIDSafe,
	/* [X] */ 0x58: MarkerIDSafe,
	/* [Y] */ 0x59: MarkerIDSafe,
	/* [Z] */ 0x5a: MarkerIDSafe,
	/* [[] */ 0x5b: LowSymbolOrLookalike,
	/* [\] */ 0x5c: LowSymbolOrLookalike | QuotedTextDelimiter | ArrayDelimiter | MarkupDelimiter,
	/* []] */ 0x5d: LowSymbolOrLookalike,
	/* [^] */ 0x5e: LowSymbolOrLookalike,
	/* [_] */ 0x5f: MarkerIDSafe,
	/* [`] */ 0x60: LowSymbolOrLookalike | MarkupDelimiter,
	/* [a] */ 0x61: MarkerIDSafe,
	/* [b] */ 0x62: MarkerIDSafe,
	/* [c] */ 0x63: MarkerIDSafe,
	/* [d] */ 0x64: MarkerIDSafe,
	/* [e] */ 0x65: MarkerIDSafe,
	/* [f] */ 0x66: MarkerIDSafe,
	/* [g] */ 0x67: MarkerIDSafe,
	/* [h] */ 0x68: MarkerIDSafe,
	/* [i] */ 0x69: MarkerIDSafe,
	/* [j] */ 0x6a: MarkerIDSafe,
	/* [k] */ 0x6b: MarkerIDSafe,
	/* [l] */ 0x6c: MarkerIDSafe,
	/* [m] */ 0x6d: MarkerIDSafe,
	/* [n] */ 0x6e: MarkerIDSafe,
	/* [o] */ 0x6f: MarkerIDSafe,
	/* [p] */ 0x70: MarkerIDSafe,
	/* [q] */ 0x71: MarkerIDSafe,
	/* [r] */ 0x72: MarkerIDSafe,
	/* [s] */ 0x73: MarkerIDSafe,
	/* [t] */ 0x74: MarkerIDSafe,
	/* [u] */ 0x75: MarkerIDSafe,
	/* [v] */ 0x76: MarkerIDSafe,
	/* [w] */ 0x77: MarkerIDSafe,
	/* [x] */ 0x78: MarkerIDSafe,
	/* [y] */ 0x79: MarkerIDSafe,
	/* [z] */ 0x7a: MarkerIDSafe,
	/* [{] */ 0x7b: LowSymbolOrLookalike,
	/* [|] */ 0x7c: LowSymbolOrLookalike | ArrayDelimiter,
	/* [}] */ 0x7d: LowSymbolOrLookalike,
	/* [~] */ 0x7e: LowSymbolOrLookalike,
	/*     */ 0x7f: Control,
	/*     */ 0x80: Control,
	/*     */ 0x81: Control,
	/*     */ 0x82: Control,
	/*     */ 0x83: Control,
	/*     */ 0x84: Control,
	/*     */ 0x85: Control,
	/*     */ 0x86: Control,
	/*     */ 0x87: Control,
	/*     */ 0x88: Control,
	/*     */ 0x89: Control,
	/*     */ 0x8a: Control,
	/*     */ 0x8b: Control,
	/*     */ 0x8c: Control,
	/*     */ 0x8d: Control,
	/*     */ 0x8e: Control,
	/*     */ 0x8f: Control,
	/*     */ 0x90: Control,
	/*     */ 0x91: Control,
	/*     */ 0x92: Control,
	/*     */ 0x93: Control,
	/*     */ 0x94: Control,
	/*     */ 0x95: Control,
	/*     */ 0x96: Control,
	/*     */ 0x97: Control,
	/*     */ 0x98: Control,
	/*     */ 0x99: Control,
	/*     */ 0x9a: Control,
	/*     */ 0x9b: Control,
	/*     */ 0x9c: Control,
	/*     */ 0x9d: Control,
	/*     */ 0x9e: Control,
	/*     */ 0x9f: Control,
	/* [ ] */ 0xa0: Whitespace,
	/* [¦] */ 0xa6: LowSymbolOrLookalike,
	/*     */ 0xad: Control,
	/* [´] */ 0xb4: LowSymbolOrLookalike,
	/*     */ 0x600: Control,
	/*     */ 0x601: Control,
	/*     */ 0x602: Control,
	/*     */ 0x603: Control,
	/*     */ 0x604: Control,
	/*     */ 0x605: Control,
	/*     */ 0x61c: Control,
	/*     */ 0x6dd: Control,
	/*     */ 0x70f: Control,
	/*     */ 0x8e2: Control,
	/*     */ 0x180e: Control,
	/* [ ] */ 0x2000: Whitespace,
	/* [ ] */ 0x2001: Whitespace,
	/* [ ] */ 0x2002: Whitespace,
	/* [ ] */ 0x2003: Whitespace,
	/* [ ] */ 0x2004: Whitespace,
	/* [ ] */ 0x2005: Whitespace,
	/* [ ] */ 0x2006: Whitespace,
	/* [ ] */ 0x2007: Whitespace,
	/* [ ] */ 0x2008: Whitespace,
	/* [ ] */ 0x2009: Whitespace,
	/* [ ] */ 0x200a: Whitespace,
	/*     */ 0x200b: Control,
	/*     */ 0x200c: Control,
	/*     */ 0x200d: Control,
	/*     */ 0x200e: Control,
	/*     */ 0x200f: Control,
	/* [‘] */ 0x2018: LowSymbolOrLookalike,
	/* [’] */ 0x2019: LowSymbolOrLookalike,
	/* [‚] */ 0x201a: LowSymbolOrLookalike,
	/* [‛] */ 0x201b: LowSymbolOrLookalike,
	/* [“] */ 0x201c: LowSymbolOrLookalike,
	/* [”] */ 0x201d: LowSymbolOrLookalike,
	/* [„] */ 0x201e: LowSymbolOrLookalike,
	/* [‟] */ 0x201f: LowSymbolOrLookalike,
	/*     */ 0x2028: Control,
	/*     */ 0x2029: Control,
	/*     */ 0x202a: Control,
	/*     */ 0x202b: Control,
	/*     */ 0x202c: Control,
	/*     */ 0x202d: Control,
	/*     */ 0x202e: Control,
	/* [ ] */ 0x202f: Whitespace,
	/* [′] */ 0x2032: LowSymbolOrLookalike,
	/* [″] */ 0x2033: LowSymbolOrLookalike,
	/* [‴] */ 0x2034: LowSymbolOrLookalike,
	/* [‵] */ 0x2035: LowSymbolOrLookalike,
	/* [‶] */ 0x2036: LowSymbolOrLookalike,
	/* [‷] */ 0x2037: LowSymbolOrLookalike,
	/* [‹] */ 0x2039: LowSymbolOrLookalike,
	/* [›] */ 0x203a: LowSymbolOrLookalike,
	/* [‼] */ 0x203c: LowSymbolOrLookalike,
	/* [⁇] */ 0x2047: LowSymbolOrLookalike,
	/* [⁈] */ 0x2048: LowSymbolOrLookalike,
	/* [⁉] */ 0x2049: LowSymbolOrLookalike,
	/* [⁎] */ 0x204e: LowSymbolOrLookalike,
	/* [⁒] */ 0x2052: LowSymbolOrLookalike,
	/* [⁕] */ 0x2055: LowSymbolOrLookalike,
	/* [⁗] */ 0x2057: LowSymbolOrLookalike,
	/* [ ] */ 0x205f: Whitespace,
	/*     */ 0x2060: Control,
	/*     */ 0x2061: Control,
	/*     */ 0x2062: Control,
	/*     */ 0x2063: Control,
	/*     */ 0x2064: Control,
	/*     */ 0x2066: Control,
	/*     */ 0x2067: Control,
	/*     */ 0x2068: Control,
	/*     */ 0x2069: Control,
	/*     */ 0x206a: Control,
	/*     */ 0x206b: Control,
	/*     */ 0x206c: Control,
	/*     */ 0x206d: Control,
	/*     */ 0x206e: Control,
	/*     */ 0x206f: Control,
	/* [∗] */ 0x2217: LowSymbolOrLookalike,
	/* [∣] */ 0x2223: LowSymbolOrLookalike,
	/* [∼] */ 0x223c: LowSymbolOrLookalike,
	/* [⎜] */ 0x239c: LowSymbolOrLookalike,
	/* [⎟] */ 0x239f: LowSymbolOrLookalike,
	/* [⎢] */ 0x23a2: LowSymbolOrLookalike,
	/* [⎥] */ 0x23a5: LowSymbolOrLookalike,
	/* [⎪] */ 0x23aa: LowSymbolOrLookalike,
	/* [⎸] */ 0x23b8: LowSymbolOrLookalike,
	/* [⎹] */ 0x23b9: LowSymbolOrLookalike,
	/* [⏐] */ 0x23d0: LowSymbolOrLookalike,
	/* [⏽] */ 0x23fd: LowSymbolOrLookalike,
	/* [　] */ 0x3000: Whitespace,
	/* [︐] */ 0xfe10: LowSymbolOrLookalike,
	/* [︑] */ 0xfe11: LowSymbolOrLookalike,
	/* [︔] */ 0xfe14: LowSymbolOrLookalike,
	/* [︕] */ 0xfe15: LowSymbolOrLookalike,
	/* [︖] */ 0xfe16: LowSymbolOrLookalike,
	/* [︱] */ 0xfe31: LowSymbolOrLookalike,
	/* [︳] */ 0xfe33: LowSymbolOrLookalike,
	/* [﹅] */ 0xfe45: LowSymbolOrLookalike,
	/* [﹆] */ 0xfe46: LowSymbolOrLookalike,
	/* [﹐] */ 0xfe50: LowSymbolOrLookalike,
	/* [﹑] */ 0xfe51: LowSymbolOrLookalike,
	/* [﹒] */ 0xfe52: NoProperties,
	/* [﹔] */ 0xfe54: LowSymbolOrLookalike,
	/* [﹕] */ 0xfe55: NoProperties,
	/* [﹖] */ 0xfe56: LowSymbolOrLookalike,
	/* [﹗] */ 0xfe57: LowSymbolOrLookalike,
	/* [﹘] */ 0xfe58: NoProperties,
	/* [﹙] */ 0xfe59: LowSymbolOrLookalike,
	/* [﹚] */ 0xfe5a: LowSymbolOrLookalike,
	/* [﹛] */ 0xfe5b: LowSymbolOrLookalike,
	/* [﹜] */ 0xfe5c: LowSymbolOrLookalike,
	/* [﹝] */ 0xfe5d: LowSymbolOrLookalike,
	/* [﹞] */ 0xfe5e: LowSymbolOrLookalike,
	/* [﹟] */ 0xfe5f: LowSymbolOrLookalike,
	/* [﹠] */ 0xfe60: LowSymbolOrLookalike,
	/* [﹡] */ 0xfe61: LowSymbolOrLookalike,
	/* [﹢] */ 0xfe62: LowSymbolOrLookalike,
	/* [﹣] */ 0xfe63: NoProperties,
	/* [﹤] */ 0xfe64: LowSymbolOrLookalike,
	/* [﹥] */ 0xfe65: LowSymbolOrLookalike,
	/* [﹦] */ 0xfe66: LowSymbolOrLookalike,
	/* [﹨] */ 0xfe68: LowSymbolOrLookalike,
	/* [﹩] */ 0xfe69: LowSymbolOrLookalike,
	/* [﹪] */ 0xfe6a: LowSymbolOrLookalike,
	/* [﹫] */ 0xfe6b: LowSymbolOrLookalike,
	/*     */ 0xfeff: Control,
	/* [！] */ 0xff01: LowSymbolOrLookalike,
	/* [＂] */ 0xff02: LowSymbolOrLookalike,
	/* [＃] */ 0xff03: LowSymbolOrLookalike,
	/* [＄] */ 0xff04: LowSymbolOrLookalike,
	/* [％] */ 0xff05: LowSymbolOrLookalike,
	/* [＆] */ 0xff06: LowSymbolOrLookalike,
	/* [＇] */ 0xff07: LowSymbolOrLookalike,
	/* [（] */ 0xff08: LowSymbolOrLookalike,
	/* [）] */ 0xff09: LowSymbolOrLookalike,
	/* [＊] */ 0xff0a: LowSymbolOrLookalike,
	/* [＋] */ 0xff0b: LowSymbolOrLookalike,
	/* [，] */ 0xff0c: LowSymbolOrLookalike,
	/* [－] */ 0xff0d: NoProperties,
	/* [．] */ 0xff0e: NoProperties,
	/* [／] */ 0xff0f: LowSymbolOrLookalike,
	/* [０] */ 0xff10: NumeralOrLookalike,
	/* [１] */ 0xff11: NumeralOrLookalike,
	/* [２] */ 0xff12: NumeralOrLookalike,
	/* [３] */ 0xff13: NumeralOrLookalike,
	/* [４] */ 0xff14: NumeralOrLookalike,
	/* [５] */ 0xff15: NumeralOrLookalike,
	/* [６] */ 0xff16: NumeralOrLookalike,
	/* [７] */ 0xff17: NumeralOrLookalike,
	/* [８] */ 0xff18: NumeralOrLookalike,
	/* [９] */ 0xff19: NumeralOrLookalike,
	/* [：] */ 0xff1a: NoProperties,
	/* [；] */ 0xff1b: LowSymbolOrLookalike,
	/* [＜] */ 0xff1c: LowSymbolOrLookalike,
	/* [＝] */ 0xff1d: LowSymbolOrLookalike,
	/* [＞] */ 0xff1e: LowSymbolOrLookalike,
	/* [？] */ 0xff1f: LowSymbolOrLookalike,
	/* [＠] */ 0xff20: LowSymbolOrLookalike,
	/* [［] */ 0xff3b: LowSymbolOrLookalike,
	/* [＼] */ 0xff3c: LowSymbolOrLookalike,
	/* [］] */ 0xff3d: LowSymbolOrLookalike,
	/* [＾] */ 0xff3e: LowSymbolOrLookalike,
	/* [＿] */ 0xff3f: NoProperties,
	/* [｀] */ 0xff40: LowSymbolOrLookalike,
	/* [｛] */ 0xff5b: LowSymbolOrLookalike,
	/* [｜] */ 0xff5c: LowSymbolOrLookalike,
	/* [｝] */ 0xff5d: LowSymbolOrLookalike,
	/* [～] */ 0xff5e: LowSymbolOrLookalike,
	/* [￤] */ 0xffe4: LowSymbolOrLookalike,
	/* [￨] */ 0xffe8: LowSymbolOrLookalike,
	/*     */ 0xfff9: Control,
	/*     */ 0xfffa: Control,
	/*     */ 0xfffb: Control,
	/* [𐆐] */ 0x10190: LowSymbolOrLookalike,
	/*     */ 0x110bd: Control,
	/*     */ 0x110cd: Control,
	/*     */ 0x13430: Control,
	/*     */ 0x13431: Control,
	/*     */ 0x13432: Control,
	/*     */ 0x13433: Control,
	/*     */ 0x13434: Control,
	/*     */ 0x13435: Control,
	/*     */ 0x13436: Control,
	/*     */ 0x13437: Control,
	/*     */ 0x13438: Control,
	/*     */ 0x16fe4: Control,
	/*     */ 0x1bca0: Control,
	/*     */ 0x1bca1: Control,
	/*     */ 0x1bca2: Control,
	/*     */ 0x1bca3: Control,
	/*     */ 0x1d173: Control,
	/*     */ 0x1d174: Control,
	/*     */ 0x1d175: Control,
	/*     */ 0x1d176: Control,
	/*     */ 0x1d177: Control,
	/*     */ 0x1d178: Control,
	/*     */ 0x1d179: Control,
	/*     */ 0x1d17a: Control,
	/* [𝇁] */ 0x1d1c1: LowSymbolOrLookalike,
	/* [𝇂] */ 0x1d1c2: LowSymbolOrLookalike,
	/* [𝟎] */ 0x1d7ce: NumeralOrLookalike,
	/* [𝟏] */ 0x1d7cf: NumeralOrLookalike,
	/* [𝟐] */ 0x1d7d0: NumeralOrLookalike,
	/* [𝟑] */ 0x1d7d1: NumeralOrLookalike,
	/* [𝟒] */ 0x1d7d2: NumeralOrLookalike,
	/* [𝟓] */ 0x1d7d3: NumeralOrLookalike,
	/* [𝟔] */ 0x1d7d4: NumeralOrLookalike,
	/* [𝟕] */ 0x1d7d5: NumeralOrLookalike,
	/* [𝟖] */ 0x1d7d6: NumeralOrLookalike,
	/* [𝟗] */ 0x1d7d7: NumeralOrLookalike,
	/* [𝟘] */ 0x1d7d8: NumeralOrLookalike,
	/* [𝟙] */ 0x1d7d9: NumeralOrLookalike,
	/* [𝟚] */ 0x1d7da: NumeralOrLookalike,
	/* [𝟛] */ 0x1d7db: NumeralOrLookalike,
	/* [𝟜] */ 0x1d7dc: NumeralOrLookalike,
	/* [𝟝] */ 0x1d7dd: NumeralOrLookalike,
	/* [𝟞] */ 0x1d7de: NumeralOrLookalike,
	/* [𝟟] */ 0x1d7df: NumeralOrLookalike,
	/* [𝟠] */ 0x1d7e0: NumeralOrLookalike,
	/* [𝟡] */ 0x1d7e1: NumeralOrLookalike,
	/* [𝟢] */ 0x1d7e2: NumeralOrLookalike,
	/* [𝟣] */ 0x1d7e3: NumeralOrLookalike,
	/* [𝟤] */ 0x1d7e4: NumeralOrLookalike,
	/* [𝟥] */ 0x1d7e5: NumeralOrLookalike,
	/* [𝟦] */ 0x1d7e6: NumeralOrLookalike,
	/* [𝟧] */ 0x1d7e7: NumeralOrLookalike,
	/* [𝟨] */ 0x1d7e8: NumeralOrLookalike,
	/* [𝟩] */ 0x1d7e9: NumeralOrLookalike,
	/* [𝟪] */ 0x1d7ea: NumeralOrLookalike,
	/* [𝟫] */ 0x1d7eb: NumeralOrLookalike,
	/* [𝟬] */ 0x1d7ec: NumeralOrLookalike,
	/* [𝟭] */ 0x1d7ed: NumeralOrLookalike,
	/* [𝟮] */ 0x1d7ee: NumeralOrLookalike,
	/* [𝟯] */ 0x1d7ef: NumeralOrLookalike,
	/* [𝟰] */ 0x1d7f0: NumeralOrLookalike,
	/* [𝟱] */ 0x1d7f1: NumeralOrLookalike,
	/* [𝟲] */ 0x1d7f2: NumeralOrLookalike,
	/* [𝟳] */ 0x1d7f3: NumeralOrLookalike,
	/* [𝟴] */ 0x1d7f4: NumeralOrLookalike,
	/* [𝟵] */ 0x1d7f5: NumeralOrLookalike,
	/* [𝟶] */ 0x1d7f6: NumeralOrLookalike,
	/* [𝟷] */ 0x1d7f7: NumeralOrLookalike,
	/* [𝟸] */ 0x1d7f8: NumeralOrLookalike,
	/* [𝟹] */ 0x1d7f9: NumeralOrLookalike,
	/* [𝟺] */ 0x1d7fa: NumeralOrLookalike,
	/* [𝟻] */ 0x1d7fb: NumeralOrLookalike,
	/* [𝟼] */ 0x1d7fc: NumeralOrLookalike,
	/* [𝟽] */ 0x1d7fd: NumeralOrLookalike,
	/* [𝟾] */ 0x1d7fe: NumeralOrLookalike,
	/* [𝟿] */ 0x1d7ff: NumeralOrLookalike,
	/*     */ 0xe0001: Control,
	/*     */ 0xe0020: Control,
	/*     */ 0xe0021: Control,
	/*     */ 0xe0022: Control,
	/*     */ 0xe0023: Control,
	/*     */ 0xe0024: Control,
	/*     */ 0xe0025: Control,
	/*     */ 0xe0026: Control,
	/*     */ 0xe0027: Control,
	/*     */ 0xe0028: Control,
	/*     */ 0xe0029: Control,
	/*     */ 0xe002a: Control,
	/*     */ 0xe002b: Control,
	/*     */ 0xe002c: Control,
	/*     */ 0xe002d: Control,
	/*     */ 0xe002e: Control,
	/*     */ 0xe002f: Control,
	/*     */ 0xe0030: Control,
	/*     */ 0xe0031: Control,
	/*     */ 0xe0032: Control,
	/*     */ 0xe0033: Control,
	/*     */ 0xe0034: Control,
	/*     */ 0xe0035: Control,
	/*     */ 0xe0036: Control,
	/*     */ 0xe0037: Control,
	/*     */ 0xe0038: Control,
	/*     */ 0xe0039: Control,
	/*     */ 0xe003a: Control,
	/*     */ 0xe003b: Control,
	/*     */ 0xe003c: Control,
	/*     */ 0xe003d: Control,
	/*     */ 0xe003e: Control,
	/*     */ 0xe003f: Control,
	/*     */ 0xe0040: Control,
	/*     */ 0xe0041: Control,
	/*     */ 0xe0042: Control,
	/*     */ 0xe0043: Control,
	/*     */ 0xe0044: Control,
	/*     */ 0xe0045: Control,
	/*     */ 0xe0046: Control,
	/*     */ 0xe0047: Control,
	/*     */ 0xe0048: Control,
	/*     */ 0xe0049: Control,
	/*     */ 0xe004a: Control,
	/*     */ 0xe004b: Control,
	/*     */ 0xe004c: Control,
	/*     */ 0xe004d: Control,
	/*     */ 0xe004e: Control,
	/*     */ 0xe004f: Control,
	/*     */ 0xe0050: Control,
	/*     */ 0xe0051: Control,
	/*     */ 0xe0052: Control,
	/*     */ 0xe0053: Control,
	/*     */ 0xe0054: Control,
	/*     */ 0xe0055: Control,
	/*     */ 0xe0056: Control,
	/*     */ 0xe0057: Control,
	/*     */ 0xe0058: Control,
	/*     */ 0xe0059: Control,
	/*     */ 0xe005a: Control,
	/*     */ 0xe005b: Control,
	/*     */ 0xe005c: Control,
	/*     */ 0xe005d: Control,
	/*     */ 0xe005e: Control,
	/*     */ 0xe005f: Control,
	/*     */ 0xe0060: Control,
	/*     */ 0xe0061: Control,
	/*     */ 0xe0062: Control,
	/*     */ 0xe0063: Control,
	/*     */ 0xe0064: Control,
	/*     */ 0xe0065: Control,
	/*     */ 0xe0066: Control,
	/*     */ 0xe0067: Control,
	/*     */ 0xe0068: Control,
	/*     */ 0xe0069: Control,
	/*     */ 0xe006a: Control,
	/*     */ 0xe006b: Control,
	/*     */ 0xe006c: Control,
	/*     */ 0xe006d: Control,
	/*     */ 0xe006e: Control,
	/*     */ 0xe006f: Control,
	/*     */ 0xe0070: Control,
	/*     */ 0xe0071: Control,
	/*     */ 0xe0072: Control,
	/*     */ 0xe0073: Control,
	/*     */ 0xe0074: Control,
	/*     */ 0xe0075: Control,
	/*     */ 0xe0076: Control,
	/*     */ 0xe0077: Control,
	/*     */ 0xe0078: Control,
	/*     */ 0xe0079: Control,
	/*     */ 0xe007a: Control,
	/*     */ 0xe007b: Control,
	/*     */ 0xe007c: Control,
	/*     */ 0xe007d: Control,
	/*     */ 0xe007e: Control,
	/*     */ 0xe007f: Control,
}

var asciiProperties = [256]CharProperty{
	/*     */ 0x00: Control,
	/*     */ 0x01: Control,
	/*     */ 0x02: Control,
	/*     */ 0x03: Control,
	/*     */ 0x04: Control,
	/*     */ 0x05: Control,
	/*     */ 0x06: Control,
	/*     */ 0x07: Control,
	/*     */ 0x08: Control,
	/* \t  */ 0x09: Whitespace | TabReturnNewline,
	/* \n  */ 0x0a: Whitespace | TabReturnNewline,
	/*     */ 0x0b: Control,
	/*     */ 0x0c: Control,
	/* \r  */ 0x0d: Whitespace | TabReturnNewline,
	/*     */ 0x0e: Control,
	/*     */ 0x0f: Control,
	/*     */ 0x10: Control,
	/*     */ 0x11: Control,
	/*     */ 0x12: Control,
	/*     */ 0x13: Control,
	/*     */ 0x14: Control,
	/*     */ 0x15: Control,
	/*     */ 0x16: Control,
	/*     */ 0x17: Control,
	/*     */ 0x18: Control,
	/*     */ 0x19: Control,
	/*     */ 0x1a: Control,
	/*     */ 0x1b: Control,
	/*     */ 0x1c: Control,
	/*     */ 0x1d: Control,
	/*     */ 0x1e: Control,
	/*     */ 0x1f: Control,
	/* [ ] */ 0x20: Whitespace,
	/* [!] */ 0x21: LowSymbolOrLookalike,
	/* ["] */ 0x22: LowSymbolOrLookalike | QuotedTextDelimiter,
	/* [#] */ 0x23: LowSymbolOrLookalike,
	/* [$] */ 0x24: LowSymbolOrLookalike,
	/* [%] */ 0x25: LowSymbolOrLookalike,
	/* [&] */ 0x26: LowSymbolOrLookalike,
	/* ['] */ 0x27: LowSymbolOrLookalike,
	/* [(] */ 0x28: LowSymbolOrLookalike,
	/* [)] */ 0x29: LowSymbolOrLookalike,
	/* [*] */ 0x2a: LowSymbolOrLookalike,
	/* [+] */ 0x2b: LowSymbolOrLookalike,
	/* [,] */ 0x2c: LowSymbolOrLookalike,
	/* [-] */ 0x2d: NoProperties,
	/* [.] */ 0x2e: NoProperties,
	/* [0] */ 0x30: NumeralOrLookalike | MarkerIDSafe,
	/* [1] */ 0x31: NumeralOrLookalike | MarkerIDSafe,
	/* [2] */ 0x32: NumeralOrLookalike | MarkerIDSafe,
	/* [3] */ 0x33: NumeralOrLookalike | MarkerIDSafe,
	/* [4] */ 0x34: NumeralOrLookalike | MarkerIDSafe,
	/* [5] */ 0x35: NumeralOrLookalike | MarkerIDSafe,
	/* [6] */ 0x36: NumeralOrLookalike | MarkerIDSafe,
	/* [7] */ 0x37: NumeralOrLookalike | MarkerIDSafe,
	/* [8] */ 0x38: NumeralOrLookalike | MarkerIDSafe,
	/* [9] */ 0x39: NumeralOrLookalike | MarkerIDSafe,
	/* [:] */ 0x3a: NoProperties,
	/* [;] */ 0x3b: LowSymbolOrLookalike,
	/* [<] */ 0x3c: LowSymbolOrLookalike | MarkupDelimiter,
	/* [=] */ 0x3d: LowSymbolOrLookalike,
	/* [>] */ 0x3e: LowSymbolOrLookalike | MarkupDelimiter,
	/* [?] */ 0x3f: LowSymbolOrLookalike,
	/* [@] */ 0x40: LowSymbolOrLookalike,
	/* [A] */ 0x41: MarkerIDSafe,
	/* [B] */ 0x42: MarkerIDSafe,
	/* [C] */ 0x43: MarkerIDSafe,
	/* [D] */ 0x44: MarkerIDSafe,
	/* [E] */ 0x45: MarkerIDSafe,
	/* [F] */ 0x46: MarkerIDSafe,
	/* [G] */ 0x47: MarkerIDSafe,
	/* [H] */ 0x48: MarkerIDSafe,
	/* [I] */ 0x49: MarkerIDSafe,
	/* [J] */ 0x4a: MarkerIDSafe,
	/* [K] */ 0x4b: MarkerIDSafe,
	/* [L] */ 0x4c: MarkerIDSafe,
	/* [M] */ 0x4d: MarkerIDSafe,
	/* [N] */ 0x4e: MarkerIDSafe,
	/* [O] */ 0x4f: MarkerIDSafe,
	/* [P] */ 0x50: MarkerIDSafe,
	/* [Q] */ 0x51: MarkerIDSafe,
	/* [R] */ 0x52: MarkerIDSafe,
	/* [S] */ 0x53: MarkerIDSafe,
	/* [T] */ 0x54: MarkerIDSafe,
	/* [U] */ 0x55: MarkerIDSafe,
	/* [V] */ 0x56: MarkerIDSafe,
	/* [W] */ 0x57: MarkerIDSafe,
	/* [X] */ 0x58: MarkerIDSafe,
	/* [Y] */ 0x59: MarkerIDSafe,
	/* [Z] */ 0x5a: MarkerIDSafe,
	/* [[] */ 0x5b: LowSymbolOrLookalike,
	/* [\] */ 0x5c: LowSymbolOrLookalike | QuotedTextDelimiter | ArrayDelimiter | MarkupDelimiter,
	/* []] */ 0x5d: LowSymbolOrLookalike,
	/* [^] */ 0x5e: LowSymbolOrLookalike,
	/* [_] */ 0x5f: MarkerIDSafe,
	/* [`] */ 0x60: LowSymbolOrLookalike | MarkupDelimiter,
	/* [a] */ 0x61: MarkerIDSafe,
	/* [b] */ 0x62: MarkerIDSafe,
	/* [c] */ 0x63: MarkerIDSafe,
	/* [d] */ 0x64: MarkerIDSafe,
	/* [e] */ 0x65: MarkerIDSafe,
	/* [f] */ 0x66: MarkerIDSafe,
	/* [g] */ 0x67: MarkerIDSafe,
	/* [h] */ 0x68: MarkerIDSafe,
	/* [i] */ 0x69: MarkerIDSafe,
	/* [j] */ 0x6a: MarkerIDSafe,
	/* [k] */ 0x6b: MarkerIDSafe,
	/* [l] */ 0x6c: MarkerIDSafe,
	/* [m] */ 0x6d: MarkerIDSafe,
	/* [n] */ 0x6e: MarkerIDSafe,
	/* [o] */ 0x6f: MarkerIDSafe,
	/* [p] */ 0x70: MarkerIDSafe,
	/* [q] */ 0x71: MarkerIDSafe,
	/* [r] */ 0x72: MarkerIDSafe,
	/* [s] */ 0x73: MarkerIDSafe,
	/* [t] */ 0x74: MarkerIDSafe,
	/* [u] */ 0x75: MarkerIDSafe,
	/* [v] */ 0x76: MarkerIDSafe,
	/* [w] */ 0x77: MarkerIDSafe,
	/* [x] */ 0x78: MarkerIDSafe,
	/* [y] */ 0x79: MarkerIDSafe,
	/* [z] */ 0x7a: MarkerIDSafe,
	/* [{] */ 0x7b: LowSymbolOrLookalike,
	/* [|] */ 0x7c: LowSymbolOrLookalike | ArrayDelimiter,
	/* [}] */ 0x7d: LowSymbolOrLookalike,
	/* [~] */ 0x7e: LowSymbolOrLookalike,
	/*     */ 0x7f: Control,
}
