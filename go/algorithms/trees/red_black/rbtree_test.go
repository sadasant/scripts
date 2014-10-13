package red_black

import (
	"fmt"
	"testing"
)

func TestRBCheck(t *testing.T) {
	tree := New(&RBTree{
		11, false, nil,
		&RBTree{
			2, true, nil,
			&RBTree{1, false, nil, nil, nil},
			&RBTree{
				7, false, nil,
				&RBTree{5, true, nil, nil, nil},
				&RBTree{8, true, nil, nil, nil},
			},
		},
		&RBTree{
			14, false, nil,
			&RBTree{15, true, nil, nil, nil},
			nil,
		},
	})

	if err := RBCheck(tree); err != nil {
		t.Errorf("Error: %s.\nTree:\n%s\n", err.Error(), tree.String())
	}

	tree.Insert(4)
	tree = tree.FindRoot()

	if err := RBCheck(tree); err != nil {
		t.Errorf("Error: %s.\nTree:\n%s\n", err.Error(), tree.String())
	}
}

func ExampleRedBlack() {
	tree := New(&RBTree{
		11, false, nil,
		&RBTree{
			2, true, nil,
			&RBTree{1, false, nil, nil, nil},
			&RBTree{
				7, false, nil,
				&RBTree{5, true, nil, nil, nil},
				&RBTree{8, true, nil, nil, nil},
			},
		},
		&RBTree{
			14, false, nil,
			&RBTree{15, true, nil, nil, nil},
			nil,
		},
	})

	fmt.Printf("\nOriginal:\n%s\n", tree.String())
	tree.Insert(4)
	tree = tree.FindRoot()
	fmt.Printf("After inserting 4:\n%s\n", tree.String())
	// Output:
	//Original:
	//11 -> 2 Red
	//11 -> 14
	//2 -> 1
	//2 -> 7
	//14 -> 15 Red
	//7 -> 5 Red
	//7 -> 8 Red
	//After inserting 4:
	//7 -> 2 Red
	//7 -> 11 Red
	//2 -> 1
	//2 -> 5
	//11 -> 8
	//11 -> 14
	//5 -> 4 Red
	//14 -> 15 Red
}

func TestDelete(t *testing.T) {

	// Either the node to be removed or one of it's childs is red
	tree := New(&RBTree{
		30, false, nil,
		&RBTree{
			20, false, nil,
			&RBTree{10, true, nil, nil, nil},
			nil,
		},
		&RBTree{40, false, nil, nil, nil},
	})
	expected := New(&RBTree{
		30, false, nil,
		&RBTree{10, false, nil, nil, nil},
		&RBTree{40, false, nil, nil, nil},
	})
	tree.Left.Delete()
	if !Compare(tree, expected) {
		t.Fatalf("Wanted:\n%s\nGot:\n%s\n", expected.String(), tree.String())
	}
	if err := RBCheck(tree); err != nil {
		t.Errorf("Error: %s.\nTree:\n%s\n", err.Error(), tree.String())
	}

	// Deleting node is black and it's childs are black, also it's sibling is black, but it has one child red
	tree = New(&RBTree{
		30, false, nil,
		&RBTree{20, false, nil, nil, nil},
		&RBTree{
			40, false, nil, nil,
			&RBTree{50, true, nil, nil, nil},
		},
	})
	expected = New(&RBTree{
		40, false, nil,
		&RBTree{30, false, nil, nil, nil},
		&RBTree{50, false, nil, nil, nil},
	})
	tree.Left.Delete()
	tree = tree.FindRoot()
	if !Compare(tree, expected) {
		t.Fatalf("Wanted:\n%s\nGot:\n%s\n", expected.String(), tree.String())
	}
	if err := RBCheck(tree); err != nil {
		t.Errorf("Error: %s.\nTree:\n%s\n", err.Error(), tree.String())
	}

	// Deleting node is black, childs are black, sibling is black and sibling's childs are black
	tree = New(&RBTree{
		20, false, nil,
		&RBTree{10, false, nil, nil, nil},
		&RBTree{30, false, nil, nil, nil},
	})
	expected = New(&RBTree{
		20, false, nil, nil,
		&RBTree{30, true, nil, nil, nil},
	})
	tree.Left.Delete()
	tree = tree.FindRoot()
	if !Compare(tree, expected) {
		t.Fatalf("Wanted:\n%s\nGot:\n%s\n", expected.String(), tree.String())
	}
	if err := RBCheck(tree); err != nil {
		t.Errorf("Error: %s.\nTree:\n%s\n", err.Error(), tree.String())
	}

	// Sibling is red
	tree = New(&RBTree{
		20, false, nil,
		&RBTree{10, false, nil, nil, nil},
		&RBTree{
			30, true, nil,
			&RBTree{25, false, nil, nil, nil},
			&RBTree{35, false, nil, nil, nil},
		},
	})
	expected = New(&RBTree{
		30, false, nil,
		&RBTree{
			20, false, nil, nil,
			&RBTree{25, false, nil, nil, nil},
		},
		&RBTree{35, true, nil, nil, nil},
	})
	tree.Left.Delete()
	tree = tree.FindRoot()
	if !Compare(tree, expected) {
		t.Fatalf("Wanted:\n%s\nGot:\n%s\n", expected.String(), tree.String())
	}
	if err := RBCheck(tree); err != nil {
		t.Errorf("Error: %s.\nTree:\n%s\n", err.Error(), tree.String())
	}
}
