/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	travel(root, &sum)
	return root
}

func travel(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	travel(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	travel(root.Left, sum)
}