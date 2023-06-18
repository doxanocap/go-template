package errs

import (
	"errors"
)

func InvalidFormat() error {
	return errors.New("invalid file format")
}

func IsInvalidFormatError(err error) bool {
	return err.Error() == "invalid file format"
}

func EmptyResult() error {
	return errors.New("empty result")
}

func IsEmptyResultError(err error) bool {
	return err.Error() == "empty result"
}
