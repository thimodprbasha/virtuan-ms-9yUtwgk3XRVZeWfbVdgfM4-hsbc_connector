package fn

import (
	"strings"
)

func StringLastIndex(s, substr string) (int, error) {
	return strings.LastIndex(s, substr), nil
}
