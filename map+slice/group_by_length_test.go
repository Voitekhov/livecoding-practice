package mapslice

import (
	"reflect"
	"testing"
)

func TestGroupByLength(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[int][]string
	}{
		{
			name:  "разные длины",
			words: []string{"a", "bb", "cc", "ddd", "ee"},
			want: map[int][]string{
				1: {"a"},
				2: {"bb", "cc", "ee"},
				3: {"ddd"},
			},
		},
		{
			name:  "все одной длины",
			words: []string{"go", "rs", "py"},
			want: map[int][]string{
				2: {"go", "rs", "py"},
			},
		},
		{
			name:  "одно слово",
			words: []string{"hello"},
			want: map[int][]string{
				5: {"hello"},
			},
		},
		{
			name:  "пустая строка имеет длину 0",
			words: []string{"", "a", ""},
			want: map[int][]string{
				0: {"", ""},
				1: {"a"},
			},
		},
		{
			name:  "порядок внутри группы сохраняется",
			words: []string{"x", "yy", "z", "ww"},
			want: map[int][]string{
				1: {"x", "z"},
				2: {"yy", "ww"},
			},
		},
		{
			name:  "пустой слайс",
			words: []string{},
			want:  map[int][]string{},
		},
		{
			name:  "nil слайс",
			words: nil,
			want:  map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByLength(tt.words)
			if got == nil {
				t.Fatalf("GroupByLength(%v) вернул nil, ожидался непустой map", tt.words)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByLength(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}
