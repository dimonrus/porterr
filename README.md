# PORTABLE ERROR
- Easy to say that your app has an error.
- Easy to say to your client that app has an error
- Easy to define any error in your logic
- Classic error representation
- Supports errors.Is for IError type

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

#### If you find this project useful or want to support the author, you can send tokens to any of these wallets
- Bitcoin: bc1qgx5c3n7q26qv0tngculjz0g78u6mzavy2vg3tf
- Ethereum: 0x62812cb089E0df31347ca32A1610019537bbFe0D
- Dogecoin: DET7fbNzZftp4sGRrBehfVRoi97RiPKajV
