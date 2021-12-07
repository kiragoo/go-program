package array

import "sort"

func LC15(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	if n < 3 {
		return res
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		r := n - 1
		for j := i + 1; j < r; {
			if j > i+1 && nums[j-1] == nums[j] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[r]
			if sum == 0 {
				tmp := []int{nums[i], nums[j], nums[r]}
				res = append(res, tmp)
				j++
			} else if sum < 0 {
				j++
			} else {
				r--
			}
		}
	}
	return res
}
