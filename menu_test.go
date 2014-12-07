package terminus

import "testing"

func TestNewExitOption(t *testing.T) {
	option := NewExitOption("Quit")

	if option.Label != "Quit" {
		t.Errorf("Exit option label should be %q, was %q", "Quit", option.Label)
	}
	if ret := option.Action(); ret != Exit {
		t.Error("Action from exit option should return Exit constant")
	}
}
