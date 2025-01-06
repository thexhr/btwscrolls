package character

import (
	"math/rand"
)

func dice(side int) int {
	r := rand.New(time.Now())
	return r.Intn(side)
}
