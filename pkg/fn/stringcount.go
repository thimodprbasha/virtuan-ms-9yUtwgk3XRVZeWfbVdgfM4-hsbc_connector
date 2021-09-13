package fn

import (
	"strings"
)

func StringCount(s, substr string) (int, error) {
	return strings.Count(s, substr), nil
}
