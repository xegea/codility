package timecomplex

import "testing"

type frogJmpTest struct {
	arg1     int
	arg2     int
	arg3     int
	expected int
}

var frogJmpTests = []frogJmpTest{
	{10, 85, 30, 3},
}

func TestSolution(t *testing.T) {
	for _, test := range frogJmpTests {
		if output := Solution(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
