package main

import "github.com/davecgh/go-spew/spew"

func canJump(nums []int) bool {
	target := len(nums) - 1

	for i := target - 1; 0 <= i; i-- {
		if target <= i + nums[i] {
			target = i
		}
	}

	return target == 0
}

func main() {
	cases := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
	}

	for _, c := range cases {
		spew.Dump(canJump(c))
	}
}
