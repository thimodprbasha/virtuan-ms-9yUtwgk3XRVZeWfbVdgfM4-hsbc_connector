package fn

import (
	"strings"
)

func StringJoin(elems []string, sep string) (string, error) {
	return strings.Join(elems, sep), nil
}
