package timecomplex

import "testing"

type TapeEquilibriumTest struct {
	arg1     []int
	expected int
}

var TapeEquilibriumTests = []TapeEquilibriumTest{
	{[]int{3, 1, 2, 4, 3}, 1},
}

func TestSolution__(t *testing.T) {
	for _, test := range TapeEquilibriumTests {
		if output := Solution__(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
