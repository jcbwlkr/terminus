package terminus

import "testing"

func TestMaxLineLength(t *testing.T) {
	var s string

	// Simple cases
	s = ""
	if l := maxLineLength(s); l != 0 {
		t.Errorf("Expected 0, got %d", l)
	}
	s = "c"
	if l := maxLineLength(s); l != 1 {
		t.Errorf("Expected 1, got %d", l)
	}
	s = "def"
	if l := maxLineLength(s); l != 3 {
		t.Errorf("Expected 3, got %d", l)
	}

	// Multiple lines
	s = "rick\ndaryl"
	if l := maxLineLength(s); l != 5 {
		t.Errorf("Expected 5, got %d", l)
	}
}
