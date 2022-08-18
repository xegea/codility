// MissingInteger
// Converting array to map improves performance
// Score 66%: 	https://app.codility.com/demo/results/demoQGKZA2-7ZE/
// Score 100%: 	https://app.codility.com/demo/results/demo3U7KP4-MTK/
package demo

func Solution(A []int) int {
	m1 := make(map[int]int)
	for i := 0; i < len(A); i++ {
		m1[A[i]] = A[i]
	}

	i := 1
	for ; i <= len(A); i++ {
		if _, ok := m1[i]; !ok {
			return i
		}
	}
	return i
}
