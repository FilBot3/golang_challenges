package hello_world

import (
    "testing"
)

func TestPrint_Hello(t *testing.T) {
    var string_tests = []struct{
        input string // Input to our test
        expected string // Exected output of the test.
    }{
        {"", "Hello World"},
    }
    
    for _, tt := range string_tests {
        result := Print_Hello()
        if result != tt.expected {
            t.Errorf("Print_Hello() did not print expected %s, instead printed %s", tt.expected, result)
        }
    }
}