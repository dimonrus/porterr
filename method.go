package porterr

import (
	"fmt"
	"net/http"
)

// Interface error method
func (e *PortError) Error() string {
	return e.ErrorData.Message
}

// GetDetails Interface error get details
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

// Origin return origin error
func (e *PortError) Origin() *PortError {
	return e
}

// PushDetail Interface error push detail
func (e *PortError) PushDetail(code interface{}, name string, message string) IError {
	e.details = append(e.details, ErrorData{Code: code, Name: name, Message: message})
	return e
}

// PopDetail Pop detail
func (e *PortError) PopDetail() IError {
	if len(e.details) > 0 {
		var item *ErrorData
		item, e.details = &e.details[len(e.details)-1], e.details[:len(e.details)-1]
		return &PortError{ErrorData: *item}
	}
	return nil
}

// MergeDetails Merge detail from other errors
func (e *PortError) MergeDetails(error ...IError) IError {
	for _, v := range error {
		if v == nil {
			continue
		}
		if pe, ok := v.(*PortError); ok {
			e.details = append(e.details, pe.details...)
		}
	}
	return e
}

// AsDetails append to details list of IError
func (e *PortError) AsDetails(error ...IError) IError {
	for _, iError := range error {
		if pe, ok := iError.(*PortError); ok {
			e.details = append(e.details, pe.ErrorData)
		}
	}
	return e
}

// Is check is IError is same as origin
func (e *PortError) Is(err error) bool {
	pe, ok := err.(*PortError)
	if !ok {
		return false
	}
	if pe.httpCode != 0 && e.httpCode != 0 {
		if pe.httpCode != e.httpCode {
			return false
		}
	}
	if len(pe.details) != len(e.details) {
		return false
	}
	if pe.Name != e.Name {
		return false
	}
	if pe.Message != e.Message {
		return false
	}
	if pe.ErrorData.Code != e.ErrorData.Code {
		return false
	}
	for i := range pe.details {
		if pe.details[i].Name != e.details[i].Name {
			return false
		}
		if pe.details[i].Message != e.details[i].Message {
			return false
		}
		if pe.details[i].Code != e.details[i].Code {
			return false
		}
	}
	return true
}

// FlushDetails Flush detail
func (e *PortError) FlushDetails() IError {
	e.details = e.details[:0]
	return e
}

// IfDetails Return nil when error does not contain details
func (e *PortError) IfDetails() IError {
	if e == nil {
		return nil
	}
	if len(e.details) > 0 {
		return e
	}
	return nil
}

// HTTP Set HTTP
func (e *PortError) HTTP(httpCode int) IError {
	e.httpCode = httpCode
	return e
}

// GetHTTP Get HTTP Code
func (e *PortError) GetHTTP() int {
	return e.httpCode
}

// Code Set Error Code
func (e *PortError) Code(code interface{}) IError {
	e.ErrorData.Code = code
	return e
}

// GetCode Get Error Code
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

// NewF error
func NewF(code interface{}, message string, args ...interface{}) IError {
	return &PortError{
		httpCode:  http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Message: fmt.Sprintf(message, args...)},
	}
}

// NewWithName error with name
func NewWithName(code interface{}, name string, message string) IError {
	return &PortError{
		httpCode:  http.StatusInternalServerError,
		ErrorData: ErrorData{Code: code, Name: name, Message: message},
	}
}

// NewFWithName error with name format
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
