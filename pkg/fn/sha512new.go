package fn

import (
	"crypto/sha512"
	"hash"
)

func Sha512New() (interface{}, error) {
	return sha512.New(), nil
}
