package craps

import "math/rand"

type Dice struct {
	FirstDie int
	SecondDie int
}

func (d Dice) Roll() {
	d.FirstDie = dieRoll()
	d.SecondDie = dieRoll()
}

func (d Dice) Value() int {
	return d.FirstDie + d.SecondDie
}

func (d Dice) Pointable() bool {
	switch d.Value() {
	case 2, 3, 7, 11, 12:
		return false
	}
	return true
}

func (d Dice) Crap() bool {
	switch d.Value() {
	case 2, 3, 12:
		return true
	}
	return false
}

func dieRoll() int {
	return 1 + rand.Intn(6)
}
