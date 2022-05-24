package porterr

import (
	"fmt"
	"net/http"
)

// Interface error method
func (e *PortError) Error() string {
	return e.ErrorData.Message
}

// Interface error get details
func (e *PortError) GetDetails() []IError {
	if e == nil {
		return nil
	}
	var items []IError
	for i := range e.details {
		items = append(items, &PortError{httpCode: e.httpCode, ErrorData: ErrorData{Message: e.details[i].Message, Code: e.details[i].Code, Name: e.details[i].Name}})
	}
	return items
}

// return origin error
func (e *PortError) Origin() *PortError {
	return e
}

// Interface error push detail
func (e *PortError) PushDetail(code interface{}, name string, message string) IError {
	e.details = append(e.details, ErrorData{Code: code, Name: name, Message: message})
	return e
}

// Pop detail
func (e *PortError) PopDetail() IError {
	if len(e.details) > 0 {
		var item *ErrorData
		item, e.details = &e.details[len(e.details)-1], e.details[:len(e.details)-1]
		return &PortError{ErrorData: *item}
	}
	return nil
}

// Merge detail from other errors
func (e *PortError) MergeDetails(error ...IError) IError {
	for _, v := range error {
		if v == nil {
			continue
		}
		detail := v.(*PortError)
		e.details = append(e.details, detail.details...)
	}
	return e
}

// Flush detail
func (e *PortError) FlushDetails() IError {
	e.details = make([]ErrorData, 0)
	return e
}

// Return nil when error does not contain details
func (e *PortError) IfDetails() IError {
	if e == nil {
		return nil
	}
	if len(e.details) > 0 {
		return e
	}
	return nil
}

// Set HTTP
func (e *PortError) HTTP(httpCode int) IError {
	e.httpCode = httpCode
	return e
}

// Get HTTP Code
func (e *PortError) GetHTTP() int {
	return e.httpCode
}

// Set Error Code
func (e *PortError) Code(code interface{}) IError {
	e.ErrorData.Code = code
	return e
}

// Get Error Code
func (e *PortError) GetCode() interface{} {
	return e.ErrorData.Code
}

// New error
func New(code interface{}, message string) IError {
	return &PortError{
		httpCode:  http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Message: message},
	}
}

// New error
func NewF(code interface{}, message string, args ...interface{}) IError {
	return &PortError{
		httpCode:  http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Message: fmt.Sprintf(message, args...)},
	}
}

// New error with name
func NewWithName(code interface{}, name string, message string) IError {
	return &PortError{
		httpCode:  http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Name: name, Message: message},
	}
}

// New error with name
func NewFWithName(code interface{}, name string, message string, args ...interface{}) IError {
	return &PortError{
		httpCode:  http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Name: name, Message: fmt.Sprintf(message, args...)},
	}
}

// HttpValidationError prepare validation error with 400 http code
func HttpValidationError(args ...string) IError {
	var code = PortErrorValidation
	var message = "Validation error"
	if len(args) > 0 {
		code = args[0]
		if len(args) > 1 {
			message = args[1]
		}
	}
	return &PortError{
		httpCode:  http.StatusBadRequest,
		ErrorData: ErrorData{Code: code, Message: message},
	}
}
