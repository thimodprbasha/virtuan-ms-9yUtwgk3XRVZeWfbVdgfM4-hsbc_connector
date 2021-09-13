package fn

import (
	"strings"
)

func StringTrimPrefix(s, prefix string) (string, error) {
	return strings.TrimPrefix(s, prefix), nil
}
