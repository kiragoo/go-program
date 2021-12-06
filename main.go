package main

import "fmt"

func main() {
	x := 4
	ans := int32(0)
	rightFn := func(x, i int) int32 {
		x = x >> i
		return int32(x)
	}
	// tmp := rightFn(x, 1)
	// fmt.Println(tmp)
	ans += rightFn(x, 2) & 1
	fmt.Println(int(ans))
}
