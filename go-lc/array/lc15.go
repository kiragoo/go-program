package array

import (
	"sort"
)

func LC15(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	if length < 3 {
		return res
	}
	r := length-1
	for f := 0; f <= len(nums)-3; f++ {
		if f > 0 && nums[f-1] == nums[f] {
			continue
		}
		for s := f + 1; s < r; {
			if s > f+1 && nums[s-1] == nums[s]{
				s++
				continue
			}
			sum := nums[f] + nums[s] + nums[r]
			if sum == 0 {
				tmp := []int{nums[f], nums[s], nums[r]}
				res = append(res, tmp)
				s++
			} else if sum > 0 {
				r--
			} else {
				s++
			}

		}
	}
	return res
}
