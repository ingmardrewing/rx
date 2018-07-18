package rx

import "testing"

func TestSubstituteAllOccurences(t *testing.T) {
	txt := "Hello World! Hello Waldo!"
	rx, err := NewRx("W[drloa]*!")
	if err != nil {
		t.Error(err)
	}

	expected := "Hello Wurst! Hello Wurst!"
	actual, err := rx.SubstituteAllOccurences(txt, "Wurst!")

	if actual != expected {
		t.Error("Expected `", expected, "` but got `", actual, "`")
	}
}
