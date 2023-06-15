package leetcode

import (
	"fmt"
)

func Run20230615() {
	cases := []struct {
		name string
		want int
		node *TreeNode
	}{
		{
			name: "test case 1",
			want: 2,
			node: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   -8,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "test case 2",
			want: 2,
			node: &TreeNode{
				Val: 989,
				Right: &TreeNode{
					Val: 10250,
					Left: &TreeNode{
						Val: 98693,
					},
					Right: &TreeNode{
						Val: -89388,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: -32127,
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		got := maxLevelSum(c.node)

		if got != c.want {
			fmt.Printf("failed %s: want is %d but got is %d\n", c.name, c.want, got)
		}
	}
}

func maxLevelSum(root *TreeNode) int {
	values := map[int]int{}
	sumValuesWithDepth(values, root, 1)

	maxDepth := 1
	maxValue := root.Val
	for depth, value := range values {
		if maxValue < value {
			maxDepth = depth
			maxValue = value
		}
	}

	return maxDepth
}

func sumValuesWithDepth(values map[int]int, node *TreeNode, depth int) {
	if node == nil {
		return
	}
	values[depth] += node.Val
	sumValuesWithDepth(values, node.Left, depth+1)
	sumValuesWithDepth(values, node.Right, depth+1)
}
