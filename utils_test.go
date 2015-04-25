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
