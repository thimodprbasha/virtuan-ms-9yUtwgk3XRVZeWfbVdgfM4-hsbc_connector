package fn

import (
	"math"
)

func MathPow(first float64, second float64) (float64, error) {
	return math.Pow(first, second), nil
}
