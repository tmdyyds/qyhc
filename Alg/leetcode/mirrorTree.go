//剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	t := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(t)

	return root
}