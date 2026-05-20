package slice

import (
	"reflect"
	"testing"
)

func TestChunkSlice(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		size int
		want [][]int
	}{
		{
			name: "длина кратна размеру",
			src:  []int{1, 2, 3, 4, 5, 6},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "последний чанк короче",
			src:  []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "размер больше длины слайса",
			src:  []int{1, 2, 3},
			size: 10,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "размер равен длине слайса",
			src:  []int{1, 2, 3},
			size: 3,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "размер 1",
			src:  []int{1, 2, 3},
			size: 1,
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "size равен нулю",
			src:  []int{1, 2, 3},
			size: 0,
			want: nil,
		},
		{
			name: "size отрицательный",
			src:  []int{1, 2, 3},
			size: -2,
			want: nil,
		},
		{
			name: "пустой слайс",
			src:  []int{},
			size: 2,
			want: nil,
		},
		{
			name: "nil слайс",
			src:  nil,
			size: 2,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChunkSlice(tt.src, tt.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkSlice(%v, %d) = %v, want %v", tt.src, tt.size, got, tt.want)
			}
		})
	}
}
