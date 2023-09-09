package main

import (
	"math"

	"github.com/davecgh/go-spew/spew"
)

func jump(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i <= nums[0]; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n - 1; i++ {
		for j := i + 1; j <= i+nums[i]; j++ {
			dp[i][j] = dp[i-1][i] + 1
		}
	}

	min := math.MaxInt32
	for _, d := range dp[:n-1] {
		if d[n-1] == 0 {
			continue
		}
		if d[n-1] < min {
			min = d[n-1]
		}
	}

	return min
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
