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
	"fmt"
	"unicode/utf8"

	"github.com/kstenerud/go-concise-encoding/internal/chars"
)

func getStringRequirements(str []byte) (escapeCount int, requiresQuotes bool) {
	if len(str) == 0 {
		return 0, true
	}

	firstRune, _ := utf8.DecodeRune(str)
	if chars.RuneHasProperty(firstRune, chars.CharNeedsQuoteFirst) {
		requiresQuotes = true
	}

	for _, ch := range string(str) {
		props := chars.GetRuneProperty(ch)
		if props.HasProperty(chars.CharNeedsQuote) {
			requiresQuotes = true
		}
		if props.HasProperty(chars.CharNeedsEscapeQuoted) {
			escapeCount++
		}
	}
	return
}

func needsEscapesStringlikeArray(str []byte) bool {
	for _, ch := range string(str) {
		if chars.RuneHasProperty(ch, chars.CharNeedsEscapeArray) {
			return true
		}
	}
	return false
}

func needsEscapesMarkup(str []byte) bool {
	for _, ch := range string(str) {
		if chars.RuneHasProperty(ch, chars.CharNeedsEscapeMarkup) {
			return true
		}
	}
	return false
}

func containsEscapes(str []byte) bool {
	for _, b := range str {
		if b == '\\' {
			return true
		}
	}
	return false
}

// ============================================================================

func escapeCharQuoted(ch rune) []byte {
	switch ch {
	case '\t':
		return []byte(`\t`)
	case '\r':
		return []byte(`\r`)
	case '\n':
		return []byte(`\n`)
	case '"':
		return []byte(`\"`)
	case '*':
		return []byte(`\*`)
	case '/':
		return []byte(`\/`)
	case '\\':
		return []byte(`\\`)
	}
	return unicodeEscape(ch)
}

func unicodeEscape(ch rune) []byte {
	hex := fmt.Sprintf("%x", ch)
	return []byte(fmt.Sprintf("\\%d%s", len(hex), hex))
}

func escapeCharStringArray(ch rune) []byte {
	switch ch {
	case '|':
		return []byte(`\|`)
	case '\\':
		return []byte(`\\`)
	case '\t':
		return []byte(`\t`)
	case '\r':
		return []byte(`\r`)
	case '\n':
		return []byte(`\n`)
	}
	return unicodeEscape(ch)
}

func escapeCharMarkup(ch rune) []byte {
	switch ch {
	case '*':
		// TODO: Check ahead for /* */ instead of blindly escaping
		return []byte(`\*`)
	case '/':
		// TODO: Check ahead for /* */ instead of blindly escaping
		return []byte(`\/`)
	case '<':
		return []byte(`\<`)
	case '>':
		return []byte(`\>`)
	case 0xa0:
		return []byte(`\_`)
	case 0xad:
		return []byte(`\-`)
	case '\\':
		return []byte(`\\`)
	}
	return unicodeEscape(ch)
}

// Ordered from least common to most common, chosen to not be confused by
// a human with other CTE document structural characters.
var verbatimSentinelAlphabet = []byte("~%*+;=^_23456789ZQXJVKBPYGCFMWULDHSNOIRATE10zqxjvkbpygcfmwuldhsnoirate")

func generateVerbatimSentinel(str []byte) []byte {
	// Try all 1, 2, and 3-character sequences picked from a safe alphabet.

	usedChars := [256]bool{}
	for _, ch := range str {
		usedChars[ch] = true
	}

	var sentinelBuff [3]byte

	for _, ch := range verbatimSentinelAlphabet {
		if !usedChars[ch] {
			return []byte{ch}
		}
	}

	for _, ch0 := range verbatimSentinelAlphabet {
		for _, ch1 := range verbatimSentinelAlphabet {
			sentinelBuff[0] = ch0
			sentinelBuff[1] = ch1
			sentinel := sentinelBuff[:2]
			if !bytes.Contains(str, sentinel) {
				return sentinel
			}
		}
	}

	for _, ch0 := range verbatimSentinelAlphabet {
		for _, ch1 := range verbatimSentinelAlphabet {
			for _, ch2 := range verbatimSentinelAlphabet {
				sentinelBuff[0] = ch0
				sentinelBuff[1] = ch1
				sentinelBuff[2] = ch2
				sentinel := sentinelBuff[:3]
				if !bytes.Contains(str, sentinel) {
					return sentinel
				}
			}
		}
	}

	// If we're here, all 450,000 three-character sequences have occurred in
	// the string. At this point, we conclude that it's a specially crafted
	// attack string, and not naturally occurring.
	panic(fmt.Errorf("could not generate verbatim sentinel for malicious string [%v]", str))
}
