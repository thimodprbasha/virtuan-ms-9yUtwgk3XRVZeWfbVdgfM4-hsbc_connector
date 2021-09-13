package fn

import (
	"strings"
)

func StringTrim(s, cutset string) (string, error) {
	return strings.Trim(s, cutset), nil
}
