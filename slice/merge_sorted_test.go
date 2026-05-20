package slice

import (
	"reflect"
	"testing"
)

func TestMergeSorted(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "слайсы одинаковой длины, чередуются",
			a:    []int{1, 3, 5},
			b:    []int{2, 4, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "слайсы разной длины",
			a:    []int{1, 2},
			b:    []int{3, 4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "все элементы a меньше всех элементов b",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "все элементы b меньше всех элементов a",
			a:    []int{10, 20},
			b:    []int{1, 2},
			want: []int{1, 2, 10, 20},
		},
		{
			name: "с одинаковыми элементами в a и b",
			a:    []int{1, 3, 3},
			b:    []int{2, 3, 4},
			want: []int{1, 2, 3, 3, 3, 4},
		},
		{
			name: "первый слайс пустой",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "второй слайс пустой",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "оба слайса пустые",
			a:    []int{},
			b:    []int{},
			want: []int{},
		},
		{
			name: "оба слайса nil",
			a:    nil,
			b:    nil,
			want: []int{},
		},
		{
			name: "один элемент в каждом",
			a:    []int{2},
			b:    []int{1},
			want: []int{1, 2},
		},
		{
			name: "с отрицательными числами",
			a:    []int{-5, -1, 3},
			b:    []int{-3, 0, 4},
			want: []int{-5, -3, -1, 0, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSorted(tt.a, tt.b)
			if len(got) != len(tt.want) {
				t.Fatalf("MergeSorted(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
			if len(got) != 0 && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSorted(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMergeSortedDoesNotMutateInputs(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	aCopy := []int{1, 3, 5}
	bCopy := []int{2, 4, 6}

	_ = MergeSorted(a, b)

	if !reflect.DeepEqual(a, aCopy) {
		t.Errorf("MergeSorted изменил слайс a: было %v, стало %v", aCopy, a)
	}
	if !reflect.DeepEqual(b, bCopy) {
		t.Errorf("MergeSorted изменил слайс b: было %v, стало %v", bCopy, b)
	}
}
