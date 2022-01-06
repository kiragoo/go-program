package tree

func LC144(root *TreeNode) (res []int) {
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	return
}

func LC144_1(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
