/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
    l, r := root, root //左右节点
    q := []*TreeNode{}

    //左右两节点入队队列
    q = append(q, l)
    q = append(q, r)

    for len(q) > 0 {
        l = q[0] //左
        r = q[1] //右
        q = q[2:]


        if l == nil && r == nil {
            continue
        }

        if l == nil || r == nil {
            return false
        }

        if l.Val != r.Val {
            return false
        }

        q = append(q, l.Left)
        q = append(q, r.Right)

        q = append(q, r.Right)
        q = append(q, l.Left)
    }

    return true
}