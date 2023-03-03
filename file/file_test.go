package fileutil_test

import (
	fileutil "github.com/yangyang5214/gou/file"
	"testing"
)

func TestFileReadLines(t *testing.T) {
	r := fileutil.FileReadLines("/etc/hosts")
	t.Log(r)
	t.Log(len(*r))
}

func TestCountLines(t *testing.T) {
	r, err := fileutil.CountLines("/etc/hosts")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}
