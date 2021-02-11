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

	"github.com/kstenerud/go-concise-encoding/events"

	"github.com/kstenerud/go-concise-encoding/options"
)

type indenter struct {
	indent []byte
}

func (_this *indenter) Reset() {
	_this.indent = _this.indent[:0]
	_this.indent = append(_this.indent, '\n')
}

func (_this *indenter) increase() {
	_this.indent = append(_this.indent, "    "...)
}

func (_this *indenter) decrease() {
	_this.indent = _this.indent[:len(_this.indent)-4]
}

func (_this *indenter) Get() []byte {
	return _this.indent
}

type EncoderContext struct {
	opts                options.CTEEncoderOptions
	indenter            indenter
	encoderStack        []Encoder
	CurrentEncoder      Encoder
	ContainerHasObjects bool
	currentPrefix       string

	Stream EncodeBuffer

	// ========================================
	// Arrays
	// arrayType              events.ArrayType
	// moreChunksFollow       bool
	// builtArrayBuffer       []byte
	// arrayMaxByteCount      uint64
	// arrayTotalByteCount    uint64
	// chunkExpectedByteCount uint64
	// chunkActualByteCount   uint64
	// utf8RemainderBacking   [4]byte
	// utf8RemainderBuffer    []byte
	// ValidateArrayDataFunc  func(data []byte)

	// // Marker/Reference
	// currentMarkerID   interface{}
	// markerObjectRule  EventRule
	// markedObjects     map[interface{}]DataType
	// forwardReferences map[interface{}]DataType
	// referenceCount    uint64
}

func (_this *EncoderContext) Init(opts *options.CTEEncoderOptions) {
	_this.opts = *opts
	_this.Reset()
}

func (_this *EncoderContext) Reset() {
	_this.indenter.Reset()
	_this.Stream.Reset()
	_this.encoderStack = _this.encoderStack[:0]
	_this.CurrentEncoder = nil
	_this.Stack(&globalTopLevelEncoder)
	_this.SetStandardIndentPrefix()
}

func (_this *EncoderContext) Stack(encoder Encoder) {
	_this.encoderStack = append(_this.encoderStack, encoder)
	_this.CurrentEncoder = encoder
}

func (_this *EncoderContext) Unstack() {
	_this.encoderStack = _this.encoderStack[:len(_this.encoderStack)-1]
	_this.CurrentEncoder = _this.encoderStack[len(_this.encoderStack)-1]
}

func (_this *EncoderContext) ChangeEncoder(encoder Encoder) {
	_this.encoderStack[len(_this.encoderStack)-1] = encoder
	_this.CurrentEncoder = encoder
}

func (_this *EncoderContext) IncreaseIndent() {
	_this.indenter.increase()
}

func (_this *EncoderContext) DecreaseIndent() {
	_this.indenter.decrease()
}

func (_this *EncoderContext) WriteBasicIndent() {
	_this.Stream.AddBytes(_this.indenter.Get())
}

func (_this *EncoderContext) SetIndentPrefix(value string) {
	_this.currentPrefix = value
}

func (_this *EncoderContext) SetStandardIndentPrefix() {
	_this.SetIndentPrefix(string(_this.indenter.Get()))
}

func (_this *EncoderContext) SetStandardMapKeyPrefix() {
	_this.SetIndentPrefix(string(_this.indenter.Get()))
}

func (_this *EncoderContext) SetStandardMapValuePrefix() {
	_this.SetIndentPrefix(" = ")
}

func (_this *EncoderContext) ClearPrefix() {
	_this.SetIndentPrefix("")
}

func (_this *EncoderContext) WriteCurrentPrefix() {
	_this.Stream.AddString(_this.currentPrefix)
	// TODO: Need to do this?
	// _this.ClearMapPrefix()
}

func (_this *EncoderContext) PerformStandardKVSeparation() {
	_this.Stream.AddString(" = ")
	_this.ClearPrefix()
}

func (_this *EncoderContext) EndContainerXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX() {
	_this.CurrentEncoder.End(_this)
	_this.Unstack()
	_this.CurrentEncoder.ChildContainerFinished(_this)
}

func (_this *EncoderContext) BeginStandardList() {
	_this.Stack(&globalListEncoder)
	_this.CurrentEncoder.Begin(_this)
}

func (_this *EncoderContext) BeginStandardMap() {
	_this.Stack(&globalMapKeyEncoder)
	_this.CurrentEncoder.Begin(_this)
}

func (_this *EncoderContext) BeginStandardMarkup() {
	panic(fmt.Errorf("TODO: EncoderContext.BeginMarkup"))
}

func (_this *EncoderContext) BeginStandardMetadata() {
	_this.Stack(&globalMetadataKeyEncoder)
	_this.CurrentEncoder.Begin(_this)
}

func (_this *EncoderContext) BeginStandardComment() {
	panic(fmt.Errorf("TODO: EncoderContext.BeginComment"))
}

func (_this *EncoderContext) BeginStandardMarker() {
	panic(fmt.Errorf("TODO: EncoderContext.BeginMarker"))
}

func (_this *EncoderContext) BeginStandardReference() {
	panic(fmt.Errorf("TODO: EncoderContext.BeginReference"))
}

func (_this *EncoderContext) BeginStandardConcatenate() {
	panic(fmt.Errorf("TODO: EncoderContext.BeginConcatenate"))
}

func (_this *EncoderContext) BeginStandardConstant(name []byte, explicitValue bool) {
	panic(fmt.Errorf("TODO: EncoderContext.BeginConstant"))
}

func (_this *EncoderContext) BeginStandardNA() {
	_this.Stack(&globalNAEncoder)
	_this.CurrentEncoder.Begin(_this)
}

func (_this *EncoderContext) BeginStandardArray(arayType events.ArrayType) {
	panic(fmt.Errorf("TODO: EncoderContext.BeginArray"))
}

// pre-write (indent)
// post-write (lf?)
// list-type
// map-type (key section, value section)
// array-type
// metadata follow
// comment follow

// string in comment
// string in markup contents
// string as ID
// int as ID
// constant name
// custom string
// NA stuff

// Types that can be printed differently:
// - string (quoted, unquoted) (escape/noescape) (trim/notrim)
// - int (pos or neg) (bin, oct, dec, hex) (with/without prefix)
// - binary float (dec, hex) (with/without prefix)

// =============================================================================

type ArrayRenderer interface {
	RenderArrayPortion(ctx *EncoderContext, data []byte)
	RenderArrayComplete(ctx *EncoderContext, data []byte)
}

var (
	arrayRendererNone ArrayRendererNone
)

type ArrayRendererNone struct{}

func (_this *ArrayRendererNone) RenderArrayPortion(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: ArrayRendererNone cannot respond to RenderArrayPortion"))
}

func (_this *ArrayRendererNone) RenderArrayComplete(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: ArrayRendererNone cannot respond to RenderArrayComplete"))
}
