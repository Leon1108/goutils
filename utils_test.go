package goutils

import "testing"

func TestToString(t *testing.T) {
	obj := MixedMsg{
		X: "x",
		Z: "z",
		TestMsg: &TestMsg{
			A: "a",
			C: "c",
		},
	}

	t.Logf("%v", ToString(obj))
}

func TestToStringWithZeroValue(t *testing.T) {
	obj := MixedMsg{}
	t.Logf(">>> %v", ToString(obj))
}

func TestToStringWithNil(t *testing.T) {
	obj := MixedMsg{TestMsg: nil}
	t.Logf(">>> %v", ToString(obj))
}
