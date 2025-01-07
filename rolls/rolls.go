package rolls

import (
	"math/rand"
)

func RollDice(number int, side int) []int {
	result := make([]int, number)
	for i := range number {
		result[i] = rand.Intn(side) + 1
	}

	return result
}


