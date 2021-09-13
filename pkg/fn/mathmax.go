package fn

import (
	"math"
)

func MathMax(first float64, second float64) (float64, error) {
	return math.Max(first, second), nil
}
