package tree

// 1
//  \
//   2
//  /
// 3
func LC94(root *TreeNode) (res []int) {
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
	helper(root)
	return
}

func LC94_1(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
