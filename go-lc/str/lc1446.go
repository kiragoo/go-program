package str

import "math"

func LC1446(s string) int {
	compare := func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}
	if len(s) == 1 {
		return 1
	}

	max := math.MinInt
	l, r := 0, 1
	for r < len(s) {
		if s[r] != s[l] {
			max = compare(max, r-l)
			l = r
		} else {
			r++
		}
	}
	return compare(max, r-l)
}
