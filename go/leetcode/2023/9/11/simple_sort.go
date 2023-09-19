// https://leetcode.com/problems/h-index/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	size := len(citations)

	if citations[size-1] == 0 {
		return 0
	} else if size == 1 {
		return 1
	}

	var result int
	for h := 1; h <= size; h++ {
		if citations[size-h] < h {
			break
		}

		result = h
	}

	return result
}

func main() {
	cases := [][]int{
		{3, 0, 6, 1, 5},
		{1, 3, 1},
	}

	for _, c := range cases {
		spew.Dump(hIndex(c))
	}
}
