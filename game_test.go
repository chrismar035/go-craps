package craps

import "testing"

func TestNewCrapsGame(t *testing.T) {
	game := CrapsGame{}
	expected_initial_point := 0
	if game.point != expected_initial_point {
		t.Errorf("New game had point set to %v; want %v",
				 game.point,
				 expected_initial_point)
	}
}

func TestIsComingOutNewGame(t *testing.T) {
	game := CrapsGame{}
	if !game.IsComingOut() {
		t.Errorf("New games should not have a point")
	}
}

func TestIsComingOutEstablished(t *testing.T) {
	game := CrapsGame{1, nil}
	if game.IsComingOut() {
		t.Errorf("Games with a point should not be coming out")
	}
}

func TestIsComingOutResetPoint(t *testing.T) {
	game := CrapsGame{1, nil}
	game.point = 0
	if !game.IsComingOut() {
		t.Errorf("Games without a point should be coming out")
	}
}
