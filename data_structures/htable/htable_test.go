package htable

import (
	"testing"
)


func TestNew(t *testing.T) {
	n := New()
	if n == nil {
		t.Error("Failed to create new table")
	}
}
