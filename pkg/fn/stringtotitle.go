package fn

import (
	"strings"
)

func StringToTitle(s string) (string, error) {
	return strings.ToTitle(s), nil
}
