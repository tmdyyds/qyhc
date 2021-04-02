/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//145. 二叉树的后序遍历
func preorderTraversal(root *TreeNode) []int {
	var recursive func(node *TreeNode)
	res := []int{}
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}

		recursive(node.Left)
		recursive(node.Right)
		res = append(res, node.Val)
	}

	recursive(root)

	return res
}