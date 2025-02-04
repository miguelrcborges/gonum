package common

import "unsafe"

func AddInt64(v1, v2, vo unsafe.Pointer, l int64) {
	for l > 0 {
		*(*uint64)(vo)= *(*uint64)(v1) + *(*uint64)(v2)
		vo = unsafe.Add(vo, 8)
		v1 = unsafe.Add(v1, 8)
		v2 = unsafe.Add(v2, 8)
		l -= 1
	}
}
