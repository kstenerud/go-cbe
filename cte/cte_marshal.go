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
	"fmt"

	"github.com/kstenerud/go-concise-encoding/builder"
	"github.com/kstenerud/go-concise-encoding/debug"
	"github.com/kstenerud/go-concise-encoding/iterator"
	"github.com/kstenerud/go-concise-encoding/rules"
)

type MarshalerOptions struct {
	Encoder  EncoderOptions
	Iterator iterator.IteratorOptions
}

// Marshal a go object into a CTE document
func Marshal(object interface{}, options *MarshalerOptions) (document []byte, err error) {
	if options == nil {
		options = &MarshalerOptions{}
	}
	defer func() {
		if !debug.DebugOptions.PassThroughPanics {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
			}
		}
	}()

	encoder := NewEncoder(&options.Encoder)
	iterator.IterateObject(object, encoder, &options.Iterator)
	document = encoder.GetBuiltDocument()
	return
}

type CTEUnmarshalerOptions struct {
	Decoder DecoderOptions
	Builder builder.BuilderOptions
	Rules   rules.RuleOptions
}

// Unmarshal a CTE document, creating an object of the same type as the template.
func Unmarshal(document []byte, template interface{}, options *CTEUnmarshalerOptions) (decoded interface{}, err error) {
	if options == nil {
		options = &CTEUnmarshalerOptions{}
	}
	defer func() {
		if !debug.DebugOptions.PassThroughPanics {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
			}
		}
	}()

	builder := builder.NewBuilderFor(template, &options.Builder)
	rules := rules.NewRules(&options.Rules, builder)
	decoder := NewDecoder(document, rules, &options.Decoder)
	decoder.Decode()
	decoded = builder.GetBuiltObject()
	return
}