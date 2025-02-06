package common

import "unsafe"

func SubInt64(v1, v2, vo unsafe.Pointer, l uint64) {
	for l > 0 {
		*(*uint64)(vo)= *(*uint64)(v1) - *(*uint64)(v2)
		vo = unsafe.Add(vo, 8)
		v1 = unsafe.Add(v1, 8)
		v2 = unsafe.Add(v2, 8)
		l -= 1
	}
}


func SubInt32(v1, v2, vo unsafe.Pointer, l uint64) {
	for l > 0 {
		*(*uint32)(vo)= *(*uint32)(v1) - *(*uint32)(v2)
		vo = unsafe.Add(vo, 4)
		v1 = unsafe.Add(v1, 4)
		v2 = unsafe.Add(v2, 4)
		l -= 1
	}
}


func SubInt16(v1, v2, vo unsafe.Pointer, l uint64) {
	for l > 0 {
		*(*uint16)(vo)= *(*uint16)(v1) - *(*uint16)(v2)
		vo = unsafe.Add(vo, 2)
		v1 = unsafe.Add(v1, 2)
		v2 = unsafe.Add(v2, 2)
		l -= 1
	}
}


func SubInt8(v1, v2, vo unsafe.Pointer, l uint64) {
	for l > 0 {
		*(*uint8)(vo)= *(*uint8)(v1) - *(*uint8)(v2)
		vo = unsafe.Add(vo, 8)
		v1 = unsafe.Add(v1, 8)
		v2 = unsafe.Add(v2, 8)
		l -= 1
	}
}
