package arrays

import "testing"

type OddOcurrencesInArrayTest struct {
	arg1     []int
	expected int
}

var oddOcurrencesInArrayTests = []OddOcurrencesInArrayTest{
	{[]int{7, 5, 3, 1, 5, 3, 7}, 1},
	{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
}

func TestSolution_(t *testing.T) {
	for _, test := range oddOcurrencesInArrayTests {
		if output := Solution_(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
