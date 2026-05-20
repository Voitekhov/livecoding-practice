package slice

import "testing"

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target int
		want   int
	}{
		{
			name:   "элемент в середине",
			src:    []int{10, 20, 30, 40},
			target: 30,
			want:   2,
		},
		{
			name:   "первый элемент",
			src:    []int{5, 6, 7},
			target: 5,
			want:   0,
		},
		{
			name:   "последний элемент",
			src:    []int{5, 6, 7},
			target: 7,
			want:   2,
		},
		{
			name:   "элемент не найден",
			src:    []int{1, 2, 3},
			target: 99,
			want:   -1,
		},
		{
			name:   "первое вхождение при дубликатах",
			src:    []int{4, 8, 8, 8, 4},
			target: 8,
			want:   1,
		},
		{
			name:   "пустой слайс",
			src:    []int{},
			target: 1,
			want:   -1,
		},
		{
			name:   "nil слайс",
			src:    nil,
			target: 1,
			want:   -1,
		},
		{
			name:   "отрицательное значение",
			src:    []int{0, -3, 7},
			target: -3,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindIndex(tt.src, tt.target)
			if got != tt.want {
				t.Errorf("FindIndex(%v, %d) = %d, want %d", tt.src, tt.target, got, tt.want)
			}
		})
	}
}
