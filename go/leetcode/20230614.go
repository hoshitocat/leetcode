package leetcode

import (
	"math"
	"sort"

	"github.com/davecgh/go-spew/spew"
)

//[1,0,48,null,null,12,49]

func Run20230614() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}

	spew.Dump(getMinimumDifference(root))
	spew.Dump(nums)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nums = make([]int, 0, 10_000)

func getMinimumDifference(root *TreeNode) int {
	getNodeValue(root)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// default
	min := 100_000
	for i := 0; i < len(nums)-1; i++ {
		diff := intDiff(nums[i], nums[i+1])

		if diff < min {
			min = diff
		}
	}

	return min
}

func getNodeValue(current *TreeNode) {
	if current == nil {
		return
	}

	nums = append(nums, current.Val)
	getNodeValue(current.Left)
	getNodeValue(current.Right)

	return
}

func intDiff(x, y int) int {
	return int(math.Abs(float64(x - y)))
}
