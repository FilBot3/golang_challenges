package helloWorld

import (
    "testing"
)

func TestPrintHello(t *testing.T) {
    var string_tests = []struct{
        input string // Input to our test
        expected string // Exected output of the test.
    }{
        {"", "Hello World"},
    }
    
    for _, tt := range string_tests {
        result := PrintHello()
        if result != tt.expected {
            t.Errorf("Print_Hello() did not print expected %s, instead printed %s", tt.expected, result)
        }
    }
}