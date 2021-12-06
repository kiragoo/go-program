package array

func LC268(nums []int) int {
	var res int
	tag := make([]int, len(nums)+1)
	for _, num := range nums {
		tag[num] = 1
	}
	for k, v := range tag {
		if v != 1 {
			res = k
		}
	}
	return res
}

func LC268_1(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	for i := 0; i <= len(nums); i++ {
		res ^= i
	}
	return res
}
