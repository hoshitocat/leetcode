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
	left := 1
	right := len(cells)
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		if canWalk(cells, row, col, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func canWalk(cells [][]int, row, col, dayAt int) bool {
	grid := make([][]int, row)
	for i := range grid {
		grid[i] = make([]int, col)
	}
	for i := 0; i < dayAt; i++ {
		grid[cells[i][0]-1][cells[i][1]-1] = 1
	}

	bfs := make([][]int, 0)
	for c := 0; c < col; c++ {
		if grid[0][c] == 0 {
			bfs = append(bfs, []int{0, c})
			grid[0][c] = 1
		}
	}

	for len(bfs) > 0 {
		cur := bfs[0]
		bfs := bfs[1:]
		r := cur[0]
		c := cur[1]
		if r == row-1 {
			return true
		}
		for i := 0; i < 4; i++ {
			nextR := r + directions[i]
			nextC := c + directions[i+1]
		}

	}

}
