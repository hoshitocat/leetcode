package leetcode

import (
	"fmt"
)

func Run20230613() {
	cases := [][][]int{
		{
			{3, 2, 1},
			{1, 7, 6},
			{2, 7, 7},
		},
		{
			{3, 1, 2, 2},
			{1, 4, 4, 5},
			{2, 4, 2, 2},
			{2, 4, 2, 2},
		},
		{
			{3, 1, 2, 2},
			{1, 4, 4, 4},
			{2, 4, 2, 2},
			{2, 5, 2, 2},
		},
		{
			{11, 1},
			{1, 11},
		},
	}

	for i, c := range cases {
		fmt.Printf("case %d: answer is %d\n", i+1, equalPairs(c))
	}
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	mapping := make(map[[200]int]int)
	arr := [200]int{}
	for i := 0; i < n; i++ {
		copy(arr[:], grid[i])
		mapping[arr]++
	}

	var sameCount int
	for i := 0; i < n; i++ {
		arr := [200]int{}
		for j := 0; j < n; j++ {
			arr[j] = grid[j][i]
		}

		count, ok := mapping[arr]
		if ok {
			sameCount += count
		}
	}

	return sameCount
}
