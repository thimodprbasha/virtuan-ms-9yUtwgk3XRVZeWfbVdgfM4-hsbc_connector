package fn

import (
	"crypto/sha1"
	"hash"
)

func Sha1New() (interface{}, error) {
	return sha1.New(), nil
}
