package dfsbfs

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type MultiNode struct {
	Val      int          // значение узла
	Children []*MultiNode // список потомков (может быть пустым)
}
