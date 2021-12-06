package bits

func LC287(nums []int) int {
	var res int
	tags := make([]int, len(nums))

	for _, num := range nums {
		if tags[num-1] == 0 {
			tags[num-1] = 1
		} else {
			res = num
			break
		}
	}
	return res
}
