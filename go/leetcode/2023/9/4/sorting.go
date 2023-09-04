// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {return nums[i] < nums[j] })
	return nums[len(nums)/2]
}

func main() {
	cases := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}

	for _, c := range cases {
		spew.Dump(majorityElement(c))
	}
}
