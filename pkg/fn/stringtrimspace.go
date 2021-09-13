package fn

import (
	"strings"
)

func StringTrimSpace(s string) (string, error) {
	return strings.TrimSpace(s), nil
}
