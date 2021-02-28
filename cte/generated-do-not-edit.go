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
	"math/big"
	"reflect"
	"time"

	"github.com/kstenerud/go-concise-encoding/events"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
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
func (_this *constantEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *constantEncoder) ChildContainerFinished(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to ChildContainerFinished", reflect.TypeOf(_this)))
}
func (_this *constantEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *constantEncoder) BeginComment(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginComment", reflect.TypeOf(_this)))
}
func (_this *constantEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *constantEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
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
func (_this *mapKeyEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
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
func (_this *mapKeyEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
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
func (_this *metadataKeyEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
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
func (_this *metadataKeyEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
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
func (_this *markupNameEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginComment(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginComment", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *markupNameEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *markupKeyEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *markupValueEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *markupValueEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *markupValueEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *markupValueEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBool", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeTrue(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTrue", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeFalse(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFalse", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodePositiveInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodePositiveInt", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNegativeInt", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeInt(ctx *EncoderContext, value int64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeInt", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeBigInt(ctx *EncoderContext, value *big.Int) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigInt", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFloat", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigFloat", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTime", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeCompactTime", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeUUID", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginMarker(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarker", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginReference(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginReference", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginConcatenate(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConcatenate", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginConstant(ctx *EncoderContext, name []byte, explicitValue bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConstant", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *markupContentsEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBool", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeTrue(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTrue", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeFalse(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFalse", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodePositiveInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodePositiveInt", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNegativeInt", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeInt(ctx *EncoderContext, value int64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeInt", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeBigInt(ctx *EncoderContext, value *big.Int) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigInt", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFloat", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigFloat", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTime", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeCompactTime", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeUUID", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginMarker(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarker", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginReference(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginReference", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginConcatenate(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConcatenate", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginConstant(ctx *EncoderContext, name []byte, explicitValue bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConstant", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *commentEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBool", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeTrue(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTrue", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeFalse(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFalse", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNegativeInt", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFloat", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigFloat", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTime", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeCompactTime", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeUUID", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginComment(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginComment", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginMarker(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarker", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginReference(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginReference", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginConcatenate(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConcatenate", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *referenceEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBool", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeTrue(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTrue", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeFalse(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFalse", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNegativeInt", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFloat", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigFloat", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTime", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeCompactTime", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeUUID", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginComment(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginComment", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginMarker(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarker", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginReference(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginReference", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginConcatenate(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConcatenate", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) BeginArrayChunk(ctx *EncoderContext, length uint64, moreChunksFollow bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArrayChunk", reflect.TypeOf(_this)))
}
func (_this *markerIDEncoder) EncodeArrayData(ctx *EncoderContext, data []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArrayData", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) Begin(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to Begin", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) End(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to End", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) ChildContainerFinished(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to ChildContainerFinished", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBool", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeTrue(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTrue", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeFalse(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFalse", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodePositiveInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodePositiveInt", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNegativeInt", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeInt(ctx *EncoderContext, value int64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeInt", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeBigInt(ctx *EncoderContext, value *big.Int) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigInt", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeFloat", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigFloat", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeBigDecimalFloat", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeNan", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeTime", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeCompactTime", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeUUID", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginList(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginList", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginMap(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMap", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginMarkup(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarkup", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginMetadata(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMetadata", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginComment(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginComment", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginMarker(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginMarker", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginReference(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginReference", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginConcatenate(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConcatenate", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginConstant(ctx *EncoderContext, name []byte, explicitValue bool) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginConstant", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginNA(ctx *EncoderContext) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginNA", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeArray(ctx *EncoderContext, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeArray", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) EncodeStringlikeArray(ctx *EncoderContext, arrayType events.ArrayType, data string) {
	panic(fmt.Errorf("BUG: %v cannot respond to EncodeStringlikeArray", reflect.TypeOf(_this)))
}
func (_this *arrayEncoder) BeginArray(ctx *EncoderContext, arrayType events.ArrayType) {
	panic(fmt.Errorf("BUG: %v cannot respond to BeginArray", reflect.TypeOf(_this)))
}
