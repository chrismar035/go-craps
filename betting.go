package craps

import "math"

type Bet struct {
	amount int
	odds Odds
}

type Odds struct {
	toPay int
	base int
}

func (b Bet) Winnings() int {
	odds := float64(b.odds.toPay) / float64(b.odds.base)
	return int(math.Floor(float64(b.amount) * odds))
}

