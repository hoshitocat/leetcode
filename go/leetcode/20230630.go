package leetcode

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Run20230630() {
	// https://leetcode.com/problems/last-day-where-you-can-still-cross/
	cases := []struct {
		name  string
		row   int
		col   int
		cells [][]int
		want  int
	}{
		{
			name: "case 1",
			row:  2,
			col:  2,
			cells: [][]int{
				{1, 1},
				{2, 1},
				{1, 2},
				{2, 2},
			},
			want: 2,
		},
		{
			name: "case 2",
			row:  2,
			col:  2,
			cells: [][]int{
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
			},
			want: 1,
		},
		{
			name: "case 3",
			row:  3,
			col:  3,
			cells: [][]int{
				{1, 2},
				{2, 1},
				{3, 3},
				{2, 2},
				{1, 1},
				{1, 3},
				{2, 3},
				{3, 2},
				{3, 1},
			},
			want: 3,
		},
	}

	for _, c := range cases {
		got := latestDayToCross(c.row, c.col, c.cells)
		if diff := cmp.Diff(c.want, got); diff != "" {
			fmt.Printf("name: %s, got: %d, want: %d\n", c.name, got, c.want)
			fmt.Printf("diff: %s\n", diff)
		}
	}
}

var directions = []int{0, 1, 0, -1, 0}

func latestDayToCross(row int, col int, cells [][]int) int {
	l, r, day := 0, len(cells), 0
	for l <= r {
		day = l + (r-l)/2
		if isItPossibleToGetBottomLine() {
			l = day + 1
		} else {
			r = day - 1
		}
	}

	return r
}

func isPossibleTopToBottom(flooded map[string]struct{}, row int, col int, cells [][]int) bool {
	for r := 0; r < row; r++ {
		if r+1 == row {
			return true
		}
		for c := 0; c < col; c++ {
			key := keyFunc([]int{r, c})
			if _, ok := flooded[key]; ok {
				continue
			}
		}

		if _, ok := flooded[key]; ok {
			continue
		}

		if i+1 == row {
			return true
		}
	}
	return false
}

func keyFunc(s []int) string {
	return fmt.Sprintf("%d-%d", s[0], s[1])
}
