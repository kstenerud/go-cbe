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

	"github.com/kstenerud/go-concise-encoding/options"

	"github.com/cockroachdb/apd/v2"
	"github.com/kstenerud/go-compact-float"
	"github.com/kstenerud/go-compact-time"
)

type uintBuilder struct {
	// Static data
	session *Session
	dstType reflect.Type
}

func newUintBuilder(dstType reflect.Type) ObjectBuilder {
	return &uintBuilder{
		dstType: dstType,
	}
}

func (_this *uintBuilder) String() string {
	return fmt.Sprintf("%v<%v>", reflect.TypeOf(_this), _this.dstType)
}

func (_this *uintBuilder) InitTemplate(session *Session) {
	_this.session = session
}

func (_this *uintBuilder) NewInstance(root *RootBuilder, parent ObjectBuilder, options *options.BuilderOptions) ObjectBuilder {
	return _this
}

func (_this *uintBuilder) SetParent(parent ObjectBuilder) {
}

func (_this *uintBuilder) BuildFromNil(dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Nil")
}

func (_this *uintBuilder) BuildFromBool(value bool, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Bool")
}

func (_this *uintBuilder) BuildFromInt(value int64, dst reflect.Value) {
	setUintFromInt(value, dst)
}

func (_this *uintBuilder) BuildFromUint(value uint64, dst reflect.Value) {
	setUintFromUint(value, dst)
}

func (_this *uintBuilder) BuildFromBigInt(value *big.Int, dst reflect.Value) {
	setUintFromBigInt(value, dst)
}

func (_this *uintBuilder) BuildFromFloat(value float64, dst reflect.Value) {
	setUintFromFloat(value, dst)
}

func (_this *uintBuilder) BuildFromBigFloat(value *big.Float, dst reflect.Value) {
	setUintFromBigFloat(value, dst)
}

func (_this *uintBuilder) BuildFromDecimalFloat(value compact_float.DFloat, dst reflect.Value) {
	setUintFromDecimalFloat(value, dst)
}

func (_this *uintBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, dst reflect.Value) {
	setUintFromBigDecimalFloat(value, dst)
}

func (_this *uintBuilder) BuildFromUUID(value []byte, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "UUID")
}

func (_this *uintBuilder) BuildFromString(value []byte, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "String")
}

func (_this *uintBuilder) BuildFromVerbatimString(value []byte, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "VerbatimString")
}

func (_this *uintBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Bytes")
}

func (_this *uintBuilder) BuildFromCustomBinary(value []byte, dst reflect.Value) {
	if err := _this.session.GetCustomBinaryBuildFunction()(value, dst); err != nil {
		BuilderPanicBuildFromCustomBinary(_this, value, dst.Type(), err)
	}
}

func (_this *uintBuilder) BuildFromCustomText(value []byte, dst reflect.Value) {
	if err := _this.session.GetCustomTextBuildFunction()(value, dst); err != nil {
		BuilderPanicBuildFromCustomText(_this, value, dst.Type(), err)
	}
}

func (_this *uintBuilder) BuildFromURI(value *url.URL, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "URI")
}

func (_this *uintBuilder) BuildFromTime(value time.Time, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Time")
}

func (_this *uintBuilder) BuildFromCompactTime(value *compact_time.Time, dst reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "CompactTime")
}

func (_this *uintBuilder) BuildBeginList() {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "List")
}

func (_this *uintBuilder) BuildBeginMap() {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Map")
}

func (_this *uintBuilder) BuildEndContainer() {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "ContainerEnd")
}

func (_this *uintBuilder) BuildBeginMarker(id interface{}) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Marker")
}

func (_this *uintBuilder) BuildFromReference(id interface{}) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "Reference")
}

func (_this *uintBuilder) PrepareForListContents() {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "PrepareForListContents")
}

func (_this *uintBuilder) PrepareForMapContents() {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "PrepareForMapContents")
}

func (_this *uintBuilder) NotifyChildContainerFinished(value reflect.Value) {
	BuilderWithTypePanicBadEvent(_this, _this.dstType, "NotifyChildContainerFinished")
}
