package porterr

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Interface error method
func (e *PortError) Error() string {
	return e.ErrorData.Message
}

// Interface error get stack
func (e *PortError) GetStack() []byte {
	return e.stack
}

// Interface error get details
func (e *PortError) GetDetails() []IError {
	var items []IError
	for i := range e.details {
		items = append(items, &PortError{httpCode: e.httpCode, ErrorData: ErrorData{Message: e.details[i].Message, Code: e.details[i].Code, Name: e.details[i].Name, stack: e.details[i].stack}})
	}
	return items
}

// Interface error push detail
func (e *PortError) PushDetail(code interface{}, name string, message string) IError {
	e.details = append(e.details, ErrorData{Code: code, Name: name, Message: message, stack: debug.Stack()})
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
func (e *PortError) MergeDetails(error ... IError) IError {
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
		httpCode:http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Message: message, stack: debug.Stack()},
	}
}

// New error
func NewF(code interface{}, message string, args ...interface{}) IError {
	return &PortError{
		httpCode:http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Message: fmt.Sprintf(message, args...), stack: debug.Stack()},
	}
}

// New error with name
func NewWithName(code interface{}, name string, message string) IError {
	return &PortError{
		httpCode:http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Name: name, Message: message, stack: debug.Stack()},
	}
}

// New error with name
func NewFWithName(code interface{}, name string, message string, args ...interface{}) IError {
	return &PortError{
		httpCode:http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Name: name, Message: fmt.Sprintf(message, args...), stack: debug.Stack()},
	}
}
