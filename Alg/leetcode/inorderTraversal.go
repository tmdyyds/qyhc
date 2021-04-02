//94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var recursive func(node *TreeNode)
	res := []int{}
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}

		recursive(node.Left)
		res = append(res, node.Val)
		recursive(node.Right)
	}

	recursive(root)

	return res
}


//第二种
var r []int
func inorderTraversal(root *TreeNode) []int {
	r = make([]int, 0)
	helper(root)
	return r
}

func helper(root *TreeNode) {
	if root != nil {
		helper(root.Left)
		r = append(r, root.Val)
		helper(root.Right)
	}
}