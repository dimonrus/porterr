package porterr

import (
	"fmt"
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
		items = append(items, &PortError{ErrorData: ErrorData{Message: e.details[i].Message, Code: e.details[i].Code, Name: e.details[i].Name, stack: e.details[i].stack}})
	}
	return items
}

// Interface error push detail
func (e *PortError) PushDetail(message string, code interface{}, name string) IError {
	e.details = append(e.details, ErrorData{Message: message, Code: code, Name: name, stack: debug.Stack()})
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

// Flush detail
func (e *PortError) FlushDetails() IError {
	e.details = make([]ErrorData, 0)
	return e
}

// New error
func New(message string, code interface{}, name string) IError {
	return &PortError{
		ErrorData: ErrorData{Message: message, Code: code, Name: name, stack: debug.Stack()},
	}
}

// New error
func NewF(message string, code interface{}, name string, args ...interface{}) IError {
	return &PortError{
		ErrorData: ErrorData{Message: fmt.Sprintf(message, args...), Code: code, Name: name, stack: debug.Stack()},
	}
}