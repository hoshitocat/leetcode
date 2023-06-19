package leetcode

import "github.com/davecgh/go-spew/spew"

func Run20230619() {
	// https://leetcode.com/problems/find-the-highest-altitude/
	spew.Dump(largestAltitude([]int{-5, 1, 5, 0, -7}))
}

func largestAltitude(gain []int) int {
	var highest int
	var current int
	for _, g := range gain {
		current += g

		if highest < current {
			highest = current
		}
	}

	return highest
}
