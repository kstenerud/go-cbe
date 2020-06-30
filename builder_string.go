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

type stringBuilder struct {
}

var globalStringBuilder = &stringBuilder{}

func newStringBuilder() ObjectBuilder {
	return globalStringBuilder
}

func (_this *stringBuilder) IsContainerOnly() bool {
	return false
}

func (_this *stringBuilder) PostCacheInitBuilder() {
}

func (_this *stringBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	return _this
}

func (_this *stringBuilder) SetParent(parent ObjectBuilder) {
}

func (_this *stringBuilder) BuildFromNil(dst reflect.Value) {
	// Go doesn't have the concept of a nil string.
	dst.SetString("")
}

func (_this *stringBuilder) BuildFromBool(value bool, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "Bool")
}

func (_this *stringBuilder) BuildFromInt(value int64, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "Int")
}

func (_this *stringBuilder) BuildFromUint(value uint64, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "Uint")
}

func (_this *stringBuilder) BuildFromBigInt(value *big.Int, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "BigInt")
}

func (_this *stringBuilder) BuildFromFloat(value float64, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "Float")
}

func (_this *stringBuilder) BuildFromBigFloat(value *big.Float, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "BigFloat")
}

func (_this *stringBuilder) BuildFromDecimalFloat(value compact_float.DFloat, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "DecimalFloat")
}

func (_this *stringBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "BigDecimalFloat")
}

func (_this *stringBuilder) BuildFromUUID(value []byte, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "UUID")
}

func (_this *stringBuilder) BuildFromString(value string, dst reflect.Value) {
	dst.SetString(value)
}

func (_this *stringBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "Bytes")
}

func (_this *stringBuilder) BuildFromURI(value *url.URL, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "URI")
}

func (_this *stringBuilder) BuildFromTime(value time.Time, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "Time")
}

func (_this *stringBuilder) BuildFromCompactTime(value *compact_time.Time, dst reflect.Value) {
	builderPanicBadEvent(_this, typeString, "CompactTime")
}

func (_this *stringBuilder) BuildBeginList() {
	builderPanicBadEvent(_this, typeString, "List")
}

func (_this *stringBuilder) BuildBeginMap() {
	builderPanicBadEvent(_this, typeString, "Map")
}

func (_this *stringBuilder) BuildEndContainer() {
	builderPanicBadEvent(_this, typeString, "ContainerEnd")
}

func (_this *stringBuilder) BuildFromMarker(id interface{}) {
	panic("TODO: stringBuilder.Marker")
}

func (_this *stringBuilder) BuildFromReference(id interface{}) {
	panic("TODO: stringBuilder.Reference")
}

func (_this *stringBuilder) PrepareForListContents() {
	builderPanicBadEvent(_this, typeString, "PrepareForListContents")
}

func (_this *stringBuilder) PrepareForMapContents() {
	builderPanicBadEvent(_this, typeString, "PrepareForMapContents")
}

func (_this *stringBuilder) NotifyChildContainerFinished(value reflect.Value) {
	builderPanicBadEvent(_this, typeString, "NotifyChildContainerFinished")
}
