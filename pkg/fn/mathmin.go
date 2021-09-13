package fn

import (
	"math"
)

func MathMin(first float64, second float64) (float64, error) {
	return math.Min(first, second), nil
}
