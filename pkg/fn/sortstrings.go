package fn

import (
	"sort"
)

func SortStrings(x []string) ([]string, error) {
	sort.Strings(x)
	return x, nil
}
