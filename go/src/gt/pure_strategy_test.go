package gt

import (
	"testing"
)


func Test_Normal_Consistency(t *testing.T) {
	PD := NormalGame([][][2]int{
		{{-1, -1}, {-5,  0}},
		{{ 0, -5}, {-2, -2}},
	})
	for i := 0; i < len(PD.A); i++ {
		for j := 0; j < len(PD.A[0]); j++ {
			if PD.A[i][j][0] != *PD.N[0].Strategies[i][j] {
				t.Errorf("Mismatched Strategies for P0: %v and %v", PD.A[i][j][0], *PD.N[0].Strategies[i][j])
			}
			if PD.A[i][j][1] != *PD.N[1].Strategies[j][i] {
				t.Errorf("Mismatched Strategies for P1: %v and %v", PD.A[i][j][1], *PD.N[1].Strategies[j][i])
			}
		}
	}
	// Testing best response...
	if PD.N[0].BestResponse() != 1 {
		t.Errorf("Wrong Best Response!")
	}
	// Testing players' pointers after changing the game...
	PD.A[0][1][0] = 0
	if PD.N[0].BestResponse() != 0 {
		t.Errorf("Wrong Best Response after the original game changed!")
	}
	return
}



func Test_Prisoners_Dilema(t *testing.T) {
	PD := NormalGame([][][2]int{
		{{-1, -1}, {-5,  0}},
		{{ 0, -5}, {-2, -2}},
	})
	eq := PD.NashEquilibrium()
	if eq[0] != -2 && eq[1] != -2 {
		t.Errorf("Wrong Nash Equilibrium: %v", eq)
	}
	return
}


func Test_Coordination(t *testing.T) {
	PD := NormalGame([][][2]int{
		{{4, 4}, {1, 3}},
		{{3, 1}, {2, 2}},
	})
	eq := PD.NashEquilibrium()
	if eq[0] != 4 && eq[1] != 4 {
		t.Errorf("Wrong Nash Equilibrium: %v", eq)
	}
	return
}
