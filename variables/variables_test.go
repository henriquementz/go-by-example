package variables

import "testing"

func TestInt(t *testing.T) {
	integer := GetInteger()
	expected := 2

	if integer != expected {
		t.Errorf("expected '%q' but got '%q'", expected, integer)
	}

}

func TestBool(t *testing.T) {
	boolean := GetBoolean()
	expected := !false

	if boolean != expected {
		t.Errorf("expected '%t' but got '%t'", expected, boolean)
	}

}
