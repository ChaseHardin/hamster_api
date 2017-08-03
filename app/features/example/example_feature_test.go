package example

import (
	"testing"
)

func TestTrueIsTruelyTrue(t *testing.T) {
	if true != true {
		t.Errorf("Is not true")
	}
}
func TestFalseIsNotTrueTrue(t *testing.T) {
	if false == true {
		t.Errorf("Is not true")
	}
}
