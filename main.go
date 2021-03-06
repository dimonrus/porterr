package porterr

// Portable error
type PortError struct {
	httpCode  int         // Http Error code
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
	PushDetail(code interface{}, name string, message string) IError // Add error detail
	PopDetail() IError                                               // Get detail from
	MergeDetails(e ...IError) IError                                 // Merge detail to error
	FlushDetails() IError                                            // Reset all details
	HTTP(httpCode int) IError                                        // Set HTTP code
	GetHTTP() int                                                    // Get http code
	Code(code interface{}) IError                                    // Set error code
	GetCode() interface{}                                            // Get error code
	IfDetails() IError                                               // Return error if error has details else nil
}
