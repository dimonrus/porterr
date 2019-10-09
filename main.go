package porterr

// Portable error
type PortError struct {
	ErrorData             // Error data
	details   []ErrorData // Error details
}

// Detailed error
type ErrorData struct {
	Message string      // Error message
	Code    interface{} // Code message
	Name    string      // Name of detail
	stack   []byte      // Stacktrace
}

// Common Error Interface
type IError interface {
	Error() string                                                   // Error interface std
	GetStack() []byte                                                // Get stacktrace
	GetDetails() []IError                                            // Get error details
	PushDetail(message string, code interface{}, name string) IError // Add error detail
	PopDetail() IError                                               // Get detail from
	FlushDetails() IError                                            // Reset all details
}
