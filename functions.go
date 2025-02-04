package gonum

import (
	"unsafe"

	"github.com/miguelrcborges/gonum/common"
)

var _AddInt64 func(v1, v2, vo unsafe.Pointer, l int64) = common.AddInt64
func AddInt64(v1, v2 []int64) []int64 {
	if len(v1) != len(v2) {
		panic(invalid_slices_combination)
	}
	l := len(v1)
	vo := make([]int64, len(v1))
	if l == 0 {
		return vo
	}
	_AddInt64(unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), int64(l))
	return vo
}
