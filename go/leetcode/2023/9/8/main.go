package main

import "github.com/davecgh/go-spew/spew"

func canJump(nums []int) bool {
	farthest := 0
	length := len(nums)
	for i, n := range nums {
		if i+n > farthest {
			farthest = i+n
		}

		if n == 0 && farthest < length - 1 && i == farthest {
			return false
		}
	}

	return true
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
