package fn

import (
	"strings"
)

func StringSplitN(s, sep string, n int) ([]string, error) {
	return strings.SplitN(s, sep, n), nil
}
