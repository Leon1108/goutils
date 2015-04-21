package goutils

import "testing"

func TestSearchFloor(t *testing.T) {
	is := []int64{1, 3, 4, 6, 8, 11, 39, 210}
	val, idx, err := SearchFloor(5, is)
	if val != 4 || idx != 2 || err != nil {
		t.Fatalf("Wanted: val = 4, idx = 2; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, _ = SearchFloor(8, is)
	if val != 8 || idx != 4 {
		t.Fatalf("Wanted: val = 8, idx = 4; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, err = SearchFloor(0, is)
	t.Logf("%v", err)
	if val != -1 || idx != -1 || err == nil {
		t.Fatalf("Wanted: val = -1, idx = -1; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, _ = SearchFloor(999, is)
	if val != 210 || idx != 7 {
		t.Fatalf("Wanted: val = 210, idx = 7; Actual: val = %v, idx = %v", val, idx)
	}
}
