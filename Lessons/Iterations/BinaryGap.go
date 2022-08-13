package iterations

import (
	"strconv"
)

func Solution(N int) int {
	n := int64(N)
	b := strconv.FormatInt(n, 2)

	count := -1
	bg := 0
	for i := len(b) - 1; i >= 0; i-- {
		if string(b[i]) == "0" && count < 0 {
			continue
		}
		if string(b[i]) == "1" && count >= 0 {
			if count > bg {
				bg = count
			}
			count = -1
		}
		if string(b[i]) == "0" || count == -1 {
			count++
		}
	}
	return bg
}
