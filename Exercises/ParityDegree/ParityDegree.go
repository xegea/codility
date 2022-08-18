package codingskills

func Solution(N int) int {
	for i := 0; i <= 30; i++ {
		p := power(2, i)
		if N%p == 0 {
			continue
		}
		return i - 1
	}
	return 0
}

func power(base int, exp int) int {
	p := 1
	for exp != 0 {
		p *= base
		exp -= 1
	}
	return p
}
