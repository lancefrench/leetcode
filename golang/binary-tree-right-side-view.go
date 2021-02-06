/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	rightView(root, &res, 0)
	return res
}

func rightView(curr *TreeNode, res *[]int, currDepth int) {
	if curr == nil {
		return
	}
	if currDepth == len(*res) {
		*res = append(*res, curr.Val)
	}

	rightView(curr.Right, res, currDepth+1)
	rightView(curr.Left, res, currDepth+1)
}