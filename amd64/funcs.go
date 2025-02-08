package amd64

import "unsafe"

func CpuId(flag uint32) (eax, ebx, ecx, edx uint32)

func AddInt64Scalar(v1, v2, vo unsafe.Pointer, l uint64)
func AddInt64SSE2(v1, v2, vo unsafe.Pointer, l uint64)
func AddInt64AVX2(v1, v2, vo unsafe.Pointer, l uint64)

func AddInt32SSE2(v1, v2, vo unsafe.Pointer, l uint64)
func AddInt32AVX2(v1, v2, vo unsafe.Pointer, l uint64)

func AddInt16SSE2(v1, v2, vo unsafe.Pointer, l uint64)
func AddInt16AVX2(v1, v2, vo unsafe.Pointer, l uint64)
