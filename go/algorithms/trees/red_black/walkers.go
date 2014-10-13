package red_black

// Walkers launch Walks in a new goroutine,
// and returns a read-only channel of nodes.

func DepthWalker(t *RBTree) <-chan *RBTree {
	ch := make(chan *RBTree)
	go func() {
		DepthWalk(ch, t)
		close(ch)
	}()
	return ch
}

func BreadthWalker(t *RBTree) <-chan *RBTree {
	ch := make(chan *RBTree)
	go func() {
		BreadthWalk(ch, t)
		close(ch)
	}()
	return ch
}

// Walks traverses a tree depth-first, or breadth-first,
// sending each node on a channel.

func DepthWalk(ch chan *RBTree, t *RBTree) {
	if t == nil {
		return
	}
	DepthWalk(ch, t.Left)
	ch <- t
	DepthWalk(ch, t.Right)
}

func BreadthWalk(ch chan *RBTree, ts ...*RBTree) {
	if len(ts) == 0 {
		return
	}
	var next []*RBTree
	for _, t := range ts {
		if t == nil {
			continue
		}
		ch <- t
		next = append(next, t.Left)
		next = append(next, t.Right)
	}
	BreadthWalk(ch, next...)
}
