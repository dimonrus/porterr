package porterr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	if e.Error() != message {
		t.Fatal("Wrong message")
	}
}

func TestNewF(t *testing.T) {
	message := "Filed with message with %s param and id: %v"
	e := NewF("SOME_CODE", message, "custom", 1)
	if e.Error() != "Filed with message with custom param and id: 1" {
		t.Fatal("Wrong message")
	}
	r := reflect.TypeOf(e)
	if r.Elem().Name() != "PortError" {
		t.Fatal("Type is wrong")
	}
	if e.GetCode() != "SOME_CODE" {
		t.Fatal("code is wrong")
	}
	e = e.Code("New_Code")
	if e.GetCode() != "New_Code" {
		t.Fatal("Code is wrong")
	}
}

func TestPortError_PushDetail(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("SOME_CODE", "item", "New detail")
	e = e.PushDetail(http.StatusBadRequest, "item second", "New detail 2")
	if len(e.GetDetails()) != 2 {
		t.Fatal("wrong detail count")
	}
}

func TestPortError_PopDetail(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("SOME_CODE", "item", "New detail")
	e = e.PushDetail(http.StatusBadRequest, "item second", "New detail 2")
	er := e.PopDetail()
	if er.Error() != "New detail 2" {
		t.Fatal("wrong detail message")
	}
	if er.GetCode() != http.StatusBadRequest {
		t.Fatal("code is not works")
	}
	if len(e.GetDetails()) != 1 {
		t.Fatal("pop does not works")
	}
	er = e.PopDetail()
	er = e.PopDetail()
	if er != nil {
		t.Fatal("wrong pop")
	}
}

func TestPortError_GetDetails(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("SOME_CODE", "item", "New detail")
	e = e.PushDetail(http.StatusBadRequest, "item second", "New detail 2")
	if len(e.GetDetails()) != 2 {
		t.Fatal("Wrong detail count")
	}
}

func TestPortError_FlushDetails(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("SOME_CODE", "item", "New detail")
	e = e.PushDetail(http.StatusBadRequest, "item second", "New detail 2")
	e = e.FlushDetails()
	if len(e.GetDetails()) != 0 {
		t.Fatal("Flush is not works")
	}
}

func TestPortError_MarshalJSON(t *testing.T) {
	message := "Filed with message"
	e := New(PortErrorSystem, message).HTTP(http.StatusInternalServerError)
	e = e.PushDetail("SOME_CODE", "item", "New detail")
	e = e.PushDetail(http.StatusBadRequest, "item second", "")

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal("Marshal error")
	}

	if fmt.Sprintf("%s", data) != `{"message":"Filed with message","code":"PORTABLE_ERROR_SYSTEM","data":[{"message":"New detail","code":"SOME_CODE","name":"item"},{"code":400,"name":"item second"}]}` {
		t.Fatal("wrong marshal")
	}

	if e.GetHTTP() != 500 {
		t.Fatal("wrong http code")
	}
}

func TestPortError_UnmarshalJSON(t *testing.T) {
	data := []byte(`{"message":"Filed with message","code":500,"name":"Unknown","data":[{"message":"New detail","code":"SOME_CODE","name":"item"},{"message":"New detail 2","code":400,"name":"item second"}]}`)
	e := &PortError{}
	err := json.Unmarshal(data, e)
	if err != nil {
		t.Fatal(err)
	}
	if e.Error() != "Filed with message" {
		t.Fatal("Wrong unmarshal")
	}
	if len(e.GetDetails()) != 2 {
		t.Fatal("wrong unmarshal details count")
	}
}

func TestPortError_UnmarshalJSON2(t *testing.T) {
	e := &PortError{}
	data := []byte(`{"message":233}`)
	err := json.Unmarshal(data, e)
	if err == nil {
		t.Fatal("Must be an error")
	}
}

func TestNewWithName(t *testing.T) {
	message := "Filed with message"
	e := NewWithName(http.StatusInternalServerError, "Unknown", message)
	if e.Error() != message {
		t.Fatal("Wrong message")
	}
}

func TestNewFWithName(t *testing.T) {
	message := "Filed with message with %s param and id: %v"
	e := NewFWithName("SOME_CODE", "Unknown", message, "custom", 1)
	if e.Error() != "Filed with message with custom param and id: 1" {
		t.Fatal("Wrong message")
	}
	r := reflect.TypeOf(e)
	if r.Elem().Name() != "PortError" {
		t.Fatal("Type is wrong")
	}
	if e.GetCode() != "SOME_CODE" {
		t.Fatal("code is wrong")
	}
	if e.(*PortError).Name != "Unknown" {
		t.Fatal("name is wrong")
	}
}

func TestPortError_MergeDetails(t *testing.T) {
	data := []byte(`{"message":"Filed with message","code":500,"name":"Unknown","data":[{"message":"New detail","code":"SOME_CODE","name":"item"},{"message":"New detail 2","code":400,"name":"item second"}]}`)
	e := New(PortErrorIO, "Filed with message")
	err := json.Unmarshal(data, e)
	if err != nil {
		t.Fatal(err)
	}
	if e.Error() != "Filed with message" {
		t.Fatal("Wrong unmarshal")
	}
	if len(e.GetDetails()) != 2 {
		t.Fatal("wrong unmarshal details count")
	}

	data = []byte(`{"message":"Filed with message","code":500,"name":"Unknown","data":[{"message":"New detail","code":"SOME_CODE_2","name":"item"}]}`)
	e2 := New(PortErrorIO, "Filed with message")
	err = json.Unmarshal(data, e2)
	if err != nil {
		t.Fatal(err)
	}

	data = []byte(`{"message":"Filed with message","code":500,"name":"Unknown"}`)
	e3 := New(PortErrorIO, "Filed with message")
	err = json.Unmarshal(data, e3)
	if err != nil {
		t.Fatal(err)
	}

	var e4 IError
	e = e.MergeDetails(e2, e3, e4)
	if len(e.GetDetails()) != 3 {
		t.Fatal("must be 3")
	}

	d, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(d))
}

func TestPortError_IfDetails(t *testing.T) {
	e := New(PortErrorSearch, "some test error")
	e = e.PushDetail(PortErrorIO, "name", "try it")

	if e.IfDetails() == nil {
		t.Fatal("if detail isn work properly")
	}

	e = New(PortErrorSearch, "some test error")
	if e.IfDetails() != nil {
		t.Fatal("if detail isn work properly. nil expected")
	}
}

func TestHttpValidationError(t *testing.T) {
	t.Run("classic", func(t *testing.T) {
		e := HttpValidationError()
		if e.Error() != "Validation error" {
			t.Fatal("wrong message")
		}
		if e.GetCode() != PortErrorValidation {
			t.Fatal("wrong code")
		}
		if e.GetHTTP() != http.StatusBadRequest {
			t.Fatal("wrong http code")
		}
	})
	t.Run("conflict", func(t *testing.T) {
		e := HttpValidationError(PortErrorConflict)
		if e.Error() != "Validation error" {
			t.Fatal("wrong message")
		}
		if e.GetCode() != PortErrorConflict {
			t.Fatal("wrong code")
		}
		if e.GetHTTP() != http.StatusBadRequest {
			t.Fatal("wrong http code")
		}
	})
}

func BenchmarkHttpValidationError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HttpValidationError(PortErrorConflict)
	}
	b.ReportAllocs()
}
