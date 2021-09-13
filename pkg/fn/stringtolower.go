package fn

import (
	"strings"
)

func StringToLower(s string) (string, error) {
	return strings.ToLower(s), nil
}
