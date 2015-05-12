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
		S:  []string{"S1", "S2"},
		S2: [][]string{{"S11", "S12"}, {"S21", "S22"}},
		SS: []*TestMsg{
			&TestMsg{A: "a1", C: "c1"},
		},
		SSS: []TestMsg{
			TestMsg{A: "a1", C: "c1"},
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
