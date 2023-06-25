package errs

import (
	"errors"
	"fmt"
)

func InvalidFormat() error {
	return errors.New(invalidFormat)
}

func EmptyResult() error {
	return errors.New(emptyResult)
}

func HttpConflict(msg string) error {
	return errors.New(fmt.Sprintf("%s: %s", conflict, msg))
}

func HttpNotFound(msg string) error {
	return errors.New(fmt.Sprintf("%s: %s", notFound, msg))
}
