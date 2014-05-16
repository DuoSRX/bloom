package bloom

import (
	"testing"
)

func TestSimple(t *testing.T) {
	bf := New(1000)
	bf.Add("whatever")

	if !bf.IsPresent("whatever") {
		t.Error("'whatever' should be present")
	}

	if bf.IsPresent("notfound") {
		t.Error("'notfound' shouldn't be present")
	}
}
