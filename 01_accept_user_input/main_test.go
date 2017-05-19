package accept_user_input

import (
    "testing"
)

func TestPrint_Name(t *testing.T) {
    var name_table = []struct{
        input string
        expected string
    }{
        {"Phillip", "Your name is: Phillip"},
        {"Phil", "Your name is: Phil"},
        {"David", "Your name is: David"},
    }
    
    for _, tt := range name_table {
        result := Print_Name(tt.input)
        if result != tt.expected {
            t.Errorf("Print_Name() took %s, expected %s, but instead got %s", tt.input, tt.expected, result)
        }
    }
}