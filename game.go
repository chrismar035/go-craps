package craps

import "fmt"

type CrapsGame struct {
	Point int
	Bets []PlayerBet
	dice Dice
	Winners []PlayerBet
}

func (g CrapsGame) FirstDie() int {
	return g.dice.FirstDie
}

func (g CrapsGame) SecondDie() int {
	return g.dice.SecondDie
}

func (g CrapsGame) RollValue() int {
	return g.dice.Value()
}

func (g CrapsGame) IsComingOut() bool {
	return g.Point == 0
}

func (g CrapsGame) Roll() {
	g.Winners = nil

	g.dice.Roll()

	if g.IsComingOut() {
		if g.dice.Pointable() {
			g.Point = g.dice.Value()
		} else if g.dice.Crap() {
			// lose pass line bets
			g.Bets = nil

		} else {
			// front line winner
			g.Winners = g.Bets
		}
	} else {
		switch g.dice.Value() {
		case 7:
			g.Bets = nil
			g.Point = 0
		case g.Point:
			g.Winners = g.Bets
			g.Bets = nil
			g.Point = 0
		}
	}
}

func (g CrapsGame) PlaceBet(newBet PlayerBet) {
	fmt.Printf("Before: %v\n", len(g.Bets))
	if cap(g.Bets) == 0 {
		fmt.Println("new")
		g.Bets = []PlayerBet{newBet}
	} else {
		g.Bets = append(g.Bets, newBet)
	}
	fmt.Printf("After: %v\n", g.Bets[0].Amount)
}
