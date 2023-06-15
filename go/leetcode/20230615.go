package leetcode

import (
	"fmt"
	"math"
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
	level := 0
	minLevel := 0
	maxValue := math.MinInt64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level++
		currentSum := 0
		n := len(queue)

		for n > 0 {
			node := queue[0]
			queue = queue[1:]
			currentSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			n--
		}

		if maxValue < currentSum {
			maxValue = currentSum
			minLevel = level
		}
	}

	return minLevel
}
