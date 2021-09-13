package fn

import (
	"strings"
)

func StringHasSuffix(s, suffix string) (bool, error) {
	return strings.HasSuffix(s, suffix), nil
}
