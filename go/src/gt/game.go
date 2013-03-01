package gt

import (
//	"errors"
)

// References:
// -   [Pointer Comparison in Go](http://golang.org/ref/spec#Comparison_operators)


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
	//     not only on his own actions but on the action of others.
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
// I'm using a mean as an utility function.
// Means can be compared to the estimated outcome
// of a strategy based on the other player strategies
// (which correspond to the elements in the set A).
//
func (p *Player) Utility(A []*int) int {
	var m, l int
	l = len(A)
	for i := 0; i < l; i++ {
		m += *A[i]
	}
	return m/l
}

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
func (p *Player) BestResponse() (strategy int) {
	var u, bestu int
	for s := 0; s < len(p.Strategies); s++ {
		u = p.Utility(p.Strategies[s])
		// Rational Preference: ui(a*i) ≥ ui(ai)
		if s == 0 || u > bestu {
			bestu = u
			strategy = s
		}
	}
	return
}



// It could be useful to make it an interface
// if in that way we could use general functions
// over different games.
type Game interface {
	NashEquilibrium() [2]int
}

// Normal Form Games:
// -   Players: N = {1, ..., n}
// -   Actions: (a1, ..., an) ∈ A = A1 × ... × An
//     is an action profile.
// -   Utility function: ui : A → R
//     It's per player, so it's a function of each player!
//     -   u = (u1, ..., un)
//         is a profile of utility functions.
type Normal struct {
	N [2]*Player
	A [][][2]int // Action
}

// Definition of Nash Equilibrium:
//
//    a = ⟨ai, ..., an⟩ iff ∀i, ai ∈ BR(a-i)
//
// Which means, a Nash Equilibrium
// is the set of actions that matches
// the best responses of every player.
//
func (g *Normal) NashEquilibrium() [2]int {
	br0 := g.N[0].BestResponse()
	br1 := g.N[1].BestResponse()
	return g.A[br0][br1]
}

// Create a new game with n players
// Set of all actions per player 1
// ⊃ Set of all actions per player 2
// ⊃ Set of outcomes per a1 × a2
func NormalGame(a [][][2]int) (game Normal) {
	game.N = [2]*Player{new(Player), new(Player)}

	// Building the Sets of Strategies
	var i, j, lena, lena0 int
	lena = len(a)
	lena0 = len(a[0])
	game.N[0].Strategies = make([][]*int, lena)
	game.N[1].Strategies = make([][]*int, lena0)
	for i = 0; i < lena; i++ {
		game.N[0].Strategies[i] = make([]*int, lena0)
		for j = 0; j < lena0; j++ {
			if i == 0 {
				game.N[1].Strategies[j] = make([]*int, lena)
			}
			game.N[0].Strategies[i][j] = &a[i][j][0]
			game.N[1].Strategies[j][i] = &a[i][j][1]
		}
	}

	game.A = a
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


