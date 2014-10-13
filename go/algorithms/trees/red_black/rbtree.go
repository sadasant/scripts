package red_black

import (
	"fmt"
)

type RBTree struct {
	Value  int
	Red    bool
	Parent *RBTree
	Left   *RBTree
	Right  *RBTree
}

func New(t *RBTree) *RBTree {
	ch := BreadthWalker(t)
	for {
		node, ok := <-ch
		if !ok {
			return t
		}
		if node.Left != nil {
			node.Left.Parent = node
		}
		if node.Right != nil {
			node.Right.Parent = node
		}
	}
}

func (t *RBTree) Insert(v int) {
	if v < t.Value {
		if t.Left == nil {
			t.Left = &RBTree{v, true, t, nil, nil}
			Balance(t.Left)
			return
		}
		t.Left.Insert(v)
		return
	}
	if t.Right == nil {
		t.Right = &RBTree{v, true, t, nil, nil}
		Balance(t.Right)
		return
	}
	t.Right.Insert(v)
}

func (t *RBTree) String() (str string) {
	ch := BreadthWalker(t)
	for {
		node, ok := <-ch
		if !ok {
			return str[:len(str)-1]
		}
		// println(node.Value)
		if node.Left != nil {
			var red string
			if node.Left.Red {
				red = " Red"
			}
			str += fmt.Sprintf("%d -> %d%s\n", node.Value, node.Left.Value, red)
		}
		if node.Right != nil {
			var red string
			if node.Right.Red {
				red = " Red"
			}
			str += fmt.Sprintf("%d -> %d%s\n", node.Value, node.Right.Value, red)
		}
	}
	return
}
