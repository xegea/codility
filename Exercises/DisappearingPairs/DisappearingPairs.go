// DisappearingPairs
// Reduce a string containing instances of the letters "A", "B" and "C" via the following rule:
// remove one occurrence of "AA", "BB" or "CC".
// Score 50%: 	https://app.codility.com/demo/results/trainingBPFXGJ-976/
// Score 83%:
// Score 100%: 	https://app.codility.com/demo/results/trainingTDS2UA-8A8/
package algorithmicskills

import "strings"

func Solution(S string) string {
	b := []byte(S)
	for i := len(b) - 1; i > 0; i-- {
		if i == len(b) {
			continue
		}
		if b[i] == b[i-1] {
			b = append(b[:i-1], b[i+1:]...)
		}
	}
	return string(b)
}

func Solution2(s string) string {
	// Score 83% Detected time complexity: Detected time complexity: O(N) or O(N ** 2)
	l := len(s)

	s = strings.Join(strings.Split(s, "AA"), "")
	s = strings.Join(strings.Split(s, "BB"), "")
	s = strings.Join(strings.Split(s, "CC"), "")

	if len(s) < l {
		s = Solution2(s)
	}
	return s
}

func Solution1(s string) string {
	// Score 50% Detected time complexity: O(2 ** N)
	for replace := true; replace == true; {
		if strings.Contains(s, "AA") {
			s = strings.Replace(s, "AA", "", 1)
		} else if strings.Contains(s, "BB") {
			s = strings.Replace(s, "BB", "", 1)
		} else if strings.Contains(s, "CC") {
			s = strings.Replace(s, "CC", "", 1)
		} else {
			replace = false
		}
	}
	return s
}
