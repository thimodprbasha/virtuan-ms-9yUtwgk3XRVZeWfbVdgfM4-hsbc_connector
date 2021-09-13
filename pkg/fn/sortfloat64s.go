package fn

import (
	"sort"
)

func SortFloat64s(x []float64) ([]float64, error) {
	sort.Float64s(x)
	return x, nil
}
