package fn

import (
	"strings"
)

func StringRepeat(s string, count int) (string, error) {
	return strings.Repeat(s, count), nil
}
