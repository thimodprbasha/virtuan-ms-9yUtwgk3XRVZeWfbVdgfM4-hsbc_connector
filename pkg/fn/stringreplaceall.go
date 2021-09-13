package fn

import (
	"strings"
)

func StringReplaceAll(replaceString string, old string, new string) (string, error) {
	return strings.ReplaceAll(replaceString, old, new), nil
}
