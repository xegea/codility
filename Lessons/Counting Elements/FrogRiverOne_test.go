package countingelements

import "testing"

type FrogRiverOneTest struct {
	arg1     int
	arg2     []int
	expected int
}

var frogRiverOneTests = []FrogRiverOneTest{
	{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
}

func TestSolution_(t *testing.T) {
	for _, test := range frogRiverOneTests {
		if output := Solution_(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
