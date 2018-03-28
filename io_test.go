package goutils

import (
	"testing"
	"os"
	"path"
	"time"
)

func TestFileExists(t *testing.T) {
	p := path.Join(os.TempDir(), "abc")
	err := os.MkdirAll(p, os.ModeDir)
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(p)

	if !IsDir(p) {
		t.Fatalf("Wanted: true; Actual: %v", false)
	}

	if !FileExists(p) {
		t.Fatalf("Wanted: true; Actual: %v", false)
	}

	if IsFile(p) {
		t.Fatalf("Wanted: false; Actual: %v", true)
	}

	if FileExists("/tmp/" + time.Now().String()) {
		t.Fatalf("Wanted: false; Actual: %v", true)
	}
}

func TestMkDir(t *testing.T) {
	p := path.Join(os.TempDir(), "mk_test_dir")
	err := os.MkdirAll(p, os.ModeDir)
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(p)

	// exists
	if err = MkDir(p); err != nil {
		t.Error(err)
	}
}
