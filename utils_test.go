package goutils

import (
	"testing"
	"time"
)

type TestTimeMsg struct {
	Name  string
	Age   int
	Birth time.Time
}

type TestMsgWithPrivate struct {
	Name   string
	age    int
	others *privateField
}

type privateField struct {
	sex     int
	married bool
}

func TestGetCurrentMillisecond(t *testing.T) {
	t.Logf("%v", GetCurrentMillisecond())
}

func TestGetCurrentTime(t *testing.T) {
	t.Logf("%v", GetCurrentTime())
}

func TestToStringNil(t *testing.T) {
	t.Logf("%v", ToString(nil))
	var obj *MixedMsg
	t.Logf("%v", ToString(obj))

	obj = &MixedMsg{
		TestMsg: nil,
	}
	t.Logf("%v", ToString(obj))
	t.Logf("%v", ToString(obj.TestMsg))

	obj = &MixedMsg{
		TestMsg: &TestMsg{},
	}
	t.Logf("%v", ToString(obj))
}

func TestToString(t *testing.T) {
	obj := MixedMsg{
		X: "x",
		Z: "z",
		TestMsg: &TestMsg{
			A: "a",
			C: "c",
		},
		Msg: TestMsg{},
		S:   []string{"S1", "S2"},
		S2:  [][]string{{"S11", "S12"}, {"S21", "S22"}},
		SS: []*TestMsg{
			&TestMsg{A: "a1", C: "c1"},
		},
		SSS: []TestMsg{
			TestMsg{A: "a1", C: "c1"},
		},
		M:   map[string]string{"K1": "V1", "K2": "V2"},
		MM:  map[string]TestMsg{"K1": TestMsg{A: "M1", C: "C1"}},
		MMM: map[string]*TestMsg{"K1": &TestMsg{A: "M1", C: "C1"}},
	}

	t.Logf("%v", ToString(obj))
}

func TestToStringArray(t *testing.T) {
	objs := []*MixedMsg{
		&MixedMsg{X: "x1"},
		&MixedMsg{X: "x2"},
	}
	t.Logf(">>> %v", ToString(objs))
}

func TestToStringMap(t *testing.T) {
	m := map[string]*MixedMsg{
		"k1": &MixedMsg{Z: "z1"},
		"k2": &MixedMsg{Z: "z2"},
	}
	t.Logf(">>> %v", ToString(m))
}

func TestToStringWithZeroValue(t *testing.T) {
	obj := MixedMsg{}
	t.Logf(">>> %v", ToString(obj))
}

func TestToStringWithNil(t *testing.T) {
	obj := MixedMsg{TestMsg: nil}
	t.Logf(">>> %v", ToString(obj))
}

func TestToStringTime(t *testing.T) {
	t.Logf(">>>>> %v", ToString(time.Now()))

	obj := TestTimeMsg{"Leon", 33, time.Now()}
	t.Logf(">>>>> %v", ToString(obj))
}

func TestToStringPrivate(t *testing.T) {
	t.Logf(">>>>> %v", ToString(TestMsgWithPrivate{
		Name: "Leon",
		age:  11,
		others: &privateField{
			sex:     0,
			married: true,
		},
	}))
}

func TestTimeFromUnixMS(t *testing.T) {
	ms := GetCurrentMillisecond()
	t.Log(TimeFromUnixMS(ms))
}
