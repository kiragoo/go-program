# 数组笔记

*模板为`golang`*

## 访问模式

* 迭代

```golang

for i,v := range items {
    ...
}
```

* 递归

```golang
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
}
```

## 常见结题思路

### 哈希表

[1. 两数之和 ](https://leetcode-cn.com/problems/two-sum/)
