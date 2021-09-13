package fn

import (
	"math"
)

func MathMod(first float64, second float64) (float64, error) {
	return math.Mod(first, second), nil
}
