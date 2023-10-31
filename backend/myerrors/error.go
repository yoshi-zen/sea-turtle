package myerrors

type MyError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (myErr *MyError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyError) UnWrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyError{ErrCode: code, Message: message, Err: err}
}
