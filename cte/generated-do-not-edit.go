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
  // IF THIS LINE SHOWS UP IN THE GIT DIFF AFTER GO FMT, THIS FILE HAS BEEN EDITED

package cte

import (
	"fmt"
	"reflect"
)
func (_this *topLevelEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *topLevelEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *topLevelEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *topLevelEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *naEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *naEncoder) ChildContainerFinished(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to ChildContainerFinished", reflect.TypeOf(_this)))
}
func (_this *naEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *naEncoder) BeginComment(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginComment", reflect.TypeOf(_this)))
}
func (_this *naEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *naEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *postInvisibleEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *postInvisibleEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *postInvisibleEncoder) ChildContainerFinished(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to ChildContainerFinished", reflect.TypeOf(_this)))
}
func (_this *postInvisibleEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *postInvisibleEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *listEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *listEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *mapKeyEncoder) ChildContainerFinished(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to ChildContainerFinished", reflect.TypeOf(_this)))
}
func (_this *mapKeyEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *mapKeyEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *mapKeyEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *mapKeyEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *mapKeyEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *mapValueEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *mapValueEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *mapValueEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *mapValueEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *metadataKeyEncoder) ChildContainerFinished(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to ChildContainerFinished", reflect.TypeOf(_this)))
}
func (_this *metadataKeyEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *metadataKeyEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *metadataKeyEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *metadataKeyEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *metadataKeyEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *metadataValueEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *metadataValueEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *metadataValueEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *metadataValueEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
