// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func removeElement(nums []int, val int) int {
	var i = 0;
	for _, n := range nums {
		if n != val {
			nums[i] = n
			i++
		}
	}

	nums = nums[:i]
	return len(nums)
}

func main() {
	cases := []struct {
		nums []int
		val int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val: 3,
		},
		{
			nums: []int{0,1,2,2,3,0,4,2},
			val: 2,
		},
	}

	for _, c := range cases {
		c := c
		fmt.Println(removeElement(c.nums, c.val))
	}
}
