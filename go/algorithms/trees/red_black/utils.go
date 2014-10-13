package red_black

import (
	"fmt"
)

func Compare(t1, t2 *RBTree) bool {
	c1, c2 := DepthWalker(t1), DepthWalker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func RBCheck(t *RBTree) error {
	if t.Red {
		return fmt.Errorf("The root must be black")
	}
	if BlackHeight(t) < 0 {
		return fmt.Errorf("Every path from a given node to any of its descendant leaves should contain the same number of black nodes")
	}
	ch := BreadthWalker(t)
	for {
		node, ok := <-ch
		if !ok {
			return nil
		}
		if node.Red {
			red_left := node.Left != nil && node.Left.Red
			red_right := node.Right != nil && node.Right.Red
			if red_left || red_right {
				return fmt.Errorf("Every red node must have two black child nodes. Failed at: %d", node.Value)
			}
		}
	}
	return nil
}

func BlackHeight(t *RBTree) int {
	if t == nil {
		return 0
	}
	var height int
	if !t.Red {
		height++
	}
	left := BlackHeight(t.Left)
	right := BlackHeight(t.Right)
	if left < 0 || right < 0 || left != right {
		return -1
	}
	if left > right {
		return height + left
	}
	return height + right
}
