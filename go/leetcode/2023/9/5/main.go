// https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "github.com/davecgh/go-spew/spew"

func rotate(nums []int, k int)  {
	l := len(nums)
	if k == 0 || k == l {
		return
	}

	if l < k {
		k %= l
	}

	kl := l - k
	tmp := make([]int, k)
	copy(tmp, nums[kl:l])
	copy(nums[k:l], nums[0:kl])
	copy(nums[0:l], tmp)
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
