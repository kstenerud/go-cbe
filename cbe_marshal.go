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

package concise_encoding

import (
	"fmt"
)

type CBEMarshalerOptions struct {
	Encoder  CBEEncoderOptions
	Iterator IteratorOptions
}

func MarshalCBE(object interface{}, options *CBEMarshalerOptions) (document []byte, err error) {
	if options == nil {
		options = &CBEMarshalerOptions{}
	}
	defer func() {
		if !DebugOptions.PassThroughPanics {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
			}
		}
	}()

	encoder := NewCBEEncoder(&options.Encoder)
	IterateObject(object, encoder, &options.Iterator)
	document = encoder.Document()
	return
}

type CBEUnmarshalerOptions struct {
	Decoder CBEDecoderOptions
	Builder BuilderOptions
	Rules   RuleOptions
}

func UnmarshalCBE(document []byte, template interface{}, options *CBEUnmarshalerOptions) (decoded interface{}, err error) {
	if options == nil {
		options = &CBEUnmarshalerOptions{}
	}
	defer func() {
		if !DebugOptions.PassThroughPanics {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
			}
		}
	}()

	builder := NewBuilderFor(template, &options.Builder)
	rules := NewRules(&options.Rules, builder)
	decoder := NewCBEDecoder(document, rules, &options.Decoder)
	decoder.Decode()
	decoded = builder.GetBuiltObject()
	return
}
