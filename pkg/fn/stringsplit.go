package fn

import (
	"strings"
)

func StringSplit(s, sep string) ([]string, error) {
	return strings.Split(s, sep), nil
}
