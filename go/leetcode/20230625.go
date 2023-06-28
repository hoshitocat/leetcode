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

func countRoutes(loc []int, start int, finish int, fuel int) int {
	var dp func(int, int) uint64

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	dp = func(i int, f int) uint64 {
		if i == finish && f == 0 {
			return 1
		}
		if f == 0 {
			return 0
		}

		var a uint64 = 0
		if i == finish {
			a++
		}

		for li := range loc {
			if li != i {
				d := abs(loc[li] - loc[i])
				if f >= d {
					a += dp(li, f-d)
				}
			}
		}
		return a % 1000000007
	}

	dp = memo(dp)
	return int(dp(start, fuel))
}

func memo(fn func(int, int) uint64) func(int, int) uint64 {
	m := make(map[string]uint64)
	d := func(i int, f int) uint64 {
		k := fmt.Sprintf("%d-%d", i, f)
		if v, ok := m[k]; ok {
			return v
		}

		v := fn(i, f)
		m[k] = v
		return v
	}
	return d
}
