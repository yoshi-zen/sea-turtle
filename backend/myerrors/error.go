package myerrors

type MyError struct {
	ErrCode
	Message string
	Err     error
}

func (myErr *MyError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyError) UnWrap() error {
	return myErr.Err
}
