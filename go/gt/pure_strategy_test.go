package gt

import (
	"testing"
)

func Test_Normal_Consistency(t *testing.T) {
	PD := NormalGame(Acts{
		{{-1, -1}, {-5, 0}},
		{{0, -5}, {-2, -2}},
	})
	// Checking the consistency of strategies
	for i := 0; i < len(PD.Actions); i++ {
		for j := 0; j < len(PD.Actions[0]); j++ {
			if PD.Actions[i][j][0] != *PD.Players[0].Strategies[i][j] {
				t.Errorf("Mismatched Strategies for P0: %v and %v", PD.Actions[i][j][0], *PD.Players[0].Strategies[i][j])
			}
			if PD.Actions[i][j][1] != *PD.Players[1].Strategies[j][i] {
				t.Errorf("Mismatched Strategies for P1: %v and %v", PD.Actions[i][j][1], *PD.Players[1].Strategies[j][i])
			}
		}
	}
	// Testing best response...
	if PD.Players[0].BestResponse(0)[0] != 1 {
		t.Errorf("Wrong Best Response!")
	}
	// Testing players' pointers after changing the game...
	PD.Actions[0][1][0] = 0
	PD.Actions[1][0][1] = 0
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
	game := NormalGame(Acts{
		{{-1, -1}, {-5, 0}},
		{{0, -5}, {-2, -2}},
	})
	eq := game.NashEquilibrium()
	eu := game.Utility(eq)
	if eu[0] != [2]int{-2, -2} {
		t.Errorf("Wrong Best Response! \n-   BR1: %v\n-   BR2: %v", game.Players[0].BestResponse(0), game.Players[1].BestResponse(0))
		t.Errorf("Wrong Nash Equilibrium! \n-   Strategy: %v\n-   Utility: %v", eq, eu)
	}
	return
}

// Next: Battle of the Sexes's Two Pure Strategy Nash Equilibrium.
//
// [Wiki](https://en.wikipedia.org/wiki/Battle_of_the_sexes_%28game_theory%29)
//
func Test_BoS(t *testing.T) {
	game := NormalGame(Acts{
		{{3, 2}, {0, 0}},
		{{0, 0}, {2, 3}},
	})
	eq := game.NashEquilibrium()
	eu := game.Utility(eq)
	if eq[0] != [2]int{0, 0} && eq[1] != [2]int{1, 1} && eu[0] != [2]int{3, 2} {
		t.Errorf("Wrong Nash Equilibrium! \n-   Strategy: %v\n-   Utility: %v\n", eq, eu)
		t.Errorf("\nP0 BR: %v", game.Players[0].BestResponse(0))
		t.Errorf("\nP1 BR: %v", game.Players[1].BestResponse(0))
		// t.Errorf("\nP0 Ss: {{%v %v}{%v %v}}", *game.Players[0].Strategies[0][0], *game.Players[0].Strategies[0][1], *game.Players[0].Strategies[1][0], *game.Players[0].Strategies[1][1])
	}
	return
}

// From: http://planetmath.org/examplesofnormalformgames
func Test_Moar(t *testing.T) {
	var n_test int
	game := NormalGame(Acts{
		{{5, 5}, {-5, 10}},
		{{10, -5}, {0, 0}},
	})
	eq := game.NashEquilibrium()
	eu := game.Utility(eq)
	if !(eq[0] == [2]int{1, 1} && eu[0] == [2]int{0, 0}) {
		t.Errorf("\nTest_Moar N%v: game.NashEquilibrium\nWanted: %v\nGot: %v", n_test, [2]int{1, 1}, eq[0])
	}
	pd := game.ParetoDominates()
	if pd[0] != [2]int{5, 5} {
		t.Errorf("\nTest_Moar N%v: game.ParetoDominates\nWanted: %v\nGot: %v", n_test, [2]int{1, 1}, pd[0])
	}
	n_test++
	game = NormalGame(Acts{
		{{5, 5}, {5, 5}},
		{{10, -5}, {0, 0}},
	})
	eq = game.NashEquilibrium()
	eu = game.Utility(eq)
	if !(eq[0] == [2]int{0, 1} && eu[0] == [2]int{5, 5}) {
		t.Errorf("\nTest_Moar N%v: game.NashEquilibrium\nWanted: %v\nGot: %v", n_test, [2]int{1, 1}, eq[0])
	}
	pd = game.ParetoDominates()
	if pd[0] != [2]int{5, 5} && pd[1] != [2]int{5, 5} {
		t.Errorf("\nTest_Moar N%v: game.ParetoDominates\nWanted: %v\nGot: %v", n_test, [2]int{1, 1}, pd[0])
	}

}
