package pe

import (
	"encoding/json"
)

const (
	PortErrorSystem        = "PORTABLE_ERROR_SYSTEM"
	PortErrorClient        = "PORTABLE_ERROR_CLIENT"
	PortErrorDatabaseQuery = "PORTABLE_ERROR_DATABASE"
	PortErrorConsole       = "PORTABLE_ERROR_CONSOLE"
	PortErrorLogic         = "PORTABLE_ERROR_LOGIC"
	PortErrorRequest       = "PORTABLE_ERROR_REQUEST"
	PortErrorResponse      = "PORTABLE_ERROR_RESPONSE"
	PortErrorConflict      = "PORTABLE_ERROR_CONFLICT"
	PortErrorTransaction   = "PORTABLE_ERROR_TRANSACTION"
	PortErrorConnection    = "PORTABLE_ERROR_CONNECTION"
	PortErrorUpdate        = "PORTABLE_ERROR_UPDATE"
	PortErrorCreate        = "PORTABLE_ERROR_CREATE"
	PortErrorDelete        = "PORTABLE_ERROR_DELETE"
	PortErrorLoad          = "PORTABLE_ERROR_LOAD"
	PortErrorSearch        = "PORTABLE_ERROR_SEARCH"
	PortErrorParam         = "PORTABLE_ERROR_PARAM"
	PortErrorValidation    = "PORTABLE_ERROR_VALIDATION"
	PortErrorScript        = "PORTABLE_ERROR_SCRIPT"
	PortErrorDescriptor    = "PORTABLE_ERROR_DESCRIPTOR"
	PortErrorNotification  = "PORTABLE_ERROR_NOTIFICATION"
	PortErrorConsumer      = "PORTABLE_ERROR_CONSUMER"
	PortErrorProducer      = "PORTABLE_ERROR_PRODUCER"
	PortErrorCore          = "PORTABLE_ERROR_CORE"
	PortErrorHandler       = "PORTABLE_ERROR_HANDLER"
	PortErrorAccess        = "PORTABLE_ERROR_ACCESS"
	PortErrorPermission    = "PORTABLE_ERROR_PERMISSION"
	PortErrorAuth          = "PORTABLE_ERROR_AUTH"
	PortErrorMigration     = "PORTABLE_ERROR_MIGRATION"
	PortErrorType          = "PORTABLE_ERROR_TYPE"
	PortErrorFunction      = "PORTABLE_ERROR_FUNCTION"
	PortErrorEncoder       = "PORTABLE_ERROR_ENCODER"
	PortErrorDecoder       = "PORTABLE_ERROR_DECODER"
	PortErrorReader        = "PORTABLE_ERROR_READER"
	PortErrorWriter        = "PORTABLE_ERROR_WRITER"
	PortErrorCloser        = "PORTABLE_ERROR_CLOSER"
	PortErrorTemplate      = "PORTABLE_ERROR_TEMPLATE"
	PortErrorCommand       = "PORTABLE_ERROR_COMMAND"
	PortErrorProcess       = "PORTABLE_ERROR_PROCESS"
	PortErrorIO            = "PORTABLE_ERROR_IO"
	PortErrorRender        = "PORTABLE_ERROR_RENDER"
	PortErrorState         = "PORTABLE_ERROR_STATE"
	PortErrorMemory        = "PORTABLE_ERROR_MEMORY"
	PortErrorNetwork       = "PORTABLE_ERROR_NETWORK"
	PortErrorDevice        = "PORTABLE_ERROR_DEVICE"
	PortErrorRecursion     = "PORTABLE_ERROR_RECURSION"
	PortErrorArgument      = "PORTABLE_ERROR_ARGUMENT"
	PortErrorBody          = "PORTABLE_ERROR_BODY"
	PortErrorHeader        = "PORTABLE_ERROR_HEADER"
	PortErrorProtocol      = "PORTABLE_ERROR_PROTOCOL"
	PortErrorSize          = "PORTABLE_ERROR_SIZE"
	PortErrorPriority      = "PORTABLE_ERROR_PRIORITY"
	PortErrorCondition     = "PORTABLE_ERROR_CONDITION"
	PortErrorIteration     = "PORTABLE_ERROR_ITERATION"
	PortErrorParser        = "PORTABLE_ERROR_PARSER"
	PortErrorEnvelop       = "PORTABLE_ERROR_ENVELOP"
	PortErrorEnvironment   = "PORTABLE_ERROR_ENVIRONMENT"
)

type errorDataJson struct {
	Code    interface{} `json:"code"`
	Name    string      `json:"name,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Portable error
type portErrorJson struct {
	errorDataJson                 // Error data
	Data          []errorDataJson `json:"data,omitempty"`
}

// MarshalJSON Std error marshal
func (e *PortError) MarshalJSON() ([]byte, error) {
	var errors = make([]errorDataJson, len(e.details))
	for i := range e.details {
		errors[i].Code = e.details[i].Code
		errors[i].Name = e.details[i].Name
		errors[i].Message = e.details[i].Message
	}
	obj := &portErrorJson{errorDataJson: errorDataJson(e.ErrorData), Data: errors}
	return json.Marshal(obj)
}

// UnmarshalJSON Unmarshal error
func (e *PortError) UnmarshalJSON(data []byte) error {
	obj := &portErrorJson{}
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}
	//  Convert to origin error
	e.ErrorData = ErrorData(obj.errorDataJson)
	e.details = make([]ErrorData, len(obj.Data))
	for i := range obj.Data {
		e.details[i].Code = obj.Data[i].Code
		e.details[i].Name = obj.Data[i].Name
		e.details[i].Message = obj.Data[i].Message
	}
	return nil
}
