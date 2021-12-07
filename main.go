package main

import "fmt"

var ans = []int{}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// ans := []int{}
	inner(nums, 4)
	fmt.Println(ans)
}

func inner(nums []int, i int) {
	if i < 0 {
		return
	}
	inner(nums, i-1)
	ans = append(ans, nums[i])
}
