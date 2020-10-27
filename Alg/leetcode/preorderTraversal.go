//144. 二叉树的前序遍历

//结构体
type root struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
    递归实现
    时间复杂度：O(N)，N为二叉树节点。
    空间复杂度：O(N),最坏O(N),平均为 O(logN)
*/
func preorderTraversal(root *TreeNode) []int {
	var recursive func(node *TreeNode)
	res := []int{}
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}

		//前序遍历
		res = append(res, node.Val)
		recursive(node.Left)
		recursive(node.Right)

		/*		//中序遍历
				recursive(node.Left)
				res = append(res, node.Val)
				recursive(node.Right)

				//后序遍历
				recursive(node.Left)
				recursive(node.Right)
				res = append(res, node.Val)*/
	}

	recursive(root)
	return res
}