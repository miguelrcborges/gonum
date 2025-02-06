package gonum

import (
	"unsafe"

	"github.com/miguelrcborges/gonum/common"
)

var UnsafeSubInt64 func(v1, v2, vo unsafe.Pointer, l uint64) = common.SubInt64
func SubInt64(v1, v2, vo []int64) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt64(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}

func SubUint64(v1, v2, vo []uint64) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt64(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}


var UnsafeSubInt32 func(v1, v2, vo unsafe.Pointer, l uint64) = common.SubInt32
func SubInt32(v1, v2, vo []int32) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt32(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}

func SubUint32(v1, v2, vo []uint32) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt32(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}


var UnsafeSubInt16 func(v1, v2, vo unsafe.Pointer, l uint64) = common.SubInt16
func SubInt16(v1, v2, vo []int16) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt16(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}

func SubUint16(v1, v2, vo []uint16) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt16(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}


var UnsafeSubInt8 func(v1, v2, vo unsafe.Pointer, l uint64) = common.SubInt8
func SubInt8(v1, v2, vo []int8) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt8(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}

func SubUint8(v1, v2, vo []uint8) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeSubInt8(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}
