package gt

import (
//	"errors"
)

// References:
// -   [Pointer Comparison in Go](http://golang.org/ref/spec#Comparison_operators)

// Acts, quick way to define the longish Actions array.
type Acts [][][2]int

// Players, or Agents:
// -   Self-Interested.
//     -   They only care about themselves
//     -   Has it own description of the world.
// -   Has a utility function:
//     -   Quantifies degree of prefference across alternatives.
//     -   Decition theory rationality: Act to maximize expected utility.
//
// Should it be an interface?
//
type Player struct {
	Payoff int
	// ## [Strategy Concept](https://en.wikipedia.org/wiki/Strategy_%28game_theory%29)
	// -   One of the options he can choose in a setting where the outcome depends
	//     not only on his own actions but on the actions of others.
	// -   A player's strategy set defines what strategies are available for them
	//     to play.
	// -   A strategy picks the outcome based on the actions...
	Strategies [][]*int
}

// Definition of Utility Function:
//
//     u : X → R
//
// Represents a preference relation on X
// iff for every x, y ∈ X, u(x) ≤ u(y)
// implies x \precsim y
// if u represents \precsim
// then this implies \precsim is complete and transitive,
// and hence rational.
//
// \precsim is preference.
//
// More Theory:
// -   Completeness:
//     All actions can be ranked in an order of preference (indifference between two or more is possible).
//     -   [Partial Order](http://en.wikipedia.org/wiki/Partial_order#Formal_definition)
//     -   [Complete Partial Order](https://en.wikipedia.org/wiki/Complete_partial_order)
//     -   [Completeness](https://en.wikipedia.org/wiki/Completeness_%28order_theory%29)
// -   Transivity: All actions can be compared with other actions.
//     -   [Transitive relation](https://en.wikipedia.org/wiki/Transitive_relation)
//
// func (p *Player) Utility(X []*int) (R int) {
// 	for i := 0; i < len(X); i++ {
// 		R += *X[i]
// 	}
// 	return
// }

// Definition of Best Response:
//
//     a*i ∈ BR(a-i) iff ∀ ai ∈ Ai, ui(a*i, a-i) ≥ ui(ai, a-i)
//
// Where:
// -   BR is the Best Response function.
// -   ai is the action of player i,
// -   a-i is the action of player -i.
// -   A is the set of actions or strategies.
// -   u is the utility function,
//     for normal games, it means getting out the value.
//
// Considerations on the algorithm:
// -   I'm actually using the median of the outcome
//     of each possible strategies per player,
//     there must be a better algorithm.
//     I shall look for a better algorithm!
// -   I'm not taking care of a-i
//     because I removed it from the
//     equation when building the game,
//     and also because I'm using median as
//     the utility function.
//     I wonder if this has implications.
//
// Todos:
// -   The formula left clears that the output
//     is at least a member of the function of BestResponses,
//     so best responses must return a SET.
//
func (p *Player) BestResponse(a_i int) (a []int) {
	var u, best_u int
	for i := 0; i < len(p.Strategies); i++ {
		u = *p.Strategies[i][a_i]
		// Rrtional Preference: ui(a*i) ≥ ui(ai)
		switch {
		case i == 0 || u > best_u:
			a = []int{i}
			best_u = u
		case u == best_u:
			a = append(a, i)
			best_u = u
		}
	}
	return
}

// It could be useful to make it an interface
// if in that way we could use general functions
// over different games.
//
type Game interface {
	NashEquilibrium() [2]int
	ParetoDominates() [2]int
}

// Normal Form Games:
//
type Normal struct {
	Players [2]*Player // Players: N = {1, ..., n}
	Actions Acts       // Actions: (a1, ..., an) ∈ A = A1 × ... × An
}

// Utility function: ui : A → R
// It's per player, so it's a function of each player!
// -   u = (u1, ..., un)
//     is a profile of utility functions.
//
func (g *Normal) Utility(Action [][2]int) (r [][2]int) {
	for _, a := range Action {
		r = append(r, g.Actions[a[0]][a[1]])
	}
	return
}

// Definition of Nash Equilibrium:
//
//    a = ⟨ai, ..., an⟩ iff ∀i, ai ∈ BR(a-i)
//
// Which means, a Nash Equilibrium
// is the set of actions that matches
// the best responses of every player.
//
// -   `a =` leaves clear that the output is at least an action []int
//
func (g *Normal) NashEquilibrium() (a [][2]int) {
	// Fixes:
	// -   I think this can be made with recursion.
	// -   Make this work for games with different shapes.
	var br0, br1 []int
	var _a [2]int
	for i := 0; i < len(g.Players[0].Strategies); i++ {
		br0 = g.Players[1].BestResponse(i)
		br1 = g.Players[0].BestResponse(br0[0]) // This should be changed to a loop
		_a = [2]int{br0[0], br1[0]}
		if i == 0 || _a != a[i-1] {
			a = append(a, _a)
		}
	}
	return
}

// Definition of Pareto Dominates:
//
//    ∀i ≤ n [ui(a*) > ui(s')]
//
// Pareto efficiency is a state of economic allocation of resources in which it
// is impossible to make any one further better off without making at least one
// individual worse off.
//
func (g *Normal) ParetoDominates() (a [][2]int) {
	var ai, aj [2]int
	var add bool
	var lasti, lastj int
loop:
	for i := 0; i < len(g.Actions); i++ {
		for j := 0; j < len(g.Actions[i]); j++ {
			aj = g.Actions[i][j]
			if aj[0] >= ai[0] && aj[1] >= ai[1] {
				ai = aj
				switch add {
				case true:
					if i == lasti && j == lastj {
						continue
					}
					a = append(a, ai)
				case false:
					lasti = i
					lastj = j
				}
			}
		}
	}
	if !add {
		add = true
		a = append(a, ai)
		goto loop
	}
	return
}

// Create a new game with n players
// Set of all actions per player 1
// ⊃ Set of all actions per player 2
// ⊃ Set of outcomes per a1 × a2
func NormalGame(a Acts) (game Normal) {
	game.Players = [2]*Player{new(Player), new(Player)}

	// Building the Sets of Strategies
	var i, j, lena, lena0 int
	lena = len(a)
	lena0 = len(a[0])
	game.Players[0].Strategies = make([][]*int, lena)
	game.Players[1].Strategies = make([][]*int, lena0)
	for i = 0; i < lena; i++ {
		game.Players[0].Strategies[i] = make([]*int, lena0)
		for j = 0; j < lena0; j++ {
			if i == 0 {
				game.Players[1].Strategies[j] = make([]*int, lena)
			}
			game.Players[0].Strategies[i][j] = &a[i][j][0]
			game.Players[1].Strategies[j][i] = &a[i][j][1]
		}
	}

	game.Actions = a
	return
}

// Extensive Form Games:
// -   Players move sequentially, represented as a tree.
// -   Keeps track of what each player knows when he or she makes each decision.
type Extensive struct {
	Players []Player
	Moves   []Move
}
type Move struct {
	Payoff   int
	Player   *Player
	Previous *Move
	Left     *Move
	Right    *Move
}
