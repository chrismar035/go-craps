package craps

import "testing"

var payoutTests = []struct {
	bet PlayerBet
	out int
}{
	{PlayerBet{5, Bet{7, 5}}, 7},
	{PlayerBet{10, Bet{7, 5}}, 14},
	{PlayerBet{10, PassLineOdds}, 10},
	{PlayerBet{10, FourTenPassLineOdds}, 20},
}

func TestWinnings(t *testing.T) {
	for _, tt := range payoutTests {
		if x := tt.bet.Winnings(); x != tt.out {
			t.Errorf("Bet{%v at %v:%v).Winnings() = %v, want %v",
				tt.bet.Amount, tt.bet.ToPay(), tt.bet.Base(), x, tt.out)
		}
	}
}
