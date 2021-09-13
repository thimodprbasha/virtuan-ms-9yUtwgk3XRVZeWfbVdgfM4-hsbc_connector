package fn

import (
	"strings"
)

func StringTrimSuffix(s, suffix string) (string, error) {
	return strings.TrimSuffix(s, suffix), nil
}
