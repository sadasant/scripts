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
