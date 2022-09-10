// https://app.codility.com/demo/results/trainingJ22URD-VDB/ 69%
package timecomplex

import "math"

func Solution__(A []int) int {
	var min float64 = 100000

	for i := 1; i < len(A)-1; i++ {
		eq := math.Abs(float64(sum(A[:i]...) - sum(A[i:]...)))
		if eq < min {
			min = eq
		}
	}
	return int(min)
}

func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
