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
	if e.GetStack() == nil {
		t.Fatal("no stack")
	}
	r := reflect.TypeOf(e)
	if r.Elem().Name() != "PortError" {
		t.Fatal("Type is wrong")
	}
	if e.(*PortError).Code != "SOME_CODE" {
		t.Fatal("code is wrong")
	}
}

func TestPortError_PushDetail(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("New detail", "SOME_CODE", "item")
	e = e.PushDetail("New detail 2", http.StatusBadRequest, "item second")
	if len(e.GetDetails()) != 2 {
		t.Fatal("wrong detail count")
	}
}

func TestPortError_PopDetail(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("New detail", "SOME_CODE", "item")
	e = e.PushDetail("New detail 2", http.StatusBadRequest, "item second")
	er := e.PopDetail()
	if er.Error() != "New detail 2" {
		t.Fatal("wrong detail message")
	}
	if er.(*PortError).Code != http.StatusBadRequest {
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
	e = e.PushDetail("New detail", "SOME_CODE", "item")
	e = e.PushDetail("New detail 2", http.StatusBadRequest, "item second")
	if len(e.GetDetails()) != 2 {
		t.Fatal("Wrong detail count")
	}
}

func TestPortError_FlushDetails(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("New detail", "SOME_CODE", "item")
	e = e.PushDetail("New detail 2", http.StatusBadRequest, "item second")
	e = e.FlushDetails()
	if len(e.GetDetails()) != 0 {
		t.Fatal("Flush is not works")
	}
}

func TestPortError_MarshalJSON(t *testing.T) {
	message := "Filed with message"
	e := New(http.StatusInternalServerError, message)
	e = e.PushDetail("New detail", "SOME_CODE", "item")
	e = e.PushDetail("", http.StatusBadRequest, "item second")

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatal("Marshal error")
	}

	if fmt.Sprintf("%s", data) != `{"message":"Filed with message","code":500,"error":[{"message":"New detail","code":"SOME_CODE","name":"item"},{"code":400,"name":"item second"}]}` {
		t.Fatal("wrong marshal")
	}
}

func TestPortError_UnmarshalJSON(t *testing.T) {
	data := []byte(`{"message":"Filed with message","code":500,"name":"Unknown","error":[{"message":"New detail","code":"SOME_CODE","name":"item"},{"message":"New detail 2","code":400,"name":"item second"}]}`)
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
	if e.GetStack() == nil {
		t.Fatal("no stack")
	}
	r := reflect.TypeOf(e)
	if r.Elem().Name() != "PortError" {
		t.Fatal("Type is wrong")
	}
	if e.(*PortError).Code != "SOME_CODE" {
		t.Fatal("code is wrong")
	}
	if e.(*PortError).Name != "Unknown" {
		t.Fatal("name is wrong")
	}
}
