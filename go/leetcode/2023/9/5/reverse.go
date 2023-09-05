// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "github.com/davecgh/go-spew/spew"

func rotate(nums []int, k int)  {
	l := len(nums)
	k %= l
	reverse(nums, 0, l-k-1)
	reverse(nums, l-k, l-1)
	reverse(nums, 0, l-1)
}

func reverse(nums []int, start, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	cases := []struct{
		nums []int
		k int
	}{
		{
			nums: []int{1,2,3,4,5,6,7},
			k: 3,
		},
		{
			nums: []int{-1, -100, 3, 99},
			k: 2,
		},
		{
			nums: []int{-1},
			k: 2,
		},
	}

	for _, c := range cases {
		c := c
		rotate(c.nums, c.k)
		spew.Dump(c.nums)
	}
}
