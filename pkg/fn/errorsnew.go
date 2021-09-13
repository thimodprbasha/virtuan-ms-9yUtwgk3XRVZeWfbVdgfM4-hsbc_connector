package fn

import (
	"errors"
)

func ErrorsNew(text string) (interface{}, error) {
	return errors.New(text), nil
}
