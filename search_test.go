package goutils

import "testing"

func TestSearchFloor(t *testing.T) {
	is := []int64{1, 3, 4, 6, 8, 11, 39, 210}
	val, idx := SearchFloor(5, is)
	if val != 4 || idx != 2 {
		t.Fatalf("Wanted: val = 4, idx = 2; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx = SearchFloor(8, is)
	if val != 8 || idx != 4 {
		t.Fatalf("Wanted: val = 8, idx = 4; Actual: val = %v, idx = %v", val, idx)
	}
}
