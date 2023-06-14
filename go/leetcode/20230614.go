package leetcode

import (
	"math"

	"github.com/davecgh/go-spew/spew"
)

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

func getMinimumDifference(root *TreeNode) int {
	_, _, diff := getMinDiffFromBST(root)
	return diff
}

func getMinDiffFromBST(node *TreeNode) (int, int, int) {
	min, max, diff := node.Val, node.Val, math.MaxInt32

	if node.Left != nil {
		lmin, lmax, ldiff := getMinDiffFromBST(node.Left)
		min = lmin
		if tmpDiff := node.Val - lmax; tmpDiff < diff {
			diff = tmpDiff
		}

		if ldiff < diff {
			diff = ldiff
		}
	}

	if node.Right != nil {
		rmin, rmax, rdiff := getMinDiffFromBST(node.Right)
		max = rmax
		if tmpDiff := rmin - node.Val; tmpDiff < diff {
			diff = tmpDiff
		}

		if rdiff < diff {
			diff = rdiff
		}
	}

	return min, max, diff
}
