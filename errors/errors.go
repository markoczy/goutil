package errors

//
// More advanced Error Handling...
//
//


// ErrorType ...
type ErrorType interface {
	Error() string
	Code() int
	Cause() error
}

type errorType struct {
	mError string
	mCode int
	mCause error
}

// New ...
func New(aError string, aCode int, aCause error) ErrorType {
	return &errorType{ aError, aCode, aCause }
}

func (err *errorType) Error() string {
	return err.mError
}

func (err *errorType) Code() int {
	return err.mCode
}

func (err *errorType) Cause() error {
	return err.mCause
}