package problem0094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type tc struct {
	params
	output
}

type params struct {
	root *TreeNode
}

type output struct {
	ans []int
}

func TestProblem0094(t *testing.T) {
	tcs := []tc{
		{
			params{
				&TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			output{
				[]int{1, 3, 2},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output
		actual := inorderTraversal(p.root)
		expected := o.ans
		assert.Equal(t, expected, actual)
	}
}
