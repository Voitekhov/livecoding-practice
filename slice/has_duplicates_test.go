package slice

import "testing"

func TestHasDuplicates(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want bool
	}{
		{
			name: "есть дубликаты",
			src:  []int{1, 2, 3, 2},
			want: true,
		},
		{
			name: "все элементы уникальны",
			src:  []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "два одинаковых элемента подряд",
			src:  []int{7, 7},
			want: true,
		},
		{
			name: "два разных элемента",
			src:  []int{7, 8},
			want: false,
		},
		{
			name: "один элемент",
			src:  []int{1},
			want: false,
		},
		{
			name: "пустой слайс",
			src:  []int{},
			want: false,
		},
		{
			name: "nil слайс",
			src:  nil,
			want: false,
		},
		{
			name: "дубликаты на краях",
			src:  []int{5, 1, 2, 3, 5},
			want: true,
		},
		{
			name: "все элементы одинаковы",
			src:  []int{9, 9, 9, 9},
			want: true,
		},
		{
			name: "дубликат среди отрицательных",
			src:  []int{-1, -2, -1},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasDuplicates(tt.src)
			if got != tt.want {
				t.Errorf("HasDuplicates(%v) = %t, want %t", tt.src, got, tt.want)
			}
		})
	}
}
