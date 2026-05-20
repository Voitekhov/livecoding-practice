package mapslice

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		wantOK bool
	}{
		{
			name:   "классический пример",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			wantOK: true,
		},
		{
			name:   "пара в середине",
			nums:   []int{3, 2, 4},
			target: 6,
			wantOK: true,
		},
		{
			name:   "два одинаковых элемента",
			nums:   []int{3, 3},
			target: 6,
			wantOK: true,
		},
		{
			name:   "пары нет",
			nums:   []int{1, 2, 3, 4},
			target: 100,
			wantOK: false,
		},
		{
			name:   "пустой слайс",
			nums:   []int{},
			target: 0,
			wantOK: false,
		},
		{
			name:   "nil слайс",
			nums:   nil,
			target: 0,
			wantOK: false,
		},
		{
			name:   "один элемент — пары нет",
			nums:   []int{5},
			target: 5,
			wantOK: false,
		},
		{
			name:   "отрицательные числа",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			wantOK: true,
		},
		{
			name:   "нельзя использовать один индекс дважды",
			nums:   []int{3, 1, 4},
			target: 6,
			wantOK: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)

			if !tt.wantOK {
				if got != nil {
					t.Errorf("TwoSum(%v, %d) = %v, ожидался nil", tt.nums, tt.target, got)
				}
				return
			}

			if len(got) != 2 {
				t.Fatalf("TwoSum(%v, %d) = %v, ожидался слайс длины 2", tt.nums, tt.target, got)
			}
			i, j := got[0], got[1]
			if i == j {
				t.Errorf("TwoSum(%v, %d) = %v, индексы должны быть различны", tt.nums, tt.target, got)
			}
			if i < 0 || j < 0 || i >= len(tt.nums) || j >= len(tt.nums) {
				t.Errorf("TwoSum(%v, %d) = %v, индексы вне диапазона", tt.nums, tt.target, got)
			}
			if tt.nums[i]+tt.nums[j] != tt.target {
				t.Errorf("TwoSum(%v, %d) = %v, сумма %d != %d", tt.nums, tt.target, got, tt.nums[i]+tt.nums[j], tt.target)
			}
		})
	}
}
