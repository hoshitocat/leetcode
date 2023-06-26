package leetcode

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Run20230624() {
	// https://leetcode.com/problems/tallest-billboard/

	cases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "case 1",
			input: []int{1, 2, 3, 6},
			want:  6,
		},
		{
			name:  "case 2",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  10,
		},
		{
			name:  "case 3",
			input: []int{1, 2},
			want:  0,
		},
	}

	for _, c := range cases {
		got := tallestBillboard(c.input)
		if diff := cmp.Diff(got, c.want); diff != "" {
			fmt.Printf("name: %s, input: %v, got: %v, want: %v\n", c.name, c.input, got, c.want)
			fmt.Printf("diff: %v\n", diff)
		}
	}
}

func tallestBillboard(rods []int) int {
	sum := 0
	for _, rod := range rods {
		sum += rod
	}

	dp := make([]int, sum+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for _, rod := range rods {
		dpCopy := make([]int, len(dp))
		copy(dpCopy, dp)

		for i := 0; i <= sum-rod; i++ {
			if dpCopy[i] < 0 {
				continue
			}

			dp[i+rod] = max(dp[i+rod], dpCopy[i])

			dp[abs(-i+rod)] = max(dp[abs(-i+rod)], dpCopy[i]+min(i, rod))
		}
	}

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
