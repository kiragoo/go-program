package array

import (
	"math"
	"sort"
)

func LC16(nums []int, target int) int {
	min := func(x, y float64) int {
		if x >=y {
			return int(y)
		}
		return int(x)
	}
	sort.Ints(nums)
	n := len(nums)
	res := math.MaxInt32
	for i:=0;i<n;i++ {
		r := n-1
		if i>0 && nums[i-1] == nums[i] {
			continue
		}
		for j:=i+1;j<r;j++ {
			if j>i+1 && nums[j-1] == nums[j] {
				j++
				continue
			}
			tmp := math.Abs(float64(target-nums[i]-nums[j]-nums[r]))
			res = min(tmp, float64(res))
		}
	}
	return res
}
