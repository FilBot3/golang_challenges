package specific_person

import (
    "testing"
)

func TestSay_Hello_To_Specific_Person(t *testing.T) {
    var name_list = []struct{
        input string
        expected string
    }{
        {"Phillip", "Hello Phillip"},
        {"Phil", "Hello Phil"},
        {"David", "You're not authorized to use this program."},
    }
    
    for _, tt := range name_list {
        result := Say_Hello_To_Specific_Person(tt.input)
        if result != tt.expected {
            t.Errorf("Print_Name() took %s, expected %s, but instead got %s", tt.input, tt.expected, result)
        }
    }
}