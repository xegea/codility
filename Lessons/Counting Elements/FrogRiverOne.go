package countingelements

func Solution_(X int, A []int) int {
	m := make(map[int]int, X)
	for i := 0; i < len(A); i++ {
		if _, ok := m[A[i]]; !ok {
			m[A[i]]++
		}
		if len(m) == X {
			return i
		}
	}
	return -1
}
