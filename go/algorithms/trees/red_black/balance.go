package red_black

func Balance(t *RBTree) {
	parnt := t.Parent
	grand := t.Grandparent()
	uncle := t.Uncle()

	if uncle == nil {
		return
	}

	if uncle.Red {
		parnt.Red = false
		uncle.Red = false
		grand.Red = true
		Balance(grand)
		return
	} else {
		switch {
		case parnt == grand.Left && t == parnt.Left:
			RotateRight(grand)
		case parnt == grand.Left && t == parnt.Right:
			RotateLeft(parnt)
			RotateRight(grand)
		case parnt == grand.Right && t == parnt.Right:
			RotateLeft(grand)
		case parnt == grand.Right && t == parnt.Left:
			RotateRight(parnt)
			RotateLeft(grand)
		}
		grand.Red = !grand.Red
		t.Red = !t.Red
	}

}

func RotateRight(t *RBTree) {
	new_parent := t.Left

	parnt := t.Parent
	if parnt != nil {
		if t == parnt.Left {
			parnt.Left = new_parent
		} else {
			parnt.Right = new_parent
		}
	} else {
		new_parent.Parent = nil
	}

	t.Parent = new_parent
	t.Left = new_parent.Right
	new_parent.Right = t
}

func RotateLeft(t *RBTree) {
	new_parent := t.Right

	parnt := t.Parent
	if parnt != nil {
		if t == parnt.Left {
			parnt.Left = new_parent
		} else {
			parnt.Right = new_parent
		}
	} else {
		new_parent.Parent = nil
	}

	t.Parent = new_parent
	t.Right = new_parent.Left
	new_parent.Left = t
}

func (t *RBTree) Grandparent() *RBTree {
	if t.Parent == nil {
		return nil
	}
	return t.Parent.Parent
}

func (t *RBTree) Uncle() *RBTree {
	if grand := t.Grandparent(); grand != nil {
		if t.Parent == grand.Left {
			return grand.Right
		}
		return grand.Left
	}
	return nil
}

func (t *RBTree) FindRoot() *RBTree {
	if t.Parent == nil {
		return t
	}
	return t.Parent.FindRoot()
}
