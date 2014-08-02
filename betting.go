package craps

import "math"

type PlayerBet struct {
	Amount int
	Bet Bet
}

type Bet struct {
	toPay int
	base int
}

func (b PlayerBet) Winnings() int {
	odds := float64(b.Bet.toPay) / float64(b.Bet.base)
	return int(math.Floor(float64(b.Amount) * odds))
}

func (b PlayerBet) ToPay() int { return b.Bet.toPay }
func (b PlayerBet) Base() int { return b.Bet.base }


var PassLineOdds =			Bet{1, 1}
var FourTenPassLineOdds =	Bet{2, 1}
var FourTenPlaceOdds =		Bet{9, 5}
var FiveNineOdds =			Bet{3 ,2}
var SixEightOdds =			Bet{6, 5}
