package amd64

import "unsafe"

func CpuId(flag uint32) (eax, ebx, ecx, edx uint32)

func AddInt64Scalar(v1, v2, vo unsafe.Pointer, l uint64)
func AddInt64AVX2(v1, v2, vo unsafe.Pointer, l uint64)
