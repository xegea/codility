package countingelements

func Solution__(A []int) int {
	var max int
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if _, ok := m[A[i]]; ok {
			return 0
		}
		m[A[i]]++
		if A[i] > max {
			max = A[i]
		}
	}
	if len(m) != max {
		return 0
	}
	return 1
}
