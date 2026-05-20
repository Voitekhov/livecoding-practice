package mapslice

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{
			name:  "несколько повторений",
			words: []string{"go", "rust", "go", "go", "rust", "c"},
			want:  map[string]int{"go": 3, "rust": 2, "c": 1},
		},
		{
			name:  "все слова уникальны",
			words: []string{"a", "b", "c"},
			want:  map[string]int{"a": 1, "b": 1, "c": 1},
		},
		{
			name:  "одно слово",
			words: []string{"hello"},
			want:  map[string]int{"hello": 1},
		},
		{
			name:  "регистр имеет значение",
			words: []string{"Go", "go", "GO"},
			want:  map[string]int{"Go": 1, "go": 1, "GO": 1},
		},
		{
			name:  "пустые строки считаются отдельным словом",
			words: []string{"", "", "x"},
			want:  map[string]int{"": 2, "x": 1},
		},
		{
			name:  "пустой слайс",
			words: []string{},
			want:  map[string]int{},
		},
		{
			name:  "nil слайс",
			words: nil,
			want:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordFrequency(tt.words)
			if got == nil {
				t.Fatalf("WordFrequency(%v) вернул nil, ожидался непустой map", tt.words)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordFrequency(%v) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}
