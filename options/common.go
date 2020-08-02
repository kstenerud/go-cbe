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

// All options that can be used to fine-tune the behavior of various aspects of
// this library.
package options

// TODO: Opt: Convert line endings to escapes
// TODO: Opt: Don't convert escapes
// TODO: Builder that converts to string
// TODO: Iterator that converts from string to smaller type (numeric)
// TODO: Some method to notify that a string field should be encoded as a different type

// Defines what (if any) implied structure forms the Concise Encoding document takes.
type ImpliedStructure int

const (
	// Assume no implied structure
	ImpliedStructureNone = iota

	//Assume that the Concise Encoding version has already been specified.
	ImpliedStructureVersion

	// Assume that the version has been specified AND that the top-level object
	// has already been opened as a list.
	ImpliedStructureList

	// Assume that the version has been specified AND that the top-level object
	// has already been opened as a map.
	ImpliedStructureMap
)
