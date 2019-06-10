package utils

type HttpError struct {
	StatusCode int
	Message    string
	Data       interface{}
}

func NewHttpError(code int, message string, data interface{}) *HttpError {
	return &HttpError{StatusCode: code, Message: message, Data: data}
}