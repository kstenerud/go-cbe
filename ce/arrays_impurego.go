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

// +build !purego

package ce

import (
	"unsafe"

	"github.com/kstenerud/go-concise-encoding/internal/common"
)

var isLittleEndian bool

func init() {
	v := uint16(1)
	bytes := (*[common.AddressSpace]byte)(unsafe.Pointer(&v))[:2]
	if bytes[1] == 1 {
		isLittleEndian = true
	}
}

func Int8SliceToBytes(data []int8) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)]
	}
	return int8SliceToBytes(data)
}

func Uint16SliceToBytes(data []uint16) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)*2]
	}
	return uint16SliceToBytes(data)
}

func Int16SliceToBytes(data []int16) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)*2]
	}
	return int16SliceToBytes(data)
}

func Uint32SliceToBytes(data []uint32) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)*4]
	}
	return uint32SliceToBytes(data)
}

func Int32SliceToBytes(data []int32) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)*4]
	}
	return int32SliceToBytes(data)
}

func Uint64SliceToBytes(data []uint64) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)*8]
	}
	return uint64SliceToBytes(data)
}

func Int64SliceToBytes(data []int64) []byte {
	if isLittleEndian {
		return (*[common.AddressSpace]byte)(unsafe.Pointer(&data[0]))[:len(data)*8]
	}
	return int64SliceToBytes(data)
}
