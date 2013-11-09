package util

import (
	"errors"
	"fmt"
)

func Errorf(format string, vars ...interface{}) error {
	return errors.New(fmt.Sprintf(format, vars...))
}
