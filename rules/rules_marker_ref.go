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

// Imposes the structural rules that enforce a well-formed concise encoding
// document.
package rules

import (
	"fmt"
	"math/big"

	"github.com/kstenerud/go-concise-encoding/events"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
)

type MarkerIDKeyableRule struct{}

func (_this *MarkerIDKeyableRule) String() string         { return "Marker ID Keyable Rule" }
func (_this *MarkerIDKeyableRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *MarkerIDKeyableRule) OnInt(ctx *Context, value int64) {
	if value < 0 {
		panic(fmt.Errorf("Marker ID (%v) cannot be negative", value))
	}
	ctx.BeginMarkedObjectKeyable(uint64(value))
}
func (_this *MarkerIDKeyableRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.BeginMarkedObjectKeyable(value)
}
func (_this *MarkerIDKeyableRule) OnBigInt(ctx *Context, value *big.Int) {
	if !value.IsUint64() {
		panic(fmt.Errorf("Marker ID (%v) is out of range", value))
	}
	ctx.BeginMarkedObjectKeyable(value.Uint64())
}
func (_this *MarkerIDKeyableRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.ValidateFullArrayMarkerID(arrayType, elementCount, data)
	ctx.BeginMarkedObjectKeyable(string(data))
}
func (_this *MarkerIDKeyableRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.ValidateFullArrayMarkerIDString(arrayType, data)
	ctx.BeginMarkedObjectKeyable(data)
}
func (_this *MarkerIDKeyableRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginStringBuilder(arrayType, ctx.ValidateContentsMarkerID)
}
func (_this *MarkerIDKeyableRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	id := ctx.GetBuiltArrayAsString()
	ctx.BeginMarkedObjectKeyable(id)
}

// =============================================================================

type MarkerIDAnyTypeRule struct{}

func (_this *MarkerIDAnyTypeRule) String() string         { return "Marker ID Any Type Rule" }
func (_this *MarkerIDAnyTypeRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *MarkerIDAnyTypeRule) OnInt(ctx *Context, value int64) {
	if value < 0 {
		panic(fmt.Errorf("Marker ID (%v) cannot be negative", value))
	}
	ctx.BeginMarkedObjectAnyType(uint64(value))
}
func (_this *MarkerIDAnyTypeRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.BeginMarkedObjectAnyType(value)
}
func (_this *MarkerIDAnyTypeRule) OnBigInt(ctx *Context, value *big.Int) {
	if !value.IsUint64() {
		panic(fmt.Errorf("Marker ID (%v) is out of range", value))
	}
	ctx.BeginMarkedObjectAnyType(value.Uint64())
}
func (_this *MarkerIDAnyTypeRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.ValidateFullArrayMarkerID(arrayType, elementCount, data)
	ctx.BeginMarkedObjectAnyType(string(data))
}
func (_this *MarkerIDAnyTypeRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.ValidateFullArrayMarkerIDString(arrayType, data)
	ctx.BeginMarkedObjectAnyType(string(data))
}
func (_this *MarkerIDAnyTypeRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginStringBuilder(arrayType, ctx.ValidateContentsMarkerID)
}
func (_this *MarkerIDAnyTypeRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	id := ctx.GetBuiltArrayAsString()
	ctx.BeginMarkedObjectKeyable(id)
}

// =============================================================================

type MarkedObjectKeyableRule struct{}

func (_this *MarkedObjectKeyableRule) String() string         { return "Marked Keyable Object Rule" }
func (_this *MarkedObjectKeyableRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *MarkedObjectKeyableRule) OnKeyableObject(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnKeyableObject(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnInt(ctx *Context, value int64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnPositiveInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnBigInt(ctx *Context, value *big.Int) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnFloat(ctx *Context, value float64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnBigFloat(ctx *Context, value *big.Float) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnDecimalFloat(ctx *Context, value compact_float.DFloat) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnBigDecimalFloat(ctx *Context, value *apd.Decimal) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnReference(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnReference(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnConstant(ctx *Context, name []byte, explicitValue bool) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnConstant(ctx, name, explicitValue)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnArray(ctx, arrayType, elementCount, data)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnStringlikeArray(ctx, arrayType, data)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectKeyableRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginArrayKeyable(arrayType)
}
func (_this *MarkedObjectKeyableRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}

// =============================================================================

type MarkedObjectAnyTypeRule struct{}

func (_this *MarkedObjectAnyTypeRule) String() string         { return "Marked Object Rule" }
func (_this *MarkedObjectAnyTypeRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *MarkedObjectAnyTypeRule) OnNonKeyableObject(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnKeyableObject(ctx)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *MarkedObjectAnyTypeRule) OnNA(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnNA(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnKeyableObject(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnKeyableObject(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnInt(ctx *Context, value int64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnPositiveInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnBigInt(ctx *Context, value *big.Int) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnFloat(ctx *Context, value float64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnBigFloat(ctx *Context, value *big.Float) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnDecimalFloat(ctx *Context, value compact_float.DFloat) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnBigDecimalFloat(ctx *Context, value *apd.Decimal) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *MarkedObjectAnyTypeRule) OnList(ctx *Context) {
	ctx.ParentRule().OnList(ctx)
}
func (_this *MarkedObjectAnyTypeRule) OnMap(ctx *Context) {
	ctx.ParentRule().OnMap(ctx)
}
func (_this *MarkedObjectAnyTypeRule) OnMarkup(ctx *Context) {
	ctx.ParentRule().OnMarkup(ctx)
}
func (_this *MarkedObjectAnyTypeRule) OnReference(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnReference(ctx)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *MarkedObjectAnyTypeRule) OnConstant(ctx *Context, name []byte, explicitValue bool) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnConstant(ctx, name, explicitValue)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *MarkedObjectAnyTypeRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnArray(ctx, arrayType, elementCount, data)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *MarkedObjectAnyTypeRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnStringlikeArray(ctx, arrayType, data)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *MarkedObjectAnyTypeRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.ParentRule().OnArrayBegin(ctx, arrayType)
}
func (_this *MarkedObjectAnyTypeRule) OnChildContainerEnded(ctx *Context, cType DataType) {
	ctx.MarkObject(cType)
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, cType)
}

// =============================================================================

type ReferenceKeyableRule struct{}

func (_this *ReferenceKeyableRule) String() string         { return "Reference To Keyable Type Rule" }
func (_this *ReferenceKeyableRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *ReferenceKeyableRule) OnInt(ctx *Context, value int64) {
	if value < 0 {
		panic(fmt.Errorf("Reference ID (%v) cannot be negative", value))
	}
	ctx.UnstackRule()
	ctx.ReferenceObject(uint64(value), AllowKeyable)
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}
func (_this *ReferenceKeyableRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.UnstackRule()
	ctx.ReferenceObject(value, AllowKeyable)
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}
func (_this *ReferenceKeyableRule) OnBigInt(ctx *Context, value *big.Int) {
	if !value.IsUint64() {
		panic(fmt.Errorf("Reference ID (%v) is out of range", value))
	}
	ctx.UnstackRule()
	ctx.ReferenceObject(value.Uint64(), AllowKeyable)
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}
func (_this *ReferenceKeyableRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	switch arrayType {
	case events.ArrayTypeString:
		ctx.ValidateFullArrayMarkerID(arrayType, elementCount, data)
		ctx.UnstackRule()
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	default:
		panic(fmt.Errorf("Reference ID cannot be type %v", arrayType))
	}
}
func (_this *ReferenceKeyableRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	switch arrayType {
	case events.ArrayTypeString:
		ctx.ValidateFullArrayMarkerIDString(arrayType, data)
		ctx.UnstackRule()
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	default:
		panic(fmt.Errorf("Reference ID cannot be type %v", arrayType))
	}
}
func (_this *ReferenceKeyableRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	switch arrayType {
	case events.ArrayTypeString:
		ctx.BeginStringBuilder(arrayType, ctx.ValidateContentsMarkerID)
	default:
		panic(fmt.Errorf("Reference ID cannot be type %v", arrayType))
	}
}
func (_this *ReferenceKeyableRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	id := ctx.GetBuiltArrayAsString()
	ctx.ReferenceObject(id, AllowAnyType)
}

// =============================================================================

type ReferenceAnyTypeRule struct{}

func (_this *ReferenceAnyTypeRule) String() string         { return "Reference To Any Type Rule" }
func (_this *ReferenceAnyTypeRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *ReferenceAnyTypeRule) OnInt(ctx *Context, value int64) {
	if value < 0 {
		panic(fmt.Errorf("Reference ID (%v) cannot be negative", value))
	}
	ctx.UnstackRule()
	ctx.ReferenceObject(uint64(value), AllowAnyType)
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}
func (_this *ReferenceAnyTypeRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.UnstackRule()
	ctx.ReferenceObject(value, AllowAnyType)
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}
func (_this *ReferenceAnyTypeRule) OnBigInt(ctx *Context, value *big.Int) {
	if !value.IsUint64() {
		panic(fmt.Errorf("Reference ID (%v) is out of range", value))
	}
	ctx.UnstackRule()
	ctx.ReferenceObject(value.Uint64(), AllowAnyType)
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}
func (_this *ReferenceAnyTypeRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	switch arrayType {
	case events.ArrayTypeString:
		ctx.ValidateFullArrayMarkerID(arrayType, elementCount, data)
		ctx.UnstackRule()
		ctx.ReferenceObject(string(data), AllowAnyType)
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	case events.ArrayTypeResourceID:
		ctx.UnstackRule()
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	default:
		panic(fmt.Errorf("Reference ID cannot be type %v", arrayType))
	}
}
func (_this *ReferenceAnyTypeRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	switch arrayType {
	case events.ArrayTypeString:
		ctx.ValidateFullArrayMarkerIDString(arrayType, data)
		ctx.UnstackRule()
		ctx.ReferenceObject(string(data), AllowAnyType)
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	case events.ArrayTypeResourceID:
		ctx.UnstackRule()
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	default:
		panic(fmt.Errorf("Reference ID cannot be type %v", arrayType))
	}
}
func (_this *ReferenceAnyTypeRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	switch arrayType {
	case events.ArrayTypeString:
		ctx.BeginStringBuilder(arrayType, ctx.ValidateContentsMarkerID)
	case events.ArrayTypeResourceID:
		ctx.BeginArrayRIDReference(arrayType)
	default:
		panic(fmt.Errorf("Reference ID cannot be type %v", arrayType))
	}
}
func (_this *ReferenceAnyTypeRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	id := ctx.GetBuiltArrayAsString()
	ctx.ReferenceObject(id, AllowAnyType)
}

// =============================================================================

type TLReferenceRIDRule struct{}

func (_this *TLReferenceRIDRule) String() string         { return "Reference To Resource ID Rule" }
func (_this *TLReferenceRIDRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *TLReferenceRIDRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	switch arrayType {
	case events.ArrayTypeResourceID:
		ctx.ValidateLengthRID(uint64(len(data)))
		ctx.ValidateContentsRID(data)
		ctx.UnstackRule()
		ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
	default:
		panic(fmt.Errorf("Top-level Reference ID cannot be type %v", arrayType))
	}
}
func (_this *TLReferenceRIDRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginArrayRIDReference(arrayType)
}
func (_this *TLReferenceRIDRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	// TODO: Make this properly
	_this.OnArray(ctx, arrayType, uint64(len(data)), []byte(data))
}
func (_this *TLReferenceRIDRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	// Toss out the result because it's a resource ID
}
