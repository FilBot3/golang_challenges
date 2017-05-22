package sumOfNumbersToN

import (
	"testing"
)

func TestSumOfNumbersToN(t *testing.T) {
	var test_numbers = []struct {
		input    int64
		expected int64
	}{
		{4, 10},
		{5, 15},
		{10, 55},
	}

	for _, tt := range test_numbers {
		result := SumOfNumbersToN(tt.input)
		if result != tt.expected {
			t.Errorf("SumOfNumbersToN had the input of: %s, expected %s, but instead got %s", tt.input, tt.expected, result)
		}
	}
}
