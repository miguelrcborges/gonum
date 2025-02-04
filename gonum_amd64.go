package gonum

import (
	"github.com/miguelrcborges/gonum/amd64"
)


func init() {
	eax, ebx, ecx, edx := amd64.CpuId(7)
	_ = eax
	_ = ebx
	_ = ecx
	_ = edx
}
