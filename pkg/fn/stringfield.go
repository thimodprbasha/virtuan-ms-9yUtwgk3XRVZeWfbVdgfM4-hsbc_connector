package fn

import (
	"strings"
)

func StringFields(s string) ([]string, error) {
	return strings.Fields(s), nil
}
