// https://app.codility.com/demo/results/trainingJ22URD-VDB/ 69%
package timecomplex

func Solution__(A []int) int {
	min := 100000

	for i := 0; i < len(A)-1; i++ {
		var left []int = A[0 : i+1]
		var right []int = A[i+1:]

		eq := abs(sum(left[:]...) - sum(right[:]...))
		if eq < min {
			min = eq
		}
	}
	return min
}

func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
