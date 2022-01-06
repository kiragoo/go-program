package tree

func LC100(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return false
	}
	return LC100(p.Left, q.Left) && LC100(p.Right, q.Right)
}
