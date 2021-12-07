package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	var ans []int
	inner(nums, 5, ans)
	fmt.Println(ans)
}

func inner(nums []int, i int, array []int) {
	if i == 0 {
		_ = append(array, nums[0])
		return
	}
	inner(nums, i-1, array)
	_ = append(array, nums[i-1])
}
