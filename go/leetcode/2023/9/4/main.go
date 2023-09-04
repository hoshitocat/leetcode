// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
package main

import "github.com/davecgh/go-spew/spew"

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	var maxVal int
	var cur int
	for n, val := range m {
		if maxVal < val {
			cur = n
			maxVal = val
		}
	}

	return cur
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
