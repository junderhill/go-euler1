package euler1

import (
	"testing"
)

func BenchmarkSumOfNaturals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumNaturalBelow(1000)
	}
}

func TestSumNaturalBelow(t *testing.T) {
	result := SumNaturalBelow(10)
	if result != 23 {
		t.Error("Expected 23, got", result)
	}
}

type multipleOf3or5TestCase struct {
	in  int
	out bool
}

var testCases = []multipleOf3or5TestCase{
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
