package fn

import (
	"strings"
)

func StringLastIndexAny(s, chars string) (int, error) {
	return strings.LastIndexAny(s, chars), nil
}
