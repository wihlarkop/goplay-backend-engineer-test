package helper

import "net/http"

const (
	ErrorCodeGeneralError = "GENERAL-01"
	ErrorInvalidParameter = "PAYLOAD-01"
)

var (
	ErrJSONParse  = NewErr(http.StatusBadRequest, "PAYLOAD-01", "Something wrong with input")
	ErrFatalQuery = NewErr(http.StatusInternalServerError, "QUERY-01", "fatal query error")
)

type Err struct {
	errorCode      string
	message        string
	httpStatusCode int
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) GetHttpStatusCode() int {
	return e.httpStatusCode
}

func (e *Err) GetErrorCode() string {
	return e.errorCode
}

func NewErr(httpStatusCode int, errorCode, message string) *Err {
	return &Err{
		httpStatusCode: httpStatusCode,
		errorCode:      errorCode,
		message:        message,
	}
}
