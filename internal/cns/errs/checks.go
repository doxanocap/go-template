package errs

func IsInvalidFormatError(err error) bool {
	return err.Error() == invalidFormat
}

func IsEmptyResultError(err error) bool {
	return err.Error() == emptyResult
}

func IsHttpConflictError(err error) bool {
	if len(err.Error()) < len(conflict) {
		return false
	}
	return err.Error()[0:len(conflict)] == conflict
}

func IsHttpNotFoundError(err error) bool {
	if len(err.Error()) < len(notFound) {
		return false
	}
	return err.Error()[0:len(notFound)] == notFound
}
