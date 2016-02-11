package util

import (
	"runtime"
	"testing"
)

func TestGetOsType(t *testing.T) {
	expexted := runtime.GOOS
	actual := GetOsType()

	// hello
	if expexted != actual {
		t.Errorf("%q != %q", actual, expexted)
	}
}
