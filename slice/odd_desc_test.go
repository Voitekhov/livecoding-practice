package slice

import (
	"reflect"
	"testing"
)

func TestOddDesc(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "смесь чётных и нечётных",
			src:  []int{4, 1, 7, 2, 3, 8, 5},
			want: []int{7, 5, 3, 1},
		},
		{
			name: "все нечётные",
			src:  []int{3, 1, 9, 5, 7},
			want: []int{9, 7, 5, 3, 1},
		},
		{
			name: "все чётные",
			src:  []int{2, 4, 6, 8},
			want: []int{},
		},
		{
			name: "ноль считается чётным",
			src:  []int{0, 1, 0, 3},
			want: []int{3, 1},
		},
		{
			name: "один нечётный элемент",
			src:  []int{7},
			want: []int{7},
		},
		{
			name: "один чётный элемент",
			src:  []int{8},
			want: []int{},
		},
		{
			name: "повторяющиеся нечётные",
			src:  []int{3, 3, 1, 1, 5},
			want: []int{5, 3, 3, 1, 1},
		},
		{
			name: "отрицательные числа",
			src:  []int{-3, -2, -7, -1, 0, 4},
			want: []int{-1, -3, -7},
		},
		{
			name: "уже отсортирован по убыванию",
			src:  []int{9, 7, 5, 3, 1},
			want: []int{9, 7, 5, 3, 1},
		},
		{
			name: "пустой слайс",
			src:  []int{},
			want: []int{},
		},
		{
			name: "nil слайс",
			src:  nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srcCopy := append([]int(nil), tt.src...)
			got := OddDesc(tt.src)

			if len(got) == 0 && len(tt.want) == 0 {
				// nil и пустой слайс считаем эквивалентными
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OddDesc(%v) = %v, want %v", tt.src, got, tt.want)
			}

			if len(tt.src) != len(srcCopy) || (len(tt.src) > 0 && !reflect.DeepEqual(tt.src, srcCopy)) {
				t.Errorf("OddDesc мутировал исходный слайс: было %v, стало %v", srcCopy, tt.src)
			}
		})
	}
}
