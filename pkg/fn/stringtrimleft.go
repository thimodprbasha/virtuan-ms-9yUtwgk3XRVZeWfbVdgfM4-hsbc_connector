package fn

import (
	"strings"
)

func StringTrimLeft(s, cutset string) (string, error) {
	return strings.TrimLeft(s, cutset), nil
}
