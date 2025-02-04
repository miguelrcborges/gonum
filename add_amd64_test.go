package gonum

import (
	"testing"
	"unsafe"

	"github.com/miguelrcborges/gonum/common"
	"github.com/miguelrcborges/gonum/amd64"
)

func getInt64BenchmarkParams(size int64) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, int64) {
	v1, v2, vo := getBigInt64Slice(size, size)
	return unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&vo[0]), size
}

func BenchmarkNaiveAdd_6(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(6)
	for n := 0; n < b.N; n += 1 {
		common.AddInt64(p1, p2, po, l)
	}
}

func BenchmarkAVX2Add_6(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(6)
	for n := 0; n < b.N; n += 1 {
		amd64.AddInt64AVX2(p1, p2, po, l)
	}
}

func BenchmarkNaiveAdd_14(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(14)
	for n := 0; n < b.N; n += 1 {
		common.AddInt64(p1, p2, po, l)
	}
}

func BenchmarkAVX2Add_14(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(14)
	for n := 0; n < b.N; n += 1 {
		amd64.AddInt64AVX2(p1, p2, po, l)
	}
}

func BenchmarkNaiveAdd_52(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(52)
	for n := 0; n < b.N; n += 1 {
		common.AddInt64(p1, p2, po, l)
	}
}

func BenchmarkAVX2Add_52(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(52)
	for n := 0; n < b.N; n += 1 {
		amd64.AddInt64AVX2(p1, p2, po, l)
	}
}

func BenchmarkNaiveAdd_168(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(168)
	for n := 0; n < b.N; n += 1 {
		common.AddInt64(p1, p2, po, l)
	}
}

func BenchmarkAVX2Add_168(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(168)
	for n := 0; n < b.N; n += 1 {
		amd64.AddInt64AVX2(p1, p2, po, l)
	}
}

func BenchmarkNaiveAdd_552(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(552)
	for n := 0; n < b.N; n += 1 {
		common.AddInt64(p1, p2, po, l)
	}
}

func BenchmarkAVX2Add_552(b *testing.B) {
	p1, p2, po, l := getInt64BenchmarkParams(552)
	for n := 0; n < b.N; n += 1 {
		amd64.AddInt64AVX2(p1, p2, po, l)
	}
}
