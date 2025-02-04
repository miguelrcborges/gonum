package gonum

import (
	"fmt"
	"unsafe"
	"testing"

	"github.com/miguelrcborges/gonum/common"
	"github.com/miguelrcborges/gonum/amd64"
)

func getInt64BenchmarkParams(size uint64) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint64) {
	v1, v2, vo := getBigInt64Slice(size, size)
	return unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), size
}


func BenchmarkInt64Adds(b *testing.B) {
    lengths := []int{6, 14, 52, 168, 552}
    for _, length := range lengths {
        b.Run(fmt.Sprintf("Int64Slice_NaiveAdd_%d", length), func(b *testing.B) {
            p1, p2, po, l := getInt64BenchmarkParams(uint64(length))
            for i := 0; i < b.N; i++ {
                common.AddInt64(p1, p2, po, l)
            }
        })

        b.Run(fmt.Sprintf("Int64Slice_AsmScalarAdd_%d", length), func(b *testing.B) {
            p1, p2, po, l := getInt64BenchmarkParams(uint64(length))
            for i := 0; i < b.N; i++ {
                amd64.AddInt64Scalar(p1, p2, po, l)
            }
        })

        b.Run(fmt.Sprintf("Int64Slice_SSE2Add_%d", length), func(b *testing.B) {
            p1, p2, po, l := getInt64BenchmarkParams(uint64(length))
            for i := 0; i < b.N; i++ {
                amd64.AddInt64SSE2(p1, p2, po, l)
            }
        })

        b.Run(fmt.Sprintf("Int64Slice_AVX2Add_%d", length), func(b *testing.B) {
            p1, p2, po, l := getInt64BenchmarkParams(uint64(length))
            for i := 0; i < b.N; i++ {
                amd64.AddInt64AVX2(p1, p2, po, l)
            }
        })
    }
}
