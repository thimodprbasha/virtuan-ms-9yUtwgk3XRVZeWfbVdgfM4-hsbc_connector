package fn

import (
	"math"
)

func MathRound(first float64) (float64, error) {
	return math.Round(first), nil
}
