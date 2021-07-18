package porterr

import (
	"net/http"
	"reflect"
)

// ValidationCallback function that performs validation rule
type ValidationCallback func(val reflect.Value, args ...string) bool

// ValidationRule validation params
type ValidationRule struct {
	// Validator name
	Name string
	// Validator argument
	Args []string
}

// Basic validation rules
// You can override by using var CustomValidationRules
var basicValidationRules = map[string]ValidationCallback{
	// Required validator
	"required": IsRequiredValid,
	// Enum validator
	"enum": IsEnumValid,
}

// Will be used in validation method
var actualValidationRules map[string]ValidationCallback

// PrepareActualValidationRules func to append basicValidationRules or replace some one rules
// customValidationRules If you want to use your own validation rules
// add the rules in to customValidationRules var
func PrepareActualValidationRules(customValidationRules map[string]ValidationCallback) {
	actualValidationRules = make(map[string]ValidationCallback)
	for s, callback := range basicValidationRules {
		actualValidationRules[s] = callback
	}
	for s, callback := range customValidationRules {
		actualValidationRules[s] = callback
	}
}

// ValidateStruct struct fields validation
func ValidateStruct(v interface{}) IError {
	e := New(PortErrorValidation, "Validation error").HTTP(http.StatusBadRequest)
	ve := reflect.ValueOf(v)
	te := reflect.TypeOf(v)

	if ve.Kind() == reflect.Ptr {
		ve = ve.Elem()
		te = te.Elem()
	}

	var fieldName string
	var rules []ValidationRule

	var f reflect.Value
	var t reflect.StructField

	for i := 0; i < ve.NumField(); i++ {
		f = ve.Field(i)
		t = te.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			e = e.MergeDetails(ValidateStruct(f.Interface()))
		case reflect.Ptr:
			if !f.IsNil() {
				if f.Elem().Kind() == reflect.Struct {
					e = e.MergeDetails(ValidateStruct(f.Interface()))
				}
			}
		}
		jsonTag := t.Tag.Get("json")
		if jsonTag != "" {
			fieldName = jsonTag
		} else {
			fieldName = t.Name
		}
		rules = ParseValidTag(t.Tag.Get("valid"))
		for _, rule := range rules {
			if vRule, ok := actualValidationRules[rule.Name]; ok {
				if !vRule(f, rule.Args...) {
					e = e.PushDetail(PortErrorParam, fieldName, "Invalid validation for "+rule.Name+" rule")
				}
			}
		}
	}
	return e.IfDetails()
}

// ParseValidTag parse validation tag for rule and arguments
// Example
// valid:"exp~[0-5]+;range~1-50;enum~5,10,15,20,25"`
func ParseValidTag(validTag string) []ValidationRule {
	if validTag == "" {
		return nil
	}
	var result = make([]ValidationRule, 4)
	if validTag[len(validTag)-1] != ';' {
		validTag += string(';')
	}
	var ruleCount int
	var nameIndexStart, nameIndexEnd int
	var argIndexStart int

	for i, symbol := range validTag {
		if symbol == '~' {
			nameIndexEnd = i
			argIndexStart = i + 1
			if ruleCount == len(result) {
				result = append(result, make([]ValidationRule, 4)...)
			}
			continue
		}
		if symbol == ';' {
			result[ruleCount].Name = validTag[nameIndexStart:nameIndexEnd]
			result[ruleCount].Args = []string{validTag[argIndexStart:i]}
			ruleCount++
			nameIndexStart = i + 1
		}
	}
	return result[:ruleCount]
}

// Init default validators
func init() {
	PrepareActualValidationRules(nil)
}
