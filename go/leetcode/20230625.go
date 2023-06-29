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
	var fn func(int, int) uint64
	fn = func(index, remainingFuel int) uint64 {
		if index == finish && remainingFuel == 0 {
			return 1
		}

		if remainingFuel == 0 {
			return 0
		}

		var a uint64
		if index == finish {
			a++
		}

		for li := range loc {
			if li != index {
				fuelConsumption := absInt(loc[li] - loc[index])
				if remainingFuel >= fuelConsumption {
					a += fn(li, remainingFuel-fuelConsumption)
				}
			}
		}
		return a
	}

	memo := make(map[string]uint64)
	memoFn := func(index int, remainingFuel int) uint64 {
		memoKey := fmt.Sprintf("%d-%d", index, remainingFuel)
		if value, ok := memo[memoKey]; ok {
			return value
		}

		value := fn(index, remainingFuel)
		memo[memoKey] = value
		return value
	}

	return int(memoFn(start, fuel))
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
