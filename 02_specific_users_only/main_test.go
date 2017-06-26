package specificPerson

import (
	"testing"
)

func TestSayHelloToSpecificPerson(t *testing.T) {
	var name_list = []struct {
		input    string
		expected string
	}{
		{"Phillip", "Hello Phillip"},
		{"Phil", "Hello Phil"},
		{"David", "You're not authorized to use this program."},
	}

	for _, tt := range name_list {
		result := SayHelloToSpecificPerson(tt.input)
		if result != tt.expected {
			t.Errorf("SayHelloToSpecificPerson() took %s, expected %s, but instead got %s", tt.input, tt.expected, result)
		}
	}
}
