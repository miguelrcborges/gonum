package gonum

import (
	"reflect"
	"testing"
)

func getBigInt64Slice(size, offset uint64) ([]int64, []int64, []int64) {
	v1 := make([]int64, size)
	v2 := make([]int64, size)
	vo := make([]int64, size)
	for i := uint64(0); i < size; i += 1 {
		v1[i] = int64(i)
		v2[i] = int64(i + offset)
		vo[i] = int64(i) + int64(i + offset)
	}
	return v1, v2, vo
}


func TestAddInt64(t *testing.T) {
	lv1, lv2, lsum := getBigInt64Slice(1024, 1024)

	tests := []struct {
		name   string
		v1, v2 []int64
		want   []int64
	}{
		{
			name: "both empty slices",
			v1: []int64{},
			v2: []int64{},
			want: []int64{},
		},
		{
			name: "large 1",
			v1: lv1,
			v2: lv2,
			want: lsum,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int64, len(tt.v1))
			AddInt64(tt.v1, tt.v2, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddInt64(%v, %v) = %v; want %v",
					tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}
