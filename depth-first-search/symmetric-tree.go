/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var mirror func(left, right *TreeNode) bool

    mirror = func(left, right *TreeNode) bool {
        if left == nil && right == nil {
            return true
        }

        if left == nil || right == nil {
            return false
        }

        return left.Val == right.Val &&
            mirror(left.Left, right.Right) &&
            mirror(left.Right, right.Left)
    }

    return mirror(root.Left, root.Right)
}