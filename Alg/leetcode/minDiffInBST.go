//783. 二叉搜索树节点最小距离
var pre int
var res int

func minDiffInBST(root *TreeNode) int {
	res = 1<<63 - 1
	pre = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if pre != 0 {
			res = min(res, root.Val-pre)
		}
		pre = root.Val
		dfs(root.Right)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


//
func minDiffInBST(root *TreeNode) int {
    ans, pre := math.MaxInt64, -1
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if pre != -1 && node.Val-pre < ans {
            ans = node.Val - pre
        }
        pre = node.Val
        dfs(node.Right)
    }
    dfs(root)
    return ans
}
