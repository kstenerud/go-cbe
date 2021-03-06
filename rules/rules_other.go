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

type BeginDocumentRule struct{}

func (_this *BeginDocumentRule) String() string               { return "Begin Document Rule" }
func (_this *BeginDocumentRule) OnBeginDocument(ctx *Context) { ctx.SwitchVersion() }

// =============================================================================

type EndDocumentRule struct{}

func (_this *EndDocumentRule) String() string             { return "End Document Rule" }
func (_this *EndDocumentRule) OnEndDocument(ctx *Context) { ctx.EndDocument() }

// =============================================================================

type TerminalRule struct{}

func (_this *TerminalRule) String() string { return "Terminal Rule" }

// =============================================================================

type VersionRule struct{}

func (_this *VersionRule) String() string { return "Version Rule" }
func (_this *VersionRule) OnVersion(ctx *Context, version uint64) {
	if version != ctx.ExpectedVersion {
		panic(fmt.Errorf("expected version %v but got version %v", ctx.ExpectedVersion, version))
	}
	ctx.SwitchTopLevel()
}

// =============================================================================

type TopLevelRule struct{}

func (_this *TopLevelRule) String() string                                 { return "Top Level Rule" }
func (_this *TopLevelRule) OnKeyableObject(ctx *Context)                   { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnNonKeyableObject(ctx *Context)                { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnNA(ctx *Context)                              { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnChildContainerEnded(ctx *Context, _ DataType) { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnPadding(ctx *Context)                         { /* Nothing to do */ }
func (_this *TopLevelRule) OnInt(ctx *Context, value int64)                { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnPositiveInt(ctx *Context, value uint64)       { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnBigInt(ctx *Context, value *big.Int)          { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnFloat(ctx *Context, value float64)            { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnBigFloat(ctx *Context, value *big.Float)      { ctx.SwitchEndDocument() }
func (_this *TopLevelRule) OnDecimalFloat(ctx *Context, value compact_float.DFloat) {
	ctx.SwitchEndDocument()
}
func (_this *TopLevelRule) OnBigDecimalFloat(ctx *Context, value *apd.Decimal) {
	ctx.SwitchEndDocument()
}
func (_this *TopLevelRule) OnList(ctx *Context)      { ctx.BeginList() }
func (_this *TopLevelRule) OnMap(ctx *Context)       { ctx.BeginMap() }
func (_this *TopLevelRule) OnMarkup(ctx *Context)    { ctx.BeginMarkup() }
func (_this *TopLevelRule) OnMetadata(ctx *Context)  { ctx.BeginMetadata() }
func (_this *TopLevelRule) OnComment(ctx *Context)   { ctx.BeginComment() }
func (_this *TopLevelRule) OnMarker(ctx *Context)    { ctx.BeginMarkerAnyType() }
func (_this *TopLevelRule) OnReference(ctx *Context) { ctx.BeginTopLevelReference() }
func (_this *TopLevelRule) OnConstant(ctx *Context, name []byte, explicitValue bool) {
	ctx.BeginConstantAnyType(name, explicitValue)
}
func (_this *TopLevelRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.ValidateFullArrayAnyType(arrayType, elementCount, data)
	ctx.SwitchEndDocument()
}
func (_this *TopLevelRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.ValidateFullArrayStringlike(arrayType, data)
	ctx.SwitchEndDocument()
}
func (_this *TopLevelRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginArrayAnyType(arrayType)
}

// =============================================================================

type NACatRule struct{}

func (_this *NACatRule) String() string                                          { return "NA (Cat) Rule" }
func (_this *NACatRule) OnKeyableObject(ctx *Context)                            { ctx.UnstackRule() }
func (_this *NACatRule) OnNonKeyableObject(ctx *Context)                         { ctx.UnstackRule() }
func (_this *NACatRule) OnNA(ctx *Context)                                       { ctx.UnstackRule() }
func (_this *NACatRule) OnChildContainerEnded(ctx *Context, _ DataType)          { ctx.UnstackRule() }
func (_this *NACatRule) OnPadding(ctx *Context)                                  { /* Nothing to do */ }
func (_this *NACatRule) OnInt(ctx *Context, value int64)                         { ctx.UnstackRule() }
func (_this *NACatRule) OnPositiveInt(ctx *Context, value uint64)                { ctx.UnstackRule() }
func (_this *NACatRule) OnBigInt(ctx *Context, value *big.Int)                   { ctx.UnstackRule() }
func (_this *NACatRule) OnFloat(ctx *Context, value float64)                     { ctx.UnstackRule() }
func (_this *NACatRule) OnBigFloat(ctx *Context, value *big.Float)               { ctx.UnstackRule() }
func (_this *NACatRule) OnDecimalFloat(ctx *Context, value compact_float.DFloat) { ctx.UnstackRule() }
func (_this *NACatRule) OnBigDecimalFloat(ctx *Context, value *apd.Decimal)      { ctx.UnstackRule() }
func (_this *NACatRule) OnList(ctx *Context)                                     { ctx.BeginList() }
func (_this *NACatRule) OnMap(ctx *Context)                                      { ctx.BeginMap() }
func (_this *NACatRule) OnMarkup(ctx *Context)                                   { ctx.BeginMarkup() }
func (_this *NACatRule) OnMetadata(ctx *Context)                                 { ctx.BeginMetadata() }
func (_this *NACatRule) OnComment(ctx *Context)                                  { ctx.BeginComment() }
func (_this *NACatRule) OnMarker(ctx *Context)                                   { ctx.BeginMarkerAnyType() }
func (_this *NACatRule) OnReference(ctx *Context)                                { ctx.BeginTopLevelReference() }
func (_this *NACatRule) OnConstant(ctx *Context, name []byte, explicitValue bool) {
	ctx.BeginConstantAnyType(name, explicitValue)
}
func (_this *NACatRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.ValidateFullArrayAnyType(arrayType, elementCount, data)
	ctx.UnstackRule()
}
func (_this *NACatRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.ValidateFullArrayStringlike(arrayType, data)
	ctx.UnstackRule()
}
func (_this *NACatRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginArrayAnyType(arrayType)
}

// =============================================================================

type ConstantKeyableRule struct{}

func (_this *ConstantKeyableRule) String() string         { return "Keyable Constant Rule" }
func (_this *ConstantKeyableRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *ConstantKeyableRule) OnKeyableObject(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnKeyableObject(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnInt(ctx *Context, value int64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnPositiveInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnBigInt(ctx *Context, value *big.Int) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnFloat(ctx *Context, value float64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnBigFloat(ctx *Context, value *big.Float) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnDecimalFloat(ctx *Context, value compact_float.DFloat) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnBigDecimalFloat(ctx *Context, value *apd.Decimal) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnReference(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnReference(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnArray(ctx, arrayType, elementCount, data)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnStringlikeArray(ctx, arrayType, data)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantKeyableRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.BeginArrayKeyable(arrayType)
}
func (_this *ConstantKeyableRule) OnChildContainerEnded(ctx *Context, _ DataType) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, DataTypeKeyable)
}

// =============================================================================

type ConstantAnyTypeRule struct{}

func (_this *ConstantAnyTypeRule) String() string         { return "Constant Rule" }
func (_this *ConstantAnyTypeRule) OnPadding(ctx *Context) { /* Nothing to do */ }
func (_this *ConstantAnyTypeRule) OnNonKeyableObject(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnKeyableObject(ctx)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *ConstantAnyTypeRule) OnNA(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnNA(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnKeyableObject(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnKeyableObject(ctx)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnInt(ctx *Context, value int64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnPositiveInt(ctx *Context, value uint64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnPositiveInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnBigInt(ctx *Context, value *big.Int) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigInt(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnFloat(ctx *Context, value float64) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnBigFloat(ctx *Context, value *big.Float) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnDecimalFloat(ctx *Context, value compact_float.DFloat) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnBigDecimalFloat(ctx *Context, value *apd.Decimal) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnBigDecimalFloat(ctx, value)
	ctx.MarkObject(DataTypeKeyable)
}
func (_this *ConstantAnyTypeRule) OnList(ctx *Context) {
	ctx.ParentRule().OnList(ctx)
}
func (_this *ConstantAnyTypeRule) OnMap(ctx *Context) {
	ctx.ParentRule().OnMap(ctx)
}
func (_this *ConstantAnyTypeRule) OnMarkup(ctx *Context) {
	ctx.ParentRule().OnMarkup(ctx)
}
func (_this *ConstantAnyTypeRule) OnReference(ctx *Context) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnReference(ctx)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *ConstantAnyTypeRule) OnArray(ctx *Context, arrayType events.ArrayType, elementCount uint64, data []uint8) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnArray(ctx, arrayType, elementCount, data)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *ConstantAnyTypeRule) OnStringlikeArray(ctx *Context, arrayType events.ArrayType, data string) {
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnStringlikeArray(ctx, arrayType, data)
	ctx.MarkObject(DataTypeAnyType)
}
func (_this *ConstantAnyTypeRule) OnArrayBegin(ctx *Context, arrayType events.ArrayType) {
	ctx.ParentRule().OnArrayBegin(ctx, arrayType)
}
func (_this *ConstantAnyTypeRule) OnChildContainerEnded(ctx *Context, cType DataType) {
	ctx.MarkObject(cType)
	ctx.UnstackRule()
	ctx.CurrentEntry.Rule.OnChildContainerEnded(ctx, cType)
}
