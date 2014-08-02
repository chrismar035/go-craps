package craps

import "math/rand"

func Roll() int {
	return 1 + rand.Intn(6)
}
