// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
package main

import "github.com/davecgh/go-spew/spew"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	i := 2
	for j := 2; j < length; j++ {
		if nums[i-2] == nums[j] && nums[i-1] == nums[j] {
			continue
		}

		nums[i] = nums[j]
		i++
	}

	nums = nums[:i]
	return len(nums)
}

func main() {
	cases := [][]int{
		{1, 1, 1, 2, 2, 3},
		{0, 0, 1, 1, 1, 1, 2, 3, 3},
	}

	for _, c := range cases {
		k := removeDuplicates(c)
		spew.Dump(k)
	}
}
