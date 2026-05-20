package slice

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		src    []int
		target int
		// wantAny — список допустимых индексов. Для уникальных значений
		// это один индекс; для дубликатов — любой из них; для «не найдено»
		// это []int{-1}.
		wantAny []int
	}{
		{
			name:    "target в середине",
			src:     []int{1, 3, 5, 7, 9, 11},
			target:  7,
			wantAny: []int{3},
		},
		{
			name:    "target — первый элемент",
			src:     []int{1, 3, 5, 7, 9},
			target:  1,
			wantAny: []int{0},
		},
		{
			name:    "target — последний элемент",
			src:     []int{1, 3, 5, 7, 9},
			target:  9,
			wantAny: []int{4},
		},
		{
			name:    "target меньше всех",
			src:     []int{2, 4, 6, 8},
			target:  1,
			wantAny: []int{-1},
		},
		{
			name:    "target больше всех",
			src:     []int{2, 4, 6, 8},
			target:  100,
			wantAny: []int{-1},
		},
		{
			name:    "target между элементами",
			src:     []int{1, 3, 5, 7, 9},
			target:  4,
			wantAny: []int{-1},
		},
		{
			name:    "один элемент — найден",
			src:     []int{42},
			target:  42,
			wantAny: []int{0},
		},
		{
			name:    "один элемент — не найден",
			src:     []int{42},
			target:  7,
			wantAny: []int{-1},
		},
		{
			name:    "два элемента — левый",
			src:     []int{1, 2},
			target:  1,
			wantAny: []int{0},
		},
		{
			name:    "два элемента — правый",
			src:     []int{1, 2},
			target:  2,
			wantAny: []int{1},
		},
		{
			name:    "два элемента — не найден",
			src:     []int{1, 2},
			target:  3,
			wantAny: []int{-1},
		},
		{
			name:    "пустой слайс",
			src:     []int{},
			target:  1,
			wantAny: []int{-1},
		},
		{
			name:    "nil слайс",
			src:     nil,
			target:  1,
			wantAny: []int{-1},
		},
		{
			name:    "дубликаты — допустим любой индекс совпадения",
			src:     []int{1, 2, 2, 2, 3},
			target:  2,
			wantAny: []int{1, 2, 3},
		},
		{
			name:    "отрицательные числа",
			src:     []int{-10, -5, -3, 0, 4, 8},
			target:  -3,
			wantAny: []int{2},
		},
		{
			name:    "чётная длина — элемент в середине",
			src:     []int{1, 2, 3, 4, 5, 6},
			target:  3,
			wantAny: []int{2},
		},
		{
			name:    "большой слайс",
			src:     []int{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80},
			target:  55,
			wantAny: []int{11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.src, tt.target)

			ok := false
			for _, w := range tt.wantAny {
				if got == w {
					ok = true
					break
				}
			}
			if !ok {
				t.Errorf("BinarySearch(%v, %d) = %d, want one of %v", tt.src, tt.target, got, tt.wantAny)
			}
		})
	}
}
