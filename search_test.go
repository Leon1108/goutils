package goutils

import (
	"testing"
	"fmt"
)

func TestSearchFloor(t *testing.T) {
	is := []int64{11, 39, 1, 3, 4, 6, 8, 210}
	val, idx, err := SearchFloor(5, is)
	if val != 4 || idx != 2 || err != nil {
		t.Fatalf("Wanted: val = 4, idx = 2; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, _ = SearchFloor(8, is)
	if val != 8 || idx != 4 {
		t.Fatalf("Wanted: val = 8, idx = 4; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, err = SearchFloor(0, is)
	if val != -1 || idx != -1 || err == nil {
		t.Fatalf("Wanted: val = -1, idx = -1; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, _ = SearchFloor(999, is)
	if val != 210 || idx != 7 {
		t.Fatalf("Wanted: val = 210, idx = 7; Actual: val = %v, idx = %v", val, idx)
	}

	val, idx, _ = SearchFloor(999, nil)
	if val != -1 && idx != -1 {
		t.Fatalf("Wanted: val = -1, idx = -1; Actual: val = %v, idx = %v", val, idx)
	}
	val, idx, _ = SearchFloor(999, []int64{})
	if val != -1 && idx != -1 {
		t.Fatalf("Wanted: val = -1, idx = -1; Actual: val = %v, idx = %v", val, idx)
	}
	val, idx, _ = SearchFloor(999, []int64{1000})
	if val != -1 && idx != -1 {
		t.Fatalf("Wanted: val = -1, idx = -1; Actual: val = %v, idx = %v", val, idx)
	}
	val, idx, _ = SearchFloor(999, []int64{998})
	if val != 998 && idx != 0 {
		t.Fatalf("Wanted: val = 998, idx = 0; Actual: val = %v, idx = %v", val, idx)
	}
	val, idx, _ = SearchFloor(999, []int64{998, 999})
	if val != 999 && idx != 1 {
		t.Fatalf("Wanted: val = 999, idx = 1; Actual: val = %v, idx = %v", val, idx)
	}
}

func TestSearchFloorV2(t *testing.T) {
	is := []int64{
		19910816,
		19920322,
		19930523,
		19940710,
		19950924,
		19960526,
		19970824,
		19991017,
		20001105,
		20020722,
		20030928,
		20070619,
		20081030,
		20121018,
		20130619,
		20140611,
		20150412,
		20160615,
		20170720}
	val, idx, err := SearchFloor(19901010, is)
	fmt.Println(val, idx, err)

	val, idx, err = SearchFloor(20150413, is)
	fmt.Println(val, idx, err)

	val, idx, err = SearchFloor(20160616, is)
	fmt.Println(val, idx, err)
	//if val != 4 || idx != 2 || err != nil {
	//	t.Fatalf("Wanted: val = 4, idx = 2; Actual: val = %v, idx = %v", val, idx)
	//}
}
