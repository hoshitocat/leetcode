package leetcode

import (
	"fmt"
)

func Run20230625() {
	// https://leetcode.com/problems/count-all-possible-routes/

	cases := []struct {
		name      string
		locations []int
		start     int
		finish    int
		fuel      int
		expected  int
	}{
		{
			name:      "example 1",
			locations: []int{2, 3, 6, 8, 4},
			start:     1,
			finish:    3,
			fuel:      5,
			expected:  4,
		},
		{
			name:      "example 2",
			locations: []int{4, 3, 1},
			start:     1,
			finish:    0,
			fuel:      6,
			expected:  5,
		},
		{
			name:      "example 3",
			locations: []int{5, 2, 1},
			start:     0,
			finish:    2,
			fuel:      3,
			expected:  0,
		},
	}

	for _, c := range cases {
		actual := countRoutes(c.locations, c.start, c.finish, c.fuel)
		if actual != c.expected {
			fmt.Printf("case %s failed, expected: %d, actual: %d\n", c.name, c.expected, actual)
		}
	}
}

func countRoutes(locations []int, start int, finish int, fuel int) int {
	ordered := make([]int, len(locations))
	for i := 1; i < len(locations); i++ {
		if finish == i {
			continue
		}
		ordered[i] = locations[i]
	}
	ordered[0] = locations[start]
	ordered[len(ordered)-1] = locations[finish]

	dp := make([][]int, len(ordered))
	dp[0] = make([]int, fuel)

	for i := 1; i < len(ordered); i++ {
		dp[i] = make([]int, fuel)
		for j := range dp[i] {
			con := absInt(ordered[j-1] - ordered[j])

			if 0 <= con {
				continue
			}
		}
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
