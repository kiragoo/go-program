package array

func LC137(nums []int) int {
	m := map[int]int{}
	var res int
	for _, v := range nums {
		_, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = m[v] + 1
		}
	}
	for k, v := range m {
		if v == 1 {
			res = k
			break
		}
	}
	return res
}

// 2: 10
// 2: 10
// 3: 11
// 2: 10

// num>>i & 01 => $ + total => total
// 01       01 => 01 + 0    => 01
// 01		 01 => 01 + 01 	 => 10
// 11		 01 => 01 + 10 	 => 11

func LC137_1(nums []int) int {
	ans := int32(0)
	for i := 0; i < len(nums); i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num>>i) & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
