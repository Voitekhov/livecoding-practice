package dfsbfs

import "testing"

func TestMaxDepth(t *testing.T) {
	// skewedLeft — пример из задачи: левая ветвь длиннее правой.
	//
	//        1
	//       / \
	//      2   3
	//     /
	//    4
	//   /
	//  5
	//
	// Самый длинный путь: 1 -> 2 -> 4 -> 5, глубина = 4.
	skewedLeft := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:  4,
				Left: &Node{Val: 5},
			},
		},
		Right: &Node{Val: 3},
	}

	// balanced — полное дерево из трёх узлов.
	//
	//      1
	//     / \
	//    2   3
	//
	// Любой путь от корня до листа имеет длину 2.
	balanced := &Node{
		Val:   1,
		Left:  &Node{Val: 2},
		Right: &Node{Val: 3},
	}

	// rightChainOnly — вырожденное дерево-цепочка только по правому ребру.
	//
	//    1
	//     \
	//      2
	//       \
	//        3
	//         \
	//          4
	//
	// Глубина равна количеству узлов в цепочке = 4.
	rightChainOnly := &Node{
		Val: 1,
		Right: &Node{
			Val: 2,
			Right: &Node{
				Val:   3,
				Right: &Node{Val: 4},
			},
		},
	}

	tests := []struct {
		name string
		root *Node
		want int
	}{
		{name: "nil дерево", root: nil, want: 0},
		{name: "один узел", root: &Node{Val: 42}, want: 1},
		{name: "сбалансированное дерево из 3 узлов", root: balanced, want: 2},
		{name: "цепочка только вправо", root: rightChainOnly, want: 4},
		{name: "пример из задачи (длиннее левая ветка)", root: skewedLeft, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			if got != tt.want {
				t.Errorf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
