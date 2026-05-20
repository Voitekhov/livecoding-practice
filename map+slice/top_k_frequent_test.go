package mapslice

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		k    int
		want []int
	}{
		{
			name: "классический пример",
			src:  []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "один элемент k=1",
			src:  []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "все элементы разные k=3",
			src:  []int{4, 5, 6, 7},
			k:    3,
			want: []int{4, 5, 6},
		},
		{
			name: "k больше количества уникальных",
			src:  []int{1, 2, 2, 3},
			k:    10,
			want: []int{2, 1, 3},
		},
		{
			name: "k = 0",
			src:  []int{1, 2, 3},
			k:    0,
			want: []int{},
		},
		{
			name: "пустой слайс",
			src:  []int{},
			k:    3,
			want: []int{},
		},
		{
			name: "nil слайс",
			src:  nil,
			k:    2,
			want: []int{},
		},
		{
			name: "отрицательные числа",
			src:  []int{-1, -1, -2, -2, -2, 3},
			k:    1,
			want: []int{-2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.src, tt.k)

			gotCopy := append([]int(nil), got...)
			wantCopy := append([]int(nil), tt.want...)
			sort.Ints(gotCopy)
			sort.Ints(wantCopy)

			if !reflect.DeepEqual(gotCopy, wantCopy) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v (порядок не важен)", tt.src, tt.k, got, tt.want)
			}
		})
	}
}
