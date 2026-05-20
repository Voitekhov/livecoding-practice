package dfsbfs

import "testing"

func TestCountLevels(t *testing.T) {
	// example — пример из задачи: левая ветвь спускается до 4-го уровня.
	//
	//        1          <- уровень 1
	//       / \
	//      2   3        <- уровень 2
	//     /
	//    4              <- уровень 3
	//   /
	//  5                <- уровень 4
	//
	// countLevels(example) == 4.
	example := &Node{
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

	// balanced — полное бинарное дерево глубины 3.
	//
	//          1          <- уровень 1
	//        /   \
	//       2     3       <- уровень 2
	//      / \   / \
	//     4   5 6   7     <- уровень 3
	//
	// countLevels(balanced) == 3.
	balanced := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}

	// rightChain — цепочка только вправо: каждый узел на своём уровне.
	//
	//    1                <- уровень 1
	//     \
	//      2              <- уровень 2
	//       \
	//        3            <- уровень 3
	//
	// countLevels(rightChain) == 3.
	rightChain := &Node{
		Val: 1,
		Right: &Node{
			Val:   2,
			Right: &Node{Val: 3},
		},
	}

	tests := []struct {
		name string
		root *Node
		want int
	}{
		{name: "nil дерево", root: nil, want: 0},
		{name: "один узел", root: &Node{Val: 1}, want: 1},
		{name: "правая цепочка из трёх", root: rightChain, want: 3},
		{name: "пример из задачи", root: example, want: 4},
		{name: "полное дерево глубины 3", root: balanced, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countLevels(tt.root)
			if got != tt.want {
				t.Errorf("countLevels() = %d, want %d", got, tt.want)
			}
		})
	}
}
