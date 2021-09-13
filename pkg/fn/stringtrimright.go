package fn

import (
	"strings"
)

func StringTrimRight(s, cutset string) (string, error) {
	return strings.TrimRight(s, cutset), nil
}
