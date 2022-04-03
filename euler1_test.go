package euler1

import (
	"testing"
)

func TestSumNaturalBelow(t *testing.T) {
	result := SumNaturalBelow(10)
	if result != 23 {
		t.Error("Expected 23, got", result)
	}
}

var testCases = []struct {
	in  int
	out bool
}{
	{3, true},
	{5, true},
	{4, false},
	{10, true},
}

func TestMultipleOf3or5(t *testing.T) {
	for _, testCase := range testCases {
		result := MultipleOf3Or5(testCase.in)
		if result != testCase.out {
			t.Error("Expected", testCase.out, "got", result)
		}
	}
}
