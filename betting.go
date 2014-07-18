package craps

import "math"

type PlayerBet struct {
	amount int
	bet Bet
}

type Bet struct {
	toPay int
	base int
}

func (b PlayerBet) Winnings() int {
	odds := float64(b.bet.toPay) / float64(b.bet.base)
	return int(math.Floor(float64(b.amount) * odds))
}

func (b PlayerBet) toPay() int { return b.bet.toPay }
func (b PlayerBet) base() int { return b.bet.base }


var passLineOdds =			Bet{1, 1}
var fourTenPassLineOdds =	Bet{2, 1}
var fourTenPlaceOdds =		Bet{9, 5}
var fiveNineOdds =			Bet{3 ,2}
var sixEightOdds =			Bet{6, 5}
