package mapslice

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeInverted(m map[int][]string) map[int][]string {
	out := make(map[int][]string, len(m))
	for k, v := range m {
		vc := append([]string(nil), v...)
		sort.Strings(vc)
		out[k] = vc
	}
	return out
}

func TestInvertMap(t *testing.T) {
	tests := []struct {
		name string
		in   map[string]int
		want map[int][]string
	}{
		{
			name: "коллизии по значению",
			in:   map[string]int{"a": 1, "b": 2, "c": 1},
			want: map[int][]string{1: {"a", "c"}, 2: {"b"}},
		},
		{
			name: "все значения уникальны",
			in:   map[string]int{"x": 10, "y": 20, "z": 30},
			want: map[int][]string{10: {"x"}, 20: {"y"}, 30: {"z"}},
		},
		{
			name: "все значения одинаковы",
			in:   map[string]int{"a": 5, "b": 5, "c": 5},
			want: map[int][]string{5: {"a", "b", "c"}},
		},
		{
			name: "одна запись",
			in:   map[string]int{"only": 42},
			want: map[int][]string{42: {"only"}},
		},
		{
			name: "отрицательные значения",
			in:   map[string]int{"a": -1, "b": -1, "c": 0},
			want: map[int][]string{-1: {"a", "b"}, 0: {"c"}},
		},
		{
			name: "пустой map",
			in:   map[string]int{},
			want: map[int][]string{},
		},
		{
			name: "nil map",
			in:   nil,
			want: map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InvertMap(tt.in)
			if got == nil {
				t.Fatalf("InvertMap(%v) вернул nil, ожидался непустой map", tt.in)
			}
			gotNorm := normalizeInverted(got)
			wantNorm := normalizeInverted(tt.want)
			if !reflect.DeepEqual(gotNorm, wantNorm) {
				t.Errorf("InvertMap(%v) = %v, want %v (порядок строк не важен)", tt.in, got, tt.want)
			}
		})
	}
}
