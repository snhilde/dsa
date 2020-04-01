package htree

import (
	"testing"
)


func TestNew(t *testing.T) {
	tr := New()
	if tr == nil {
		t.Error("Failed to create tree")
	}

	if tr.trunk != nil {
		t.Error("Trunk node is not nil")
	}

	if tr.length != 0 {
		t.Error("Tree claims to have nodes")
	}
}
