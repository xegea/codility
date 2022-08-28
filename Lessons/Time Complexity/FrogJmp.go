package timecomplex

import "math"

func Solution(X int, Y int, D int) int {
	return int(math.Ceil((float64(Y) - float64(X)) / float64(D)))
}
