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

type ignoreBuilder struct {
	// Clone inserted data
	root    *RootBuilder
	parent  ObjectBuilder
	options *BuilderOptions
}

var globalIgnoreBuilder = &ignoreBuilder{}

func newIgnoreBuilder() ObjectBuilder {
	return globalIgnoreBuilder
}

func (this *ignoreBuilder) IsContainerOnly() bool {
	return false
}

func (this *ignoreBuilder) PostCacheInitBuilder() {
}

func (this *ignoreBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	return &ignoreBuilder{
		parent:  parent,
		root:    root,
		options: options,
	}
}

func (this *ignoreBuilder) BuildFromNil(dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromBool(value bool, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromInt(value int64, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromUint(value uint64, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromBigInt(value *big.Int, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromFloat(value float64, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromBigFloat(value *big.Float, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromDecimalFloat(value compact_float.DFloat, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromUUID(value []byte, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromString(value string, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromURI(value *url.URL, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromTime(value time.Time, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildFromCompactTime(value *compact_time.Time, dst reflect.Value) {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreBuilder) BuildBeginList() {
	builder := newIgnoreContainerBuilder().CloneFromTemplate(this.root, this.parent, this.options)
	builder.PrepareForListContents()
}

func (this *ignoreBuilder) BuildBeginMap() {
	builder := newIgnoreContainerBuilder().CloneFromTemplate(this.root, this.parent, this.options)
	builder.PrepareForMapContents()
}

func (this *ignoreBuilder) BuildEndContainer() {
	builderPanicBadEvent(this, reflect.TypeOf([]interface{}{}).Elem(), "End")
}

func (this *ignoreBuilder) BuildFromMarker(id interface{}) {
	panic("TODO: ignoreBuilder.Marker")
}

func (this *ignoreBuilder) BuildFromReference(id interface{}) {
	panic("TODO: ignoreBuilder.Reference")
}

func (this *ignoreBuilder) PrepareForListContents() {
	builderPanicBadEvent(this, typePBigInt, "PrepareForListContents")
}

func (this *ignoreBuilder) PrepareForMapContents() {
	builderPanicBadEvent(this, typePBigInt, "PrepareForMapContents")
}

func (this *ignoreBuilder) NotifyChildContainerFinished(value reflect.Value) {
	builderPanicBadEvent(this, typePBigInt, "NotifyChildContainerFinished")
}

// ============================================================================

type ignoreContainerBuilder struct {
	// Clone inserted data
	root    *RootBuilder
	parent  ObjectBuilder
	options *BuilderOptions
}

var globalIgnoreContainerBuilder = &ignoreContainerBuilder{}

func newIgnoreContainerBuilder() ObjectBuilder {
	return globalIgnoreContainerBuilder
}

func (this *ignoreContainerBuilder) IsContainerOnly() bool {
	return true
}

func (this *ignoreContainerBuilder) PostCacheInitBuilder() {
}

func (this *ignoreContainerBuilder) CloneFromTemplate(root *RootBuilder, parent ObjectBuilder, options *BuilderOptions) ObjectBuilder {
	return &ignoreContainerBuilder{
		parent:  parent,
		root:    root,
		options: options,
	}
}

func (this *ignoreContainerBuilder) BuildFromNil(dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromBool(value bool, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromInt(value int64, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromUint(value uint64, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromBigInt(value *big.Int, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromFloat(value float64, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromBigFloat(value *big.Float, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromDecimalFloat(value compact_float.DFloat, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromBigDecimalFloat(value *apd.Decimal, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromUUID(value []byte, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromString(value string, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromBytes(value []byte, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromURI(value *url.URL, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromTime(value time.Time, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildFromCompactTime(value *compact_time.Time, dst reflect.Value) {
	// Ignore this directive
}

func (this *ignoreContainerBuilder) BuildBeginList() {
	builder := newIgnoreContainerBuilder().CloneFromTemplate(this.root, this, this.options)
	builder.PrepareForListContents()
}

func (this *ignoreContainerBuilder) BuildBeginMap() {
	builder := newIgnoreContainerBuilder().CloneFromTemplate(this.root, this, this.options)
	builder.PrepareForMapContents()
}

func (this *ignoreContainerBuilder) BuildEndContainer() {
	this.root.setCurrentBuilder(this.parent)
}

func (this *ignoreContainerBuilder) BuildFromMarker(id interface{}) {
	panic("TODO: ignoreContainerBuilder.Marker")
}

func (this *ignoreContainerBuilder) BuildFromReference(id interface{}) {
	panic("TODO: ignoreContainerBuilder.Reference")
}

func (this *ignoreContainerBuilder) PrepareForListContents() {
	this.root.setCurrentBuilder(this)
}

func (this *ignoreContainerBuilder) PrepareForMapContents() {
	this.root.setCurrentBuilder(this)
}

func (this *ignoreContainerBuilder) NotifyChildContainerFinished(value reflect.Value) {
}
