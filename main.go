package porterr

// PortError Portable error
type PortError struct {
	httpCode  int         // Http Error code
	ErrorData             // Error data
	details   []ErrorData // Error details
}

// ErrorData Detailed error
type ErrorData struct {
	Message string      // Error message
	Code    interface{} // Code message
	Name    string      // Name of detail
}

// IError Common Error Interface
type IError interface {
	Code(code interface{}) IError                                    // Set error code
	Error() string                                                   // Error interface std
	FlushDetails() IError                                            // Reset all details
	GetCode() interface{}                                            // Get error code
	GetDetails() []IError                                            // Get error details
	GetHTTP() int                                                    // Get http code
	HTTP(httpCode int) IError                                        // Set HTTP code
	IfDetails() IError                                               // Return error if error has details else nil
	MergeDetails(e ...IError) IError                                 // Merge detail to error
	Origin() *PortError                                              // Get portable error
	PopDetail() IError                                               // Get detail from
	PushDetail(code interface{}, name string, message string) IError // Add error detail
}
