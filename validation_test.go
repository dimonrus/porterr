package porterr

import (
	"reflect"
	"testing"
)

type ComplexStruct struct {
	Cool bool
}

type AliasOnTypeString string

type Nested struct {
	Foo int32 `json:"foo"`
	Bar *bool `json:"bar" valid:"required"`
}

type TestValidationStruct struct {
	Name      string            `json:"name" valid:"required;exp~[a-z]+"`
	Number    int               `json:"number" valid:"exp~[0-5]+;range~1:50;enum~[5,10,15,20,25]"`
	IsTrue    *bool             `json:"isTrue"`
	Complex   *ComplexStruct    `json:"complex" valid:"required"`
	Sl        []int64           `json:"sl"`
	SuperName AliasOnTypeString `json:"superName" valid:"required"`
	Nested    Nested            `json:"nested"`
}

func TestValidateStructPointer(t *testing.T) {
	v := TestValidationStruct{Complex: &ComplexStruct{}}
	e := ValidateStruct(&v)
	if e != nil {
		for _, iError := range e.GetDetails() {
			t.Log(iError.Error())
		}
		t.Log(e.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	n := Nested{Bar: new(bool)}
	e := ValidateStruct(n)
	if e == nil {
		t.Fatal("must have en error")
	}
	t.Log(e.Error())
}

//switch val.Kind() {
//	case reflect.Ptr:
//		if val.IsNil() || val.IsZero() {
//			return false
//		}
//		return IsRequiredValid(val.Elem())
//	case reflect.String:
//		return val.String() != ""
//	case reflect.Float64:
//		fallthrough
//	case reflect.Float32:
//		return val.Float() != 0
//	case reflect.Int:
//		fallthrough
//	case reflect.Int8:
//		fallthrough
//	case reflect.Int16:
//		fallthrough
//	case reflect.Int32:
//		fallthrough
//	case reflect.Int64:
//		return val.Int() != 0
//	case reflect.Uint:
//		fallthrough
//	case reflect.Uint8:
//		fallthrough
//	case reflect.Uint16:
//		fallthrough
//	case reflect.Uint32:
//		fallthrough
//	case reflect.Uint64:
//		return val.Uint() != 0
//	case reflect.Slice:
//		fallthrough
//	case reflect.Array:
//		return val.Len() > 0
//	case reflect.Chan:
//		return val.Len() > 0
//	case reflect.Struct:
//		return !val.IsZero()
//	}
//	return true

func TestValidateStruct(t *testing.T) {
	v := TestValidationStruct{Name: "foo", Complex: &ComplexStruct{}}
	e := ValidateStruct(v)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.(*PortError)
			t.Log(er.Name, er.Message)
		}
		t.Log(e.Error())
	}
}

func TestValidationRequired(t *testing.T) {
	var val int32 = 12
	v := reflect.ValueOf(val)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	val = 0
	v = reflect.ValueOf(val)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var pval = new(int32)
	*pval = 12
	v = reflect.ValueOf(pval)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	pval = new(int32)
	v = reflect.ValueOf(pval)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var val32 uint32 = 12
	v = reflect.ValueOf(val32)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	val32 = 0
	v = reflect.ValueOf(val32)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var pval32 = new(uint32)
	*pval32 = 12
	v = reflect.ValueOf(pval32)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	pval32 = new(uint32)
	v = reflect.ValueOf(pval32)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var str string
	v = reflect.ValueOf(str)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var pstr *string
	v = reflect.ValueOf(pstr)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	str = "okey"
	v = reflect.ValueOf(str)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	pstr = new(string)
	*pstr = "okey"
	v = reflect.ValueOf(pstr)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var boolean bool
	v = reflect.ValueOf(boolean)
	if IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	boolean = true
	v = reflect.ValueOf(boolean)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var pboolean *bool
	v = reflect.ValueOf(pboolean)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var fl float32
	v = reflect.ValueOf(fl)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	fl = 0.33
	v = reflect.ValueOf(fl)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var pfl *float32
	v = reflect.ValueOf(pfl)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	pfl = new(float32)
	v = reflect.ValueOf(pfl)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	*pfl = 12.1231
	v = reflect.ValueOf(pfl)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}

	var array []int
	v = reflect.ValueOf(array)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	marray := [3]int{1, 2, 3}
	v = reflect.ValueOf(marray)
	if !IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	mnarray := [0]int{}
	v = reflect.ValueOf(mnarray)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var ch chan int
	v = reflect.ValueOf(ch)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var ch1 = make(chan int, 0)
	v = reflect.ValueOf(ch1)
	if !IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	var ch2 = make(chan int, 1)
	ch2 <- 2
	v = reflect.ValueOf(ch2)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}
	close(ch2)

	var st = struct{ a int }{}
	v = reflect.ValueOf(st)
	if IsRequiredValid(v) {
		t.Fatal("Must be invalid")
	}

	st = struct{ a int }{a: 1}
	v = reflect.ValueOf(st)
	if !IsRequiredValid(v) {
		t.Fatal("Must be valid")
	}
}

type TestEnumStruct struct {
	Foo     string  `json:"foo" valid:"enum~empty,base,value"`
	Number  float32 `json:"number" valid:"enum~0.1,0.5,0.9"`
	Bar     int64   `json:"bar" valid:"enum~200,500,9000,100"`
	PNumber *int64  `json:"pNumber" valid:"enum~100,50,20,10"`
}

func TestIsEnumValid(t *testing.T) {
	s := TestEnumStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	*s.PNumber = 1001
	e := ValidateStruct(s)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.(*PortError)
			t.Log(er.Name, er.Message)
		}
	} else {
		t.Fatal("must be an error")
	}
}

func BenchmarkEnumStruct(b *testing.B) {
	s := TestEnumStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	*s.PNumber = 1001
	for i := 0; i < b.N; i++ {
		e := ValidateStruct(s)
		_ = e
	}
	b.ReportAllocs()
}

type TestRequiredStruct struct {
	Foo     string  `json:"foo" valid:"required"`
	Number  float32 `json:"number" valid:"required"`
	Bar     int64   `json:"bar" valid:"required"`
	PNumber *int64  `json:"pNumber" valid:"required"`
}

func BenchmarkRequired(b *testing.B) {
	s := TestRequiredStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	*s.PNumber = 100
	for i := 0; i < b.N; i++ {
		e := ValidateStruct(s)
		_ = e
	}
	b.ReportAllocs()
}

func BenchmarkCheckNative(b *testing.B) {
	s := TestRequiredStruct{Foo: "vad", Number: 0.19, Bar: 2100, PNumber: new(int64)}
	for i := 0; i < b.N; i++ {
		e := ValidateStruct(s)
		_ = e
	}
	b.ReportAllocs()
}

func TestParseValidTag(t *testing.T) {
	rules := ParseValidTag("exp~[0-5]+;range~1:50;enum~5,10,15,20,25")
	t.Log(rules)
}

func BenchmarkParseValidTag(b *testing.B) {
	rules := ParseValidTag("exp~[0-5]+;range~1:50;enum~5,10,15,20,25;other~cool;cool~231231")
	b.Log(rules)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseValidTag("exp~[0-5]+;range~1:50;enum~5,10,15,20,25")
	}
	b.ReportAllocs()
}

type TestRangeValidation struct {
	RangeInt     int     `json:"rangeInt" valid:"range~-1:50"`
	RangeFloat32 float32 `json:"rangeFloat32" valid:"range~9.5:10.5"`
	RangeUint    uint    `json:"rangeUint" valid:"range~0:10"`
}

func TestIsRangeValid(t *testing.T) {
	s := TestRangeValidation{
		RangeInt:     -1,
		RangeFloat32: 9.8,
		RangeUint:    0,
	}
	e := ValidateStruct(s)
	if e != nil {
		for _, iError := range e.GetDetails() {
			er := iError.(*PortError)
			t.Log(er.Name, er.Message)
		}
		t.Fatal(e)
	}
}

func TestIsRangeInValid(t *testing.T) {
	s := TestRangeValidation{
		RangeInt:     70,
		RangeFloat32: 9,
		RangeUint:    12,
	}
	e := ValidateStruct(s)
	if e == nil {
		t.Fatal("Must be an error")
	} else {
		for _, iError := range e.GetDetails() {
			er := iError.(*PortError)
			t.Log(er.Name, er.Message)
		}
	}
}

func BenchmarkIsRangeValid(b *testing.B) {
	s := TestRangeValidation{
		RangeInt:     22,
		RangeFloat32: 9.8,
		RangeUint:    0,
	}
	for i := 0; i < b.N; i++ {
		_ = ValidateStruct(s)
	}
	b.ReportAllocs()
}

func BenchmarkSeveralTag(b *testing.B) {
	s := TestRangeValidation{
		RangeInt:     22,
		RangeFloat32: 9.8,
		RangeUint:    0,
	}
	t := reflect.TypeOf(s)
	var tag reflect.StructTag
	for i := 0; i < b.N; i++ {
		for j := 0; j < t.NumField(); j++ {
			tag = t.Field(j).Tag
			_ = tag.Get("valid")
			_ = tag.Get("json")
		}
	}
	b.ReportAllocs()
}

type TestReg struct {
	Name string `json:"name" valid:"rx~^\\d+$;rx~[0-8]+"`
	//Name string `json:"name"`
}

func TestIsRegularValid(t *testing.T) {
	rs := TestReg{
		Name: "12291",
	}
	e := ValidateStruct(rs)
	if e != nil {
		t.Fatal(e.GetDetails())
	}

	v := "1"
	e = ValidateStruct(v)
	if e == nil {
		t.Fatal("must be error& wrong type")
	}
}

func BenchmarkIsRegularValid(b *testing.B) {
	rs := TestReg{
		Name: "1221",
	}
	for i := 0; i < b.N; i++ {
		_ = ValidateStruct(rs)
	}
	b.ReportAllocs()
}

type TestValidStruct struct {
	RangeInt *int `json:"rangeInt" valid:"required"`
}

type WrapStruct struct {
	*TestValidStruct
}

func TestVTestValidStruct(t *testing.T) {
	tag := "required"
	rules := ParseValidTag(tag)
	for _, rule := range rules {
		t.Log("name:", rule.Name)
		t.Log("arg:", rule.Args)
		t.Log("####")
	}
}

func BenchmarkRequiredTag(b *testing.B) {
	tag := "required;"
	for i := 0; i < b.N; i++ {
		_ = ParseValidTag(tag)
	}
	b.ReportAllocs()
}
