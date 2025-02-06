package gonum

import (
	"reflect"
	"testing"
)

type Int interface {
    int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

func getBigSlice[T Int](size, offset int) ([]T, []T, []T) {
	v1 := make([]T, size)
	v2 := make([]T, size)
	vo := make([]T, size)
	for i := 0; i < size; i += 1 {
		v1[i] = T(i)
		v2[i] = T(i + offset)
		vo[i] = T(i) + T(i + offset)
	}
	return v1, v2, vo
}


func TestAddInt64(t *testing.T) {
	lv1, lv2, lsum := getBigSlice[int64](1024, 1024)

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
			name: "small that will force to go scalar",
			v1: []int64{2, 3},
			v2: []int64{4, 5},
			want: []int64{6, 8},
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


func TestAddInt32(t *testing.T) {
	lv1, lv2, lsum := getBigSlice[int32](1024, 1024)

	tests := []struct {
		name   string
		v1, v2 []int32
		want   []int32
	}{
		{
			name: "both empty slices",
			v1: []int32{},
			v2: []int32{},
			want: []int32{},
		},
		{
			name: "small that will force to go scalar",
			v1: []int32{2, 3},
			v2: []int32{4, 5},
			want: []int32{6, 8},
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
			got := make([]int32, len(tt.v1))
			AddInt32(tt.v1, tt.v2, got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddInt64(%v, %v) = %v; want %v",
					tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}

