/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1

	var explore func(root *TreeNode, depth int) int

	explore = func(root *TreeNode, depth int) int {
		{
			if root == nil {
				return 0
			}
			return depth + max(explore(root.Left, depth), explore(root.Right, depth))
		}
	}
	return explore(root, depth)
}
