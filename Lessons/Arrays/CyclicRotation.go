package arrays

func Solution(A []int, K int) []int {
	if len(A) <= 1 {
		return A
	}

	for i := 0; i < K; i++ {
		A = append(A[len(A)-1:], A[:len(A)-1]...)
	}

	return A
}
