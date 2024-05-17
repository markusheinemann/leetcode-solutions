// https://leetcode.com/problems/delete-leaves-with-a-given-value
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left = removeLeafNodes(root.Left, target)
    root.Right = removeLeafNodes(root.Right, target)
    
    if root.Right == nil && root.Left == nil && root.Val == target {
        return nil
    }
    return root
}
