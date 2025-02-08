package gonum

import (
	"fmt"
	"unsafe"
	"testing"

	"github.com/miguelrcborges/gonum/common"
	"github.com/miguelrcborges/gonum/amd64"
)

func getBenchmarkParams[T Int](size int) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uint64) {
	v1, v2, vo := getBigSlice[T](size, size)
	return unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), uint64(size)
}


func BenchmarkInt64Adds(b *testing.B) {
    lengths := []int{6, 14, 52, 168, 552}
    for _, length := range lengths {
        b.Run(fmt.Sprintf("Naive_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int64](length)
            for i := 0; i < b.N; i++ {
                common.AddInt64(p1, p2, po, l)
            }
        })

        b.Run(fmt.Sprintf("AsmScalar_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int64](length)
            for i := 0; i < b.N; i++ {
                amd64.AddInt64Scalar(p1, p2, po, l)
            }
        })

        b.Run(fmt.Sprintf("SSE2_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int64](length)
            for i := 0; i < b.N; i++ {
                amd64.AddInt64SSE2(p1, p2, po, l)
            }
        })

        b.Run(fmt.Sprintf("AVX2_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int64](length)
            for i := 0; i < b.N; i++ {
                amd64.AddInt64AVX2(p1, p2, po, l)
            }
        })
    }
}



func BenchmarkInt32Adds(b *testing.B) {
    lengths := []int{6, 14, 52, 168, 552}
    for _, length := range lengths {
        b.Run(fmt.Sprintf("Naive_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int32](length)
            for i := 0; i < b.N; i++ {
                common.AddInt32(p1, p2, po, l)
            }
        })

		  b.Run(fmt.Sprintf("SSE2_%d", length), func(b *testing.B) {
			  p1, p2, po, l := getBenchmarkParams[int32](length)
			  for i := 0; i < b.N; i++ {
				  amd64.AddInt32SSE2(p1, p2, po, l)
			  }
		  })

        b.Run(fmt.Sprintf("AVX2_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int32](length)
            for i := 0; i < b.N; i++ {
                amd64.AddInt32AVX2(p1, p2, po, l)
            }
        })
    }
}


func BenchmarkInt16Adds(b *testing.B) {
    lengths := []int{6, 14, 52, 168, 552}
    for _, length := range lengths {
        b.Run(fmt.Sprintf("Naive_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int16](length)
            for i := 0; i < b.N; i++ {
                common.AddInt16(p1, p2, po, l)
            }
        })

		  b.Run(fmt.Sprintf("SSE2_%d", length), func(b *testing.B) {
			  p1, p2, po, l := getBenchmarkParams[int16](length)
			  for i := 0; i < b.N; i++ {
				  amd64.AddInt16SSE2(p1, p2, po, l)
			  }
		  })

        b.Run(fmt.Sprintf("AVX2_%d", length), func(b *testing.B) {
            p1, p2, po, l := getBenchmarkParams[int16](length)
            for i := 0; i < b.N; i++ {
                amd64.AddInt16AVX2(p1, p2, po, l)
            }
        })
    }
}
