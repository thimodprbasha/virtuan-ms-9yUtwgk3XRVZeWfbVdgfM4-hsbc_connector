package fn

import (
	"regexp"
)

func RegexpMatchString(pattern string, s string) (bool, error) {
	matched, _ := regexp.MatchString(pattern, s)
	return matched, nil
}
