package fn

import (
	"strings"
)

func StringContains(s, substr string) (bool, error) {
	return strings.Contains(s, substr), nil
}
