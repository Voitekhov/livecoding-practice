package slice

import (
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "обычный слайс",
			src:  []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "один элемент",
			src:  []int{42},
			want: []int{42},
		},
		{
			name: "два элемента",
			src:  []int{1, 2},
			want: []int{2, 1},
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
		{
			name: "повторяющиеся элементы",
			src:  []int{7, 7, 1, 7},
			want: []int{7, 1, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseSlice(tt.src)
			if len(got) != len(tt.want) {
				t.Fatalf("ReverseSlice(%v) = %v, want %v", tt.src, got, tt.want)
			}
			if len(got) != 0 && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSlice(%v) = %v, want %v", tt.src, got, tt.want)
			}
		})
	}
}

func TestReverseSliceDoesNotMutateSource(t *testing.T) {
	src := []int{1, 2, 3}
	srcCopy := []int{1, 2, 3}

	_ = ReverseSlice(src)

	if !reflect.DeepEqual(src, srcCopy) {
		t.Errorf("ReverseSlice изменил исходный слайс: было %v, стало %v", srcCopy, src)
	}
}
