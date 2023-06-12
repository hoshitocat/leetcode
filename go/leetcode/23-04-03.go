package leetcode

import (
	"sort"

	"github.com/davecgh/go-spew/spew"
)

func Run230403() {
	spew.Dump(numRescueBoats([]int{1, 2, 3, 4, 5}, 5))
}

func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})

	result := 0
	left := 0
	right := len(people) - 1
	for left <= right {
		sum := people[left] + people[right]

		if sum <= limit {
			result++
			left++
			right--
		} else {
			result++
			left++
		}
	}

	return result
}
