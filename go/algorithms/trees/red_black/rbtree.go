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

// This New function makes sure every node has it's parent
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
			InsertBalance(t.Left)
			return
		}
		t.Left.Insert(v)
		return
	}
	if t.Right == nil {
		t.Right = &RBTree{v, true, t, nil, nil}
		InsertBalance(t.Right)
		return
	}
	t.Right.Insert(v)
}

func (t *RBTree) Delete() {
	parnt := t.Parent
	if parnt == nil {
		return
	}
	// If it's a red leaf
	if t.Red && t.Left == nil && t.Right == nil {
		if t == parnt.Left {
			parnt.Left = nil
		} else {
			parnt.Right = nil
		}
		return
	}
	// Has one leaf and it's red or it leaf is red
	if t.Left != nil && (t.Red || t.Left.Red) {
		t.Left.Red = false
		if t == parnt.Left {
			parnt.Left = t.Left
		} else {
			parnt.Right = t.Left
		}
		return
	}
	if t.Right != nil && (t.Red || t.Right.Red) {
		t.Right.Red = false
		if t == parnt.Left {
			parnt.Left = t.Right
		} else {
			parnt.Right = t.Right
		}
		return
	}
	sibli := t.Sibling()
	if sibli == nil {
		return
	}
	var r_left, r_right bool
	if sibli.Left != nil {
		r_left = sibli.Left.Red
	}
	if sibli.Right != nil {
		r_right = sibli.Right.Red
	}
	// Double blacks and has at least one red child
	if !sibli.Red {
		if r_left || r_right {
			switch {
			case sibli == parnt.Left && r_left:
				parnt.Right = nil
				parnt.Left.Red = false
				RotateRight(parnt)
			case sibli == parnt.Left && r_right:
				parnt.Right = nil
				sibli.Right.Red = false
				RotateLeft(sibli)
				RotateRight(parnt)
			case sibli == parnt.Right && r_right:
				parnt.Left = nil
				sibli.Right.Red = false
				RotateLeft(parnt)
			case sibli == parnt.Right && r_left:
				parnt.Left = nil
				sibli.Left.Red = false
				RotateRight(sibli)
				RotateLeft(parnt)
			}
			return
		}
		// Everyone is black
		if sibli == parnt.Left {
			parnt.Right = nil
		}
		if sibli == parnt.Right {
			parnt.Left = nil
		}
		sibli.Red = true
		parnt.Red = false
		return
	}
	// Sibling is red
	sibli.Red = false
	if sibli == parnt.Right {
		parnt.Left = nil
		if sibli.Left != nil {
			sibli.Left.Red = true
		}
		RotateLeft(parnt)
	} else {
		parnt.Right = nil
		if sibli.Right != nil {
			sibli.Right.Red = true
		}
		RotateRight(parnt)
	}

}

// Prints something like graphviz's structures
func (t *RBTree) String() (str string) {
	ch := BreadthWalker(t)
	for {
		node, ok := <-ch
		if !ok {
			if str == "" {
				return
			}
			return str[:len(str)-1]
		}
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
