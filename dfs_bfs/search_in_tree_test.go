package dfsbfs

import "testing"

func TestSearch(t *testing.T) {
	// tree — основное дерево из задачи.
	//
	//          1
	//        / | \
	//       2  3  4
	//      /|     |
	//     5 6     7
	//             |
	//             8
	//
	// Глубина 4, общее количество узлов 8.
	tree := &MultiNode{
		Val: 1,
		Children: []*MultiNode{
			{
				Val: 2,
				Children: []*MultiNode{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 3},
			{
				Val: 4,
				Children: []*MultiNode{
					{
						Val:      7,
						Children: []*MultiNode{{Val: 8}},
					},
				},
			},
		},
	}

	// noGrandchildren — дерево из корня и двух листьев, у одного из которых
	// Children == nil, а у другого Children == []*MultiNode{}. Оба случая
	// должны корректно считаться «листом без детей».
	//
	//        1
	//       / \
	//      2   3
	noGrandchildren := &MultiNode{
		Val: 1,
		Children: []*MultiNode{
			{Val: 2, Children: nil},
			{Val: 3, Children: []*MultiNode{}},
		},
	}

	tests := []struct {
		name   string
		root   *MultiNode
		target int
		want   bool
	}{
		{name: "корень", root: tree, target: 1, want: true},
		{name: "лист среднего уровня", root: tree, target: 6, want: true},
		{name: "самый глубокий лист", root: tree, target: 8, want: true},
		{name: "внутренний узел", root: tree, target: 4, want: true},
		{name: "значение отсутствует", root: tree, target: 9, want: false},
		{name: "nil дерево", root: nil, target: 1, want: false},
		{name: "один узел совпадает", root: &MultiNode{Val: 42}, target: 42, want: true},
		{name: "один узел не совпадает", root: &MultiNode{Val: 42}, target: 7, want: false},
		{name: "узел без детей (Children == nil или пуст)", root: noGrandchildren, target: 3, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.root, tt.target)
			if got != tt.want {
				t.Errorf("search(_, %d) = %t, want %t", tt.target, got, tt.want)
			}
		})
	}
}
