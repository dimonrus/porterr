package porterr

// PortError Portable error
type PortError struct {
	// Http Error code
	httpCode int
	// Error data
	ErrorData
	// Error details
	details []ErrorData
}

// ErrorData Detailed error
type ErrorData struct {
	// Code message
	Code interface{}
	// Name of detail
	Name string
	// Error message
	Message string
}

// IError Common Error Interface
type IError interface {
	// Code Set error code
	Code(code interface{}) IError
	// Error interface std
	Error() string
	// FlushDetails Reset all details
	FlushDetails() IError
	// GetCode Get error code
	GetCode() interface{}
	// GetDetails Get error details
	GetDetails() []IError
	// GetHTTP Get http code
	GetHTTP() int
	// HTTP Set HTTP code
	HTTP(httpCode int) IError
	// IfDetails Return error if error has details else nil
	IfDetails() IError
	// MergeDetails Merge detail to error
	MergeDetails(e ...IError) IError
	// Origin Get portable error
	Origin() *PortError
	// PopDetail Get detail from
	PopDetail() IError
	// PushDetail Add error detail
	PushDetail(code interface{}, name string, message string) IError
}
