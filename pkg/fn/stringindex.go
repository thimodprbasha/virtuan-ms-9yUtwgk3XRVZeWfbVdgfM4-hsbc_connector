package fn

import (
	"strings"
)

func StringIndex(s, substr string) (int, error) {
	return strings.Index(s, substr), nil
}
