package craps

import "testing"

var payoutTests = []struct {
	bet Bet
	out int
}{
	{Bet{5, Odds{7, 5}}, 7},
	{Bet{10, Odds{7, 5}}, 14},
}

func TestWinnings(t *testing.T) {
	for _, tt := range payoutTests {
		if x := tt.bet.Winnings(); x != tt.out {
			t.Errorf("Bet{%v at %v:%v).Winnings() = %v, want %v",
				tt.bet.amount, tt.bet.odds.toPay, tt.bet.odds.base, x, tt.out)
		}
	}
}
