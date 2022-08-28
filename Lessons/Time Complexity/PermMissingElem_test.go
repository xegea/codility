package timecomplex

import "testing"

type PermMissingElemTest struct {
	arg1     []int
	expected int
}

var PermMissingElemTests = []PermMissingElemTest{
	{[]int{1, 2, 3, 5}, 4},
}

func TestSolution_(t *testing.T) {
	for _, test := range PermMissingElemTests {
		if output := Solution_(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
