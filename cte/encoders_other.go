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
	"math/big"
	"time"

	"github.com/kstenerud/go-concise-encoding/events"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
)

type naEncoder struct{}

var globalNAEncoder naEncoder

func (_this *naEncoder) String() string { return "naEncoder" }

func (_this *naEncoder) Begin(ctx *EncoderContext) {
	ctx.Stream.AddString("@na")
}

func (_this *naEncoder) prepareToWrite(ctx *EncoderContext) {
	ctx.Stream.WriteSeparator()
	ctx.Unstack()
}

func (_this *naEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteBool(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeTrue(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteTrue()
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeFalse(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteFalse()
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodePositiveInt(ctx *EncoderContext, value uint64) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WritePositiveInt(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteNegativeInt(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeInt(ctx *EncoderContext, value int64) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteInt(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeBigInt(ctx *EncoderContext, value *big.Int) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteBigInt(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteFloat(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteBigFloat(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteDecimalFloat(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteBigDecimalFloat(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteNan(signaling)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteTime(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteCompactTime(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteUUID(value)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) BeginList(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardList()
}
func (_this *naEncoder) BeginMap(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardMap()
}
func (_this *naEncoder) BeginMarkup(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardMarkup()
}
func (_this *naEncoder) BeginMarker(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardMarker()
}
func (_this *naEncoder) BeginReference(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardReference()
}
func (_this *naEncoder) BeginConcatenate(ctx *EncoderContext) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardConcatenate()
}
func (_this *naEncoder) BeginConstant(ctx *EncoderContext, name []byte, explicitValue bool) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardConstant(name, explicitValue)
}
func (_this *naEncoder) BeginNA(ctx *EncoderContext) {
	// Only unstack
	ctx.Unstack()
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeArray(ctx *EncoderContext, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteArray(arrayType, elementCount, data)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) EncodeStringlikeArray(ctx *EncoderContext, arrayType events.ArrayType, data string) {
	_this.prepareToWrite(ctx)
	ctx.Stream.WriteStringlikeArray(arrayType, data)
	ctx.CurrentEncoder.ChildContainerFinished(ctx)
}
func (_this *naEncoder) BeginArray(ctx *EncoderContext, arrayType events.ArrayType) {
	_this.prepareToWrite(ctx)
	ctx.BeginStandardArray(arrayType)
}

// =============================================================================

type postInvisibleEncoder struct{}

var globalPostInvisibleEncoder postInvisibleEncoder

func (_this *postInvisibleEncoder) String() string { return "postInvisibleEncoder" }

func (_this *postInvisibleEncoder) removeSelf(ctx *EncoderContext) Encoder {
	ctx.Unstack()
	ctx.ClearPrefix()
	return ctx.CurrentEncoder
}

func (_this *postInvisibleEncoder) EncodeBool(ctx *EncoderContext, value bool) {
	_this.removeSelf(ctx).EncodeBool(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeTrue(ctx *EncoderContext) {
	_this.removeSelf(ctx).EncodeTrue(ctx)
}
func (_this *postInvisibleEncoder) EncodeFalse(ctx *EncoderContext) {
	_this.removeSelf(ctx).EncodeFalse(ctx)
}
func (_this *postInvisibleEncoder) EncodePositiveInt(ctx *EncoderContext, value uint64) {
	_this.removeSelf(ctx).EncodePositiveInt(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeNegativeInt(ctx *EncoderContext, value uint64) {
	_this.removeSelf(ctx).EncodeNegativeInt(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeInt(ctx *EncoderContext, value int64) {
	_this.removeSelf(ctx).EncodeInt(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeBigInt(ctx *EncoderContext, value *big.Int) {
	_this.removeSelf(ctx).EncodeBigInt(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeFloat(ctx *EncoderContext, value float64) {
	_this.removeSelf(ctx).EncodeFloat(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeBigFloat(ctx *EncoderContext, value *big.Float) {
	_this.removeSelf(ctx).EncodeBigFloat(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeDecimalFloat(ctx *EncoderContext, value compact_float.DFloat) {
	_this.removeSelf(ctx).EncodeDecimalFloat(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeBigDecimalFloat(ctx *EncoderContext, value *apd.Decimal) {
	_this.removeSelf(ctx).EncodeBigDecimalFloat(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeNan(ctx *EncoderContext, signaling bool) {
	_this.removeSelf(ctx).EncodeNan(ctx, signaling)
}
func (_this *postInvisibleEncoder) EncodeTime(ctx *EncoderContext, value time.Time) {
	_this.removeSelf(ctx).EncodeTime(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeCompactTime(ctx *EncoderContext, value compact_time.Time) {
	_this.removeSelf(ctx).EncodeCompactTime(ctx, value)
}
func (_this *postInvisibleEncoder) EncodeUUID(ctx *EncoderContext, value []byte) {
	_this.removeSelf(ctx).EncodeUUID(ctx, value)
}
func (_this *postInvisibleEncoder) BeginList(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginList(ctx)
}
func (_this *postInvisibleEncoder) BeginMap(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginMap(ctx)
}
func (_this *postInvisibleEncoder) BeginMarkup(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginMarkup(ctx)
}
func (_this *postInvisibleEncoder) BeginMetadata(ctx *EncoderContext) {
	ctx.BeginStandardMetadata()
}
func (_this *postInvisibleEncoder) BeginComment(ctx *EncoderContext) {
	ctx.BeginStandardComment()
}
func (_this *postInvisibleEncoder) BeginMarker(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginMarker(ctx)
}
func (_this *postInvisibleEncoder) BeginReference(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginReference(ctx)
}
func (_this *postInvisibleEncoder) BeginConcatenate(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginConcatenate(ctx)
}
func (_this *postInvisibleEncoder) BeginConstant(ctx *EncoderContext, name []byte, explicitValue bool) {
	_this.removeSelf(ctx).BeginConstant(ctx, name, explicitValue)
}
func (_this *postInvisibleEncoder) BeginNA(ctx *EncoderContext) {
	_this.removeSelf(ctx).BeginNA(ctx)
}
func (_this *postInvisibleEncoder) EncodeArray(ctx *EncoderContext, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	_this.removeSelf(ctx).EncodeArray(ctx, arrayType, elementCount, data)
}
func (_this *postInvisibleEncoder) EncodeStringlikeArray(ctx *EncoderContext, arrayType events.ArrayType, data string) {
	_this.removeSelf(ctx).EncodeStringlikeArray(ctx, arrayType, data)
}
func (_this *postInvisibleEncoder) BeginArray(ctx *EncoderContext, arrayType events.ArrayType) {
	_this.removeSelf(ctx).BeginArray(ctx, arrayType)
}