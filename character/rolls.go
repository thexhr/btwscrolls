package character

import (
	"math/rand"
	"time"
)

func dice(side int) int {
	r := rand.New(time.Now())
	return r.Intn(side)
}
