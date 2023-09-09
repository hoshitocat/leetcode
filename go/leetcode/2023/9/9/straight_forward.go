package main

import (
	"github.com/davecgh/go-spew/spew"
)

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func jump(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] = max(nums[i]+i, nums[i-1])
	}

	var index, answer int
	for index < n-1 {
		answer++
		index = nums[index]
	}

	return answer
}

func main() {
	cases := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 0, 1, 4},
	}

	for _, c := range cases {
		spew.Dump(jump(c))
	}
}
