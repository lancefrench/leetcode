package problem0095

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestProblem0094(t *testing.T) {
	testCases := []struct {
		n    int
		want [](*TreeNode)
	}{
		{1, []*TreeNode{
			{Val: 1, Left: nil, Right: nil},
		}},
		{3, []*TreeNode{
			{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}},
			{Val: 1, Left: nil, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}},
			{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			{Val: 3, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}, Right: nil},
			{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: nil}, Right: nil},
		}},
	}

	for _, tc := range testCases {
		actual := generateTrees(tc.n)
		expected := tc.want
		assert.Equal(t, expected, actual)
	}
}
