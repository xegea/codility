package countingelements

import "testing"

type PermCheckTest struct {
	arg1     []int
	expected int
}

var permCheckTests = []PermCheckTest{
	{[]int{4, 1, 3, 2}, 1},
	{[]int{4, 1, 3}, 0},
}

func TestSolution__(t *testing.T) {
	for _, test := range permCheckTests {
		if output := Solution__(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}

}
