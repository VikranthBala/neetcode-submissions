/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var explore func(*TreeNode) int

	explore = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return 1 + max(explore(root.Left), explore(root.Right))
	}

	return explore(root)
}
