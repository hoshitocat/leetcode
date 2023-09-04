// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
package main

import "github.com/davecgh/go-spew/spew"

func majorityElement(nums []int) int {
	var candidate int
	var count int

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}

		if candidate == n {
			count++
		} else {
			count--
		}
	}

	return candidate
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
