package codingskills

import "testing"

type CodingskillsTest struct {
	arg1     int
	expected int
}

var codingskillsTests = []CodingskillsTest{
	{536870912, 29},
	{24, 3},
}

func TestSolution(t *testing.T) {
	for _, test := range codingskillsTests {
		if output := Solution(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
