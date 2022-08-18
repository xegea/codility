package algorithmicskills

import "testing"

type DisappearingPairsTest struct {
	arg1     string
	expected string
}

var disappearingPairsTests = []DisappearingPairsTest{
	{"ACCAABBC", "AC"},
	{"ABCBBCBA", ""},
	{"BABABA", "BABABA"},
	{"AAABCBAA", "ABCB"},
}

func TestSolution(t *testing.T) {
	for _, test := range disappearingPairsTests {
		if output := Solution(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
