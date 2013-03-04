package gt

import (
	"testing"
)

func Test_Normal_Consistency(t *testing.T) {
	PD := NormalGame([][][2]int{
		{{-1, -1}, {-5, 0}},
		{{0, -5}, {-2, -2}},
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
	if PD.N[0].BestResponse(0)[0] != 1 {
		t.Errorf("Wrong Best Response!")
	}
	// Testing players' pointers after changing the game...
	PD.A[0][1][0] = 0
	PD.A[1][0][1] = 0
	if PD.NashEquilibrium()[0] != [2]int{1, 0} {
		t.Errorf("Wrong Best Response after the original game changed!")
	}
	return
}

// [Prisoner's dilemma](https://en.wikipedia.org/wiki/Prisoner%27s_dilemma)
//
// Canonical PD payoff matrix:
//
//                Cooperate   Defect
//     Cooperate { { R, R }, { S, T } }
//     Defect    { { S, T }, { P, P } }
//
// For it to be a prisoner's dilemma game in the strong sense, the following
// condition must hold for the payoffs:
//
//     T > R > P > S
//
func Test_Prisoners_Dilema(t *testing.T) {
	game := NormalGame([][][2]int{
		{{-1, -1}, {-5, 0}},
		{{0, -5}, {-2, -2}},
	})
	eq := game.NashEquilibrium()
	eu := game.Utility(eq)
	if eu[0] != [2]int{-2, -2} {
		t.Errorf("Wrong Best Response! \n-   BR1: %v\n-   BR2: %v", game.N[0].BestResponse(0), game.N[1].BestResponse(0))
		t.Errorf("Wrong Nash Equilibrium! \n-   Strategy: %v\n-   Utility: %v", eq, eu)
	}
	return
}

// Next: Battle of the Sexes's Two Pure Strategy Nash Equilibrium.
//
// [Wiki](https://en.wikipedia.org/wiki/Battle_of_the_sexes_%28game_theory%29)
//
func Test_BoS(t *testing.T) {
	game := NormalGame([][][2]int{
		{{3, 2}, {0, 0}},
		{{0, 0}, {2, 3}},
	})
	eq := game.NashEquilibrium()
	eu := game.Utility(eq)
	if eq[0] != [2]int{0, 0} && eq[1] != [2]int{1, 1} && eu[0] != [2]int{3, 2} {
		t.Errorf("Wrong Nash Equilibrium! \n-   Strategy: %v\n-   Utility: %v\n", eq, eu)
		t.Errorf("\nP0 BR: %v", game.N[0].BestResponse(0))
		t.Errorf("\nP1 BR: %v", game.N[1].BestResponse(0))
		// t.Errorf("\nP0 Ss: {{%v %v}{%v %v}}", *game.N[0].Strategies[0][0], *game.N[0].Strategies[0][1], *game.N[0].Strategies[1][0], *game.N[0].Strategies[1][1])
	}
	return
}
