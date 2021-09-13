package fn

import (
	"strings"
)

func StringCompare(a, b string) (int, error) {
	return strings.Compare(a, b), nil
}
