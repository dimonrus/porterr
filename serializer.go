package porterr

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	// PackControlSymbolTilda ~ symbol
	PackControlSymbolTilda = byte(126)
	// PackControlSymbolVertical | symbol
	PackControlSymbolVertical = byte(124)
)

var (
	// ErrorPackSign Define firse byte for extended error serial protocol
	ErrorPackSign = []byte("xe")
	// PortErrorPackIdentifier check bytes to confirm that is PortError package
	PortErrorPackIdentifier = []byte("pe")
	// PackErrorPrefix Whole preffix for packed error
	PackErrorPrefix = []byte{120, 101, 58, 112, 101, 58}
)

// Pack error
func (e ErrorData) Pack(buf *bytes.Buffer) {
	switch v := e.Code.(type) {
	case string:
		buf.WriteString(v)
	case int:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case int8:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case int16:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case int32:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case int64:
		buf.WriteString(strconv.FormatInt(v, 10))
	case uint:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case uint8:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case uint16:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case uint32:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	case uint64:
		buf.WriteString(strconv.FormatInt(int64(v), 10))
	}
	if len(e.Name) > 0 {
		buf.WriteString("|" + e.Name)
	}
	if len(e.Message) > 0 {
		buf.WriteString("|" + e.Message)
	}
	return
}

// Pack pack error into []byte
func (e *PortError) Pack(b *bytes.Buffer) {
	// first 6 bytes is 'xe:pe:'
	b.Write(PackErrorPrefix)
	// next 2 byte is length of all error
	b.Write([]byte{0, 0})
	// pack error
	e.ErrorData.Pack(b)
	// pack details
	for _, detail := range e.details {
		b.WriteByte(PackControlSymbolTilda)
		detail.Pack(b)
	}
	// append http code
	b.WriteByte(PackControlSymbolTilda)
	b.WriteByte(byte(e.httpCode >> 8))
	b.WriteByte(byte(e.httpCode))
	// collect length
	data := b.Bytes()
	errorLen := b.Len() - 8
	// set length into data
	data[6] = byte(errorLen >> 8)
	data[7] = byte(errorLen)
	return
}

// get detail length
func packDetailLen(data []byte) int {
	// prefix + 2 bytes for len
	if len(data) < len(PackErrorPrefix)+2 {
		return -1
	}
	// check prefix
	for i := range PackErrorPrefix {
		if data[i] != PackErrorPrefix[i] {
			return -1
		}
	}
	var d = -1
	for i := 0; i < len(data); i++ {
		if data[i] == PackControlSymbolTilda {
			d++
		}
	}
	return d
}

func injectInString(s *string, data []byte) {
	pSlice := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	pString := (*reflect.StringHeader)(unsafe.Pointer(s))
	pString.Data = pSlice.Data
	pString.Len = pSlice.Len
}

// UnPack if bytes contains Port Error
func UnPack(data []byte) IError {
	// get details len
	errorLen := packDetailLen(data)
	if errorLen == -1 {
		return nil
	}
	// get len
	l := binary.BigEndian.Uint16(data[len(PackErrorPrefix) : len(PackErrorPrefix)+2])
	data = data[len(PackErrorPrefix)+2:]
	var e = PortError{details: make([]ErrorData, errorLen)}
	var a, k int
	var d = -1
	var isString bool
	var ed = &e.ErrorData
	for i := uint16(0); i < l; {
		for i < l && data[i] != PackControlSymbolTilda {
			if data[i] == PackControlSymbolVertical {
				if k == 0 {
					if isString {
						var s string
						injectInString(&s, data[a:i])
						ed.Code = s
					} else {
						var s string
						injectInString(&s, data[a:i])
						ed.Code, _ = strconv.Atoi(s)
					}
				} else if k == 1 {
					injectInString(&ed.Name, data[a:i])
				} else if k == 2 {
					injectInString(&ed.Message, data[a:i])
				}
				k++
				a = int(i + 1)
				isString = false
			} else {
				if (data[i] < '0' || data[i] > '9') && !isString {
					isString = true
				}
			}
			i++
		}
		if k == 0 {
			e.httpCode = int(binary.BigEndian.Uint16(data[a:i]))
			break
		} else if k == 1 {
			injectInString(&ed.Name, data[a:i])
		} else if k == 2 {
			injectInString(&ed.Message, data[a:i])
		}
		i++
		d++
		a = int(i)
		k = 0
		if d < errorLen {
			ed = &e.details[d]
		}
	}
	return &e
}
