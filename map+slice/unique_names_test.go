package mapslice

import (
	"reflect"
	"testing"
)

func TestUniqueNames(t *testing.T) {
	tests := []struct {
		name  string
		names []string
		want  []string
	}{
		{
			name:  "обычный случай с повторами",
			names: []string{"Alice", "Bob", "Alice", "Charlie", "Bob"},
			want:  []string{"Alice", "Bob", "Alice(1)", "Charlie", "Bob(1)"},
		},
		{
			name:  "все имена уникальны — без изменений",
			names: []string{"Alice", "Bob", "Charlie"},
			want:  []string{"Alice", "Bob", "Charlie"},
		},
		{
			name:  "все имена одинаковы",
			names: []string{"Alice", "Alice", "Alice", "Alice"},
			want:  []string{"Alice", "Alice(1)", "Alice(2)", "Alice(3)"},
		},
		{
			name:  "одно имя",
			names: []string{"Alice"},
			want:  []string{"Alice"},
		},
		{
			name:  "суффикс уже занят на входе — пропускаем k=1",
			names: []string{"a", "a(1)", "a"},
			want:  []string{"a", "a(1)", "a(2)"},
		},
		{
			name:  "несколько уже занятых суффиксов подряд",
			names: []string{"a", "a(1)", "a(2)", "a", "a"},
			want:  []string{"a", "a(1)", "a(2)", "a(3)", "a(4)"},
		},
		{
			name:  "регистр имеет значение",
			names: []string{"Alice", "alice", "Alice"},
			want:  []string{"Alice", "alice", "Alice(1)"},
		},
		{
			name:  "пустая строка — валидное имя",
			names: []string{"", "", "Alice"},
			want:  []string{"", "(1)", "Alice"},
		},
		{
			name:  "повторы перемежаются разными именами",
			names: []string{"x", "y", "x", "y", "x"},
			want:  []string{"x", "y", "x(1)", "y(1)", "x(2)"},
		},
		{
			name:  "пустой слайс",
			names: []string{},
			want:  []string{},
		},
		{
			name:  "nil слайс",
			names: nil,
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueNames(tt.names)

			if len(got) != len(tt.names) && !(len(got) == 0 && len(tt.names) == 0) {
				t.Fatalf("UniqueNames(%v): длина результата %d != длине входа %d", tt.names, len(got), len(tt.names))
			}

			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueNames(%v) = %v, want %v", tt.names, got, tt.want)
			}
		})
	}
}
