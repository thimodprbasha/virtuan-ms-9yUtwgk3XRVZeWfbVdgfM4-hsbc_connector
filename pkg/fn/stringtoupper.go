package fn

import (
	"strings"
)

func StringToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}
