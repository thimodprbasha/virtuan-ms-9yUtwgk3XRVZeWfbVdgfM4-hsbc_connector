package fn

import (
	"strings"
)

func StringHasPrefix(s, prefix string) (bool, error) {
	return strings.HasPrefix(s, prefix), nil
}
