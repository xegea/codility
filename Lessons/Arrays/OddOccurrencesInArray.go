package arrays

func Solution_(A []int) int {
	m := make(map[int]int)
	for _, v := range A {
		m[v]++
	}

	for n, o := range m {
		if o%2 != 0 {
			return n
		}
	}
	return 0
}
