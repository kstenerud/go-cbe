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
	"math/big"
	"net/url"
	"reflect"
	"time"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
)

type arrayBuilder struct {
	// Const data
	dstType reflect.Type

	// Cloned data (must be populated)
	elemBuilder ObjectBuilder

	// Clone inserted data
	root   *RootBuilder
	parent ObjectBuilder

	// Variable data (must be reset)
	container reflect.Value
	index     int
}

func newArrayBuilder(dstType reflect.Type) ObjectBuilder {
	return &arrayBuilder{
		dstType: dstType,
	}
}

func (_this *arrayBuilder) IsContainerOnly() bool {
	return true
}

func (_this *arrayBuilder) PostCacheInitBuilder() {
	_this.elemBuilder = getBuilderForType(_this.dstType.Elem())
}

func (_this *arrayBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	that := &arrayBuilder{
		dstType: _this.dstType,
		parent:  parent,
		root:    root,
	}
	that.elemBuilder = _this.elemBuilder.CloneFromTemplate(root, that, options)
	that.reset()
	return that
}

func (_this *arrayBuilder) SetParent(parent ObjectBuilder) {
	_this.parent = parent
}

func (_this *arrayBuilder) reset() {
	_this.container = reflect.New(_this.dstType).Elem()
	_this.index = 0
}

func (_this *arrayBuilder) currentElem() reflect.Value {
	return _this.container.Index(_this.index)
}

func (_this *arrayBuilder) BuildFromNil(ignored reflect.Value) {
	_this.elemBuilder.BuildFromNil(_this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromBool(value bool, ignored reflect.Value) {
	_this.elemBuilder.BuildFromBool(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromInt(value int64, ignored reflect.Value) {
	_this.elemBuilder.BuildFromInt(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromUint(value uint64, ignored reflect.Value) {
	_this.elemBuilder.BuildFromUint(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromBigInt(value *big.Int, ignored reflect.Value) {
	_this.elemBuilder.BuildFromBigInt(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromFloat(value float64, ignored reflect.Value) {
	_this.elemBuilder.BuildFromFloat(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromBigFloat(value *big.Float, ignored reflect.Value) {
	_this.elemBuilder.BuildFromBigFloat(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromDecimalFloat(value compact_float.DFloat, ignored reflect.Value) {
	_this.elemBuilder.BuildFromDecimalFloat(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, ignored reflect.Value) {
	_this.elemBuilder.BuildFromBigDecimalFloat(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromUUID(value []byte, ignored reflect.Value) {
	_this.elemBuilder.BuildFromUUID(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromString(value string, ignored reflect.Value) {
	_this.elemBuilder.BuildFromString(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	_this.elemBuilder.BuildFromBytes(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromURI(value *url.URL, ignored reflect.Value) {
	_this.elemBuilder.BuildFromURI(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromTime(value time.Time, ignored reflect.Value) {
	_this.elemBuilder.BuildFromTime(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildFromCompactTime(value *compact_time.Time, ignored reflect.Value) {
	_this.elemBuilder.BuildFromCompactTime(value, _this.currentElem())
	_this.index++
}

func (_this *arrayBuilder) BuildBeginList() {
	_this.elemBuilder.PrepareForListContents()
}

func (_this *arrayBuilder) BuildBeginMap() {
	_this.elemBuilder.PrepareForMapContents()
}

func (_this *arrayBuilder) BuildEndContainer() {
	object := _this.container
	_this.reset()
	_this.parent.NotifyChildContainerFinished(object)
}

func (_this *arrayBuilder) BuildFromMarker(id interface{}) {
	panic("TODO: arrayBuilder.BuildFromMarker")
}

func (_this *arrayBuilder) BuildFromReference(id interface{}) {
	panic("TODO: arrayBuilder.BuildFromReference")
}

func (_this *arrayBuilder) PrepareForListContents() {
	_this.root.SetCurrentBuilder(_this)
}

func (_this *arrayBuilder) PrepareForMapContents() {
	builderPanicBadEvent(_this, _this.dstType, "PrepareForMapContents")
}

func (_this *arrayBuilder) NotifyChildContainerFinished(value reflect.Value) {
	_this.root.SetCurrentBuilder(_this)
	_this.currentElem().Set(value)
	_this.index++
}

// ============================================================================

type bytesArrayBuilder struct {
}

var globalBytesArrayBuilder bytesArrayBuilder

func newBytesArrayBuilder() ObjectBuilder {
	return &globalBytesArrayBuilder
}

func (_this *bytesArrayBuilder) IsContainerOnly() bool {
	return false
}

func (_this *bytesArrayBuilder) PostCacheInitBuilder() {
}

func (_this *bytesArrayBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	return _this
}

func (_this *bytesArrayBuilder) SetParent(parent ObjectBuilder) {
}

func (_this *bytesArrayBuilder) BuildFromNil(ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromNil")
}

func (_this *bytesArrayBuilder) BuildFromBool(value bool, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromBool")
}

func (_this *bytesArrayBuilder) BuildFromInt(value int64, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromInt")
}

func (_this *bytesArrayBuilder) BuildFromUint(value uint64, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromUint")
}

func (_this *bytesArrayBuilder) BuildFromBigInt(value *big.Int, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromBigInt")
}

func (_this *bytesArrayBuilder) BuildFromFloat(value float64, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromFloat")
}

func (_this *bytesArrayBuilder) BuildFromBigFloat(value *big.Float, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromBigFloat")
}

func (_this *bytesArrayBuilder) BuildFromDecimalFloat(value compact_float.DFloat, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromDecimalFloat")
}

func (_this *bytesArrayBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromBigDecimalFloat")
}

func (_this *bytesArrayBuilder) BuildFromUUID(value []byte, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromUUID")
}

func (_this *bytesArrayBuilder) BuildFromString(value string, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromString")
}

func (_this *bytesArrayBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	// TODO: Is there a more efficient way?
	for i := 0; i < len(value); i++ {
		elem := dst.Index(i)
		elem.SetUint(uint64(value[i]))
	}
}

func (_this *bytesArrayBuilder) BuildFromURI(value *url.URL, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromURI")
}

func (_this *bytesArrayBuilder) BuildFromTime(value time.Time, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromTime")
}

func (_this *bytesArrayBuilder) BuildFromCompactTime(value *compact_time.Time, ignored reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "BuildFromCompactTime")
}

func (_this *bytesArrayBuilder) BuildBeginList() {
	builderPanicBadEvent(_this, typeBytes, "BuildBeginList")
}

func (_this *bytesArrayBuilder) BuildBeginMap() {
	builderPanicBadEvent(_this, typeBytes, "BuildBeginMap")
}

func (_this *bytesArrayBuilder) BuildEndContainer() {
	builderPanicBadEvent(_this, typeBytes, "BuildEndContainer")
}

func (_this *bytesArrayBuilder) BuildFromMarker(id interface{}) {
	panic("TODO: bytesArrayBuilder.BuildFromMarker")
}

func (_this *bytesArrayBuilder) BuildFromReference(id interface{}) {
	panic("TODO: bytesArrayBuilder.BuildFromReference")
}

func (_this *bytesArrayBuilder) PrepareForListContents() {
	builderPanicBadEvent(_this, typeBytes, "PrepareForListContents")
}

func (_this *bytesArrayBuilder) PrepareForMapContents() {
	builderPanicBadEvent(_this, typeBytes, "PrepareForMapContents")
}

func (_this *bytesArrayBuilder) NotifyChildContainerFinished(value reflect.Value) {
	builderPanicBadEvent(_this, typeBytes, "NotifyChildContainerFinished")
}
