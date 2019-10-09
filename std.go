package porterr

import "encoding/json"

type errorDataJson struct {
	Message string      `json:"message,omitempty"`
	Code    interface{} `json:"code"`
	Name    string      `json:"name,omitempty"`
	stack   []byte      `json:"-"`
}

// Portable error
type portErrorJson struct {
	errorDataJson                 // Error data
	Error         []errorDataJson `json:"error,omitempty"`
}

// Std error marshal
func (e *PortError) MarshalJSON() ([]byte, error) {
	// Prepare errors
	var errors []errorDataJson
	for i := range e.details {
		errors = append(errors, errorDataJson(e.details[i]))
	}
	// Create object
	obj := &portErrorJson{errorDataJson: errorDataJson(e.ErrorData), Error:errors}

	return json.Marshal(obj)
}

// Unmarshal error
func (e *PortError) UnmarshalJSON(data []byte) error {
	// Create object
	obj := &portErrorJson{}

	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}

	//  Convert to origin error
	e.ErrorData = ErrorData(obj.errorDataJson)
	for i := range obj.Error {
		e.details = append(e.details, ErrorData(obj.Error[i]))
	}

	return nil
}
