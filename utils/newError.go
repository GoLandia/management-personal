package utils

type Error struct {
	err error
}

func NewError(err error) *Error {
	return &Error{err: err}
}

func (e *Error) HandleError() string {
	if e.err != nil {
		return e.err.Error()
	}

	return ""
}
