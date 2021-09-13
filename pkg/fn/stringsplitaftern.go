package fn

import (
	"strings"
)

func StringSplitAfterN(s, sep string, n int) ([]string, error) {
	return strings.SplitAfterN(s, sep, n), nil
}
