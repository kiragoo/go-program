package array

import (
	"sort"
	"strconv"
)

func LC506(score []int) []string {
	rawS := make([]int,len(score))
	copy(rawS, score)
	m := map[int]string{
		0: "Gold Medal",
		1: "Silver Medal",
		2: "Bronze Medal",
	}
	pm := map[int]int{}
	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	for i,v := range score {
		pm[v] = i
	}
	res := []string{}
	for _, v := range rawS{
		_, ok := m[pm[v]]
		if ok {
			res = append(res, m[pm[v]])
		}else {
			res = append(res, strconv.Itoa(pm[v]+1))
		}
	}
	return res
}
