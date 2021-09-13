package fn

import (
	"strings"
)

func StringIndexAny(s, chars string) (int, error) {
	return strings.IndexAny(s, chars), nil
}
