package fn

import (
	"strings"
)

func StringContainAny(s, chars string) (bool, error) {
	return strings.ContainsAny(s, chars), nil
}
