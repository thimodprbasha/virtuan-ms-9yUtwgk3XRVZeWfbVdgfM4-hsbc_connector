package fn

import (
	"strings"
)

func StringTitle(s string) (string, error) {
	return strings.Title(s), nil
}
