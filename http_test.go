package goutils

import (
	"reflect"
	"testing"
)

type TestMsg struct {
	A string `key:"a"`
	C string `key:"c"`
	E string `key:"e"`
}

func TestToObject(t *testing.T) {
	query := "/a=b&c=d&e=f"
	p2obj := &TestMsg{}
	ToObject(query, reflect.ValueOf(p2obj))
	if nil == p2obj {
		t.Fatal("Nil is not wanted!")
	}
	t.Logf("%v", p2obj)
}
