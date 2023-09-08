package main

import "github.com/davecgh/go-spew/spew"

func canJump(nums []int) bool {
	var max int
	for i, n := range nums {
		if max < i {
			return false
		}

		if max < i+n {
			max = i + n
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
