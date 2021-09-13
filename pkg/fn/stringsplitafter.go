package fn

import (
	"strings"
)

func StringSplitAfter(s, sep string) ([]string, error) {
	return strings.SplitAfter(s, sep), nil
}
