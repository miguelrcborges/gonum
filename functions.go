package gonum

import (
	"unsafe"

	"github.com/miguelrcborges/gonum/common"
)

var UnsafeAddInt64 func(v1, v2, vo unsafe.Pointer, l uint64) = common.AddInt64
func AddInt64(v1, v2, vo []int64) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeAddInt64(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}

func AddUint64(v1, v2, vo []uint64) {
	if len(v1) != len(v2) || len(v1) != len(vo) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	if l != 0 {
		UnsafeAddInt64(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(l))
	}
}
