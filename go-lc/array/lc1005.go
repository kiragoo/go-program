package array

import "sort"


func LC1005(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	for i, v := range nums {
		if i < k {
			ans += -1 * v
		} else {
			ans += v
		}
	}
	return ans
}

