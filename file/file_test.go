package fileutil

import (
	"testing"
)

func TestFileReadLines(t *testing.T) {
	r := FileReadLines("/etc/hosts")
	t.Log(r)
	t.Log(len(r))
}

func TestCountLines(t *testing.T) {
	r, err := CountLines("/etc/hosts")
	if err != nil {
		t.Error(err)
	}
	t.Log(r)
}
