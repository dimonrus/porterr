package porterr

import (
	"math"
	"reflect"
	"strconv"
	"strings"
)

// IsRequiredValid Required validation rule
func IsRequiredValid(val reflect.Value, args ...string) bool {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return false
		}
		return !val.Elem().IsZero()
	}
	return !val.IsZero()
}

// IsEnumValid In list validation rule
func IsEnumValid(val reflect.Value, args ...string) bool {
	if len(args) == 0 {
		return true
	}
	values := strings.Split(args[0], ",")
	if len(values) == 0 {
		return true
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return false
		}
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.String:
		v := val.String()
		for _, value := range values {
			if v == value {
				return true
			}
		}
		return false
	case reflect.Float64:
		fallthrough
	case reflect.Float32:
		v := val.Float()
		for _, value := range values {
			comp, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return false
			}
			if math.Abs(v - comp) < 1e-7 {
				return true
			}
		}
		return false
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		v := val.Int()
		for _, value := range values {
			comp, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return false
			}
			if v == comp {
				return true
			}
		}
		return false
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		v := val.Uint()
		for _, value := range values {
			comp, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return false
			}
			if v == comp {
				return true
			}
		}
		return false
	}
	return true
}
