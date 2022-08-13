package iterations

import "testing"

type binaryGapTest struct {
	arg1     int
	expected int
}

var binaryGapTests = []binaryGapTest{
	{32, 0},
	{529, 4},
	{1041, 5},
	{15, 0},
	{74901729, 4},   // fail
	{1376796946, 5}, // fail
}

func TestSolution(t *testing.T) {
	for _, test := range binaryGapTests {
		if output := Solution(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
