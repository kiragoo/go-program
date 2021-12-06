package bits

func LC136(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
