package tree

func LC145(root *TreeNode) (res []int) {
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
