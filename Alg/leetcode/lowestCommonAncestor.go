//235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//根据二叉搜索树特点
	if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
		return root
	}

	//二叉搜索树 左子小于root 右树 大于root
	if p.Val < root.Val {
		root = root.Left
	} else {
		root = root.Right
	}

	return lowestCommonAncestor(root, p, q)
}