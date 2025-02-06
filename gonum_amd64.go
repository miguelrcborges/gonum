package gonum

import (
	"github.com/miguelrcborges/gonum/amd64"
)


var SSE2 bool
var SSE4_2 bool
var AVX bool
var AVX2 bool
var AVX512 bool


func init() {
	eax, ebx, ecx, edx := amd64.CpuId(1)
	_ = eax
	_ = ebx
	_ = ecx
	_ = edx
	SSE2 = true
	SSE4_2 = (ecx & (1 << 20)) != 0

	UnsafeAddInt64 = amd64.AddInt64SSE2
	UnsafeAddInt32 = amd64.AddInt32SSE2
}
