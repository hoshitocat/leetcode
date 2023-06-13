package leetcode

import (
	"fmt"
	"strconv"
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
	rows := make([]string, len(grid))
	columns := map[string]int{}
	var sameCount int

	for i := range grid {
		var row string
		var column string
		for j := range grid {
			row += strconv.Itoa(grid[i][j]) + "|"
			column += strconv.Itoa(grid[j][i]) + "|"
		}
		rows[i] = row
		columns[column]++
	}

	for _, row := range rows {
		count, ok := columns[row]
		if ok {
			sameCount += count
		}
	}

	return sameCount
}
