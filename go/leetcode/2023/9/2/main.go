// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	cases := []struct {
		nums []int
	}{
		{nums: []int{1, 1, 2}},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
	}

	for _, c := range cases {
		fmt.Println(removeDuplicates(c.nums))
	}
}
