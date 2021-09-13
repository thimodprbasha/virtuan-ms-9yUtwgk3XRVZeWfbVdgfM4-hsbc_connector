package fn

import (
	"crypto/sha256"
	"hash"
)

func Sha256New() (interface{}, error) {
	return sha256.New(), nil
}
