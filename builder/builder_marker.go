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

package builder

import (
	"fmt"
	"math/big"
	"net/url"
	"reflect"
	"time"

	"github.com/kstenerud/go-concise-encoding/internal/common"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
)

type markerIDBuilder struct {
	onID func(interface{})
}

func newMarkerIDBuilder(onID func(interface{})) *markerIDBuilder {
	return &markerIDBuilder{
		onID: onID,
	}
}

func (_this *markerIDBuilder) String() string {
	return fmt.Sprintf("%v", reflect.TypeOf(_this))
}

func (_this *markerIDBuilder) IsContainerOnly() bool {
	return false
}

func (_this *markerIDBuilder) PostCacheInitBuilder() {
}

func (_this *markerIDBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	builderPanicBadEvent(_this, "CloneFromTemplate")
	return nil
}

func (_this *markerIDBuilder) SetParent(parent ObjectBuilder) {
}

func (_this *markerIDBuilder) BuildFromNil(dst reflect.Value) {
	builderPanicBadEvent(_this, "Nil")
}

func (_this *markerIDBuilder) BuildFromBool(value bool, dst reflect.Value) {
	builderPanicBadEvent(_this, "Bool")
}

func (_this *markerIDBuilder) BuildFromInt(value int64, dst reflect.Value) {
	if value < 0 {
		builderPanicBadEvent(_this, "Int")
	}
	_this.onID(value)
}

func (_this *markerIDBuilder) BuildFromUint(value uint64, dst reflect.Value) {
	_this.onID(value)
}

func (_this *markerIDBuilder) BuildFromBigInt(value *big.Int, dst reflect.Value) {
	if common.IsBigIntNegative(value) || !value.IsUint64() {
		builderPanicBadEvent(_this, "BigInt")
	}
	_this.onID(value.Uint64())
}

func (_this *markerIDBuilder) BuildFromFloat(value float64, dst reflect.Value) {
	builderPanicBadEvent(_this, "Float")
}

func (_this *markerIDBuilder) BuildFromBigFloat(value *big.Float, dst reflect.Value) {
	builderPanicBadEvent(_this, "BigFloat")
}

func (_this *markerIDBuilder) BuildFromDecimalFloat(value compact_float.DFloat, dst reflect.Value) {
	builderPanicBadEvent(_this, "DecimalFloat")
}

func (_this *markerIDBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, dst reflect.Value) {
	builderPanicBadEvent(_this, "BigDecimalFloat")
}

func (_this *markerIDBuilder) BuildFromUUID(value []byte, dst reflect.Value) {
	builderPanicBadEvent(_this, "UUID")
}

func (_this *markerIDBuilder) BuildFromString(value string, dst reflect.Value) {
	_this.onID(value)
}

func (_this *markerIDBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	builderPanicBadEvent(_this, "Bytes")
}

func (_this *markerIDBuilder) BuildFromURI(value *url.URL, dst reflect.Value) {
	builderPanicBadEvent(_this, "URI")
}

func (_this *markerIDBuilder) BuildFromTime(value time.Time, dst reflect.Value) {
	builderPanicBadEvent(_this, "Time")
}

func (_this *markerIDBuilder) BuildFromCompactTime(value *compact_time.Time, dst reflect.Value) {
	builderPanicBadEvent(_this, "CompactTime")
}

func (_this *markerIDBuilder) BuildBeginList() {
	builderPanicBadEvent(_this, "List")
}

func (_this *markerIDBuilder) BuildBeginMap() {
	builderPanicBadEvent(_this, "Map")
}

func (_this *markerIDBuilder) BuildEndContainer() {
	builderPanicBadEvent(_this, "End")
}

func (_this *markerIDBuilder) BuildBeginMarker(id interface{}) {
	builderPanicBadEvent(_this, "Marker")
}

func (_this *markerIDBuilder) BuildFromReference(id interface{}) {
	builderPanicBadEvent(_this, "Reference")
}

func (_this *markerIDBuilder) PrepareForListContents() {
	builderPanicBadEvent(_this, "PrepareForListContents")
}

func (_this *markerIDBuilder) PrepareForMapContents() {
	builderPanicBadEvent(_this, "PrepareForMapContents")
}

func (_this *markerIDBuilder) NotifyChildContainerFinished(value reflect.Value) {
	builderPanicBadEvent(_this, "NotifyChildContainerFinished")
}

// ============================================================================

type markerObjectBuilder struct {
	parent           ObjectBuilder
	child            ObjectBuilder
	onObjectComplete func(reflect.Value)
}

func newMarkerObjectBuilder(parent, child ObjectBuilder, onObjectComplete func(reflect.Value)) *markerObjectBuilder {
	return &markerObjectBuilder{
		parent:           parent,
		child:            child,
		onObjectComplete: onObjectComplete,
	}
}

func (_this *markerObjectBuilder) String() string {
	return fmt.Sprintf("%v<%v>", reflect.TypeOf(_this), _this.child)
}

func (_this *markerObjectBuilder) IsContainerOnly() bool {
	return false
}

func (_this *markerObjectBuilder) PostCacheInitBuilder() {
}

func (_this *markerObjectBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	builderPanicBadEvent(_this, "CloneFromTemplate")
	return nil
}

func (_this *markerObjectBuilder) SetParent(parent ObjectBuilder) {
	_this.parent = parent
}

func (_this *markerObjectBuilder) BuildFromNil(dst reflect.Value) {
	_this.child.BuildFromNil(dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromBool(value bool, dst reflect.Value) {
	_this.child.BuildFromBool(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromInt(value int64, dst reflect.Value) {
	_this.child.BuildFromInt(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromUint(value uint64, dst reflect.Value) {
	_this.child.BuildFromUint(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromBigInt(value *big.Int, dst reflect.Value) {
	_this.child.BuildFromBigInt(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromFloat(value float64, dst reflect.Value) {
	_this.child.BuildFromFloat(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromBigFloat(value *big.Float, dst reflect.Value) {
	_this.child.BuildFromBigFloat(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromDecimalFloat(value compact_float.DFloat, dst reflect.Value) {
	_this.child.BuildFromDecimalFloat(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, dst reflect.Value) {
	_this.child.BuildFromBigDecimalFloat(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromUUID(value []byte, dst reflect.Value) {
	_this.child.BuildFromUUID(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromString(value string, dst reflect.Value) {
	_this.child.BuildFromString(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	_this.child.BuildFromBytes(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromURI(value *url.URL, dst reflect.Value) {
	_this.child.BuildFromURI(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromTime(value time.Time, dst reflect.Value) {
	_this.child.BuildFromTime(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildFromCompactTime(value *compact_time.Time, dst reflect.Value) {
	_this.child.BuildFromCompactTime(value, dst)
	_this.onObjectComplete(dst)
}

func (_this *markerObjectBuilder) BuildBeginList() {
	builderPanicBadEvent(_this, "List")
}

func (_this *markerObjectBuilder) BuildBeginMap() {
	builderPanicBadEvent(_this, "Map")
}

func (_this *markerObjectBuilder) BuildEndContainer() {
	builderPanicBadEvent(_this, "End")
}

func (_this *markerObjectBuilder) BuildBeginMarker(id interface{}) {
	builderPanicBadEvent(_this, "Marker")
}

func (_this *markerObjectBuilder) BuildFromReference(id interface{}) {
	builderPanicBadEvent(_this, "Reference")
}

func (_this *markerObjectBuilder) PrepareForListContents() {
	_this.child.SetParent(_this)
	_this.child.PrepareForListContents()
}

func (_this *markerObjectBuilder) PrepareForMapContents() {
	_this.child.SetParent(_this)
	_this.child.PrepareForMapContents()
}

func (_this *markerObjectBuilder) NotifyChildContainerFinished(value reflect.Value) {
	_this.onObjectComplete(value)
	_this.child.SetParent(_this.parent)
	_this.parent.NotifyChildContainerFinished(value)
}