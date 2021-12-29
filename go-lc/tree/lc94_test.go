package tree

import (
	"fmt"
	"testing"
)

func TestLC94(t *testing.T) {
	root := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(LC94(&root))
}
