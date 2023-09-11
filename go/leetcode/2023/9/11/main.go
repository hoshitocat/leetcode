// https://leetcode.com/problems/h-index/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	return citations[n/2]
}

func main() {
	cases := [][]int{
		{3,0,6,1,5},
		{1,3,1},
	}

	for _, c := range cases {
		spew.Dump(hIndex(c))
	}
}
