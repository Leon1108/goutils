package goutils

import (
	"reflect"
	"testing"
)

type SuperMsg struct {
	Leon string `key:"y"`
}

type TestMsg struct {
	SuperMsg
	A string `key:"a"`
	C string `key:"c"`
	E string `key:"e"`
}

type MixedMsg struct {
	*TestMsg
	X string `key:"x"`
	Z string `key:"z"`
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

func TestMixedToObject(t *testing.T) {
	query := "/a=b&c=d&e=f&x=x&z=z&y=y"
	p2obj := &MixedMsg{}
	ToObject(query, reflect.ValueOf(p2obj))
	if nil == p2obj {
		t.Fatal("Nil is not wanted!")
	}
	t.Logf("%v", ToString(p2obj))
}
