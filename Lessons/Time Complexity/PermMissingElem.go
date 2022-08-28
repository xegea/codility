package timecomplex

func Solution_(A []int) int {
	m := make(map[int]int)
	for _, v := range A {
		m[v] = v
	}

	for i := 1; i <= 100000; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return 100001
}
