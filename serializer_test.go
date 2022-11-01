package porterr

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestPortError_Pack(t *testing.T) {
	t.Run("size", func(t *testing.T) {
		bt := bytes.NewBuffer(nil)
		e := NewWithName(uint8(121), "foobar", "some")
		e = e.PushDetail("det_code", "some", "done")
		e = e.PushDetail("det_code2", "some2", "done2")
		e.Origin().Pack(bt)
		t.Log(bt.String())
		bt.Reset()
		e = New(1, "foobar")
		e.Origin().Pack(bt)
		t.Log(bt.String())
	})

	t.Run("unpack", func(t *testing.T) {
		e := NewWithName(uint8(109), "name", "message").HTTP(http.StatusTooManyRequests)
		e = e.PushDetail("one_code", "one_name", "one_message")
		e = e.PushDetail("two_code", "two_name", "two_message")

		bt := bytes.NewBuffer(nil)
		e.Origin().Pack(bt)
		t.Log(string(bt.Bytes()))
		t.Log(bt.Len())
		t.Log(packDetailLen(bt.Bytes()))

		e = UnPack(bt.Bytes())
		data, err := json.Marshal(e)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(data))
		if e.GetHTTP() != http.StatusTooManyRequests {
			t.Fatal("wrong http code")
		}
	})
}

func BenchmarkUnPack(b *testing.B) {
	e := NewWithName(uint8(109), "name", "message").HTTP(http.StatusTooManyRequests)
	e = e.PushDetail("one_code", "one_name", "one_message")
	e = e.PushDetail("two_code", "two_name", "two_message")
	e = e.PushDetail("three_code", "three_name", "three_message")
	e = e.PushDetail("four_code", "four_name", "four_message")
	bt := bytes.NewBuffer(nil)
	e.Origin().Pack(bt)
	for i := 0; i < b.N; i++ {
		_ = UnPack(bt.Bytes())
	}
	b.ReportAllocs()
}

func BenchmarkPackSize(b *testing.B) {
	e := New(1, "foobar")
	e = e.PushDetail("code", "name", "message")
	bt := bytes.NewBuffer(nil)
	for i := 0; i < b.N; i++ {
		e.Origin().Pack(bt)
	}
	b.ReportAllocs()
}
