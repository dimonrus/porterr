# PORTABLE ERROR
- Easy to say that your app has an error.
- Easy to say to your client that app has an error
- Easy to define any error in your logic
- Classic error representation

## How to use
Simple system error
```
e := porterr.New(PortErrorSystem, message)
```
Simple system error with http response code
```
e := porterr.New(PortErrorSystem, message).HTTP(http.StatusInternalServerError)
```
Error with details
```
e := porterr.New(PortErrorSystem, message)
e = e.PushDetail("DETAIL_CODE", "item", "New detail")
e = e.PushDetail("COMMAND_CODE", "command", "Is required")
e = e.PushDetail("COMMAND_CODE", "command", "Some other message")
```
Easy http validation error
```
e := porterr.HttpValidationError()
```
Validation http error with custom code and message
```
e := porterr.HttpValidationError("CUSTOM_CODE", "your form is invalid")
```

## How to use with validation rules
Form validation use combination of validation rules.
Possible rules as parts of 'valid' tag:
- required. Filed is required
- exp. Regular expression
- range. Range if values
- enum. Predefined enum
- min. Minimum value or length
- max. Maximum value or length
- digit. Only digits in value. Can specify length

Example: `valid:"required;exp~[0-5]+;range~1:50;enum~[5,10,15,20,25];digit~4,10;min~3;max~10"`
```
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
v := TestValidationStruct{Complex: &ComplexStruct{}}
e := porterr.ValidateStruct(&v)
if e != nil {
   panic(e)
}
```

#### If you find this project useful or want to support the author, you can send tokens to any of these wallets
- Bitcoin: bc1qgx5c3n7q26qv0tngculjz0g78u6mzavy2vg3tf
- Ethereum: 0x62812cb089E0df31347ca32A1610019537bbFe0D
- Dogecoin: DET7fbNzZftp4sGRrBehfVRoi97RiPKajV
