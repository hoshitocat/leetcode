package leetcode

import "github.com/davecgh/go-spew/spew"

func Run230329() {
	spew.Dump(twoSum([]int{1, 2, 3, 4, 5}, 5))
}

func twoSum(nums []int, target int) []int {
	records := map[int]int{}
	result := make([]int, 2)
	for key, value := range nums {
		k, ok := records[target-value]
		if ok {
			result = []int{k, key}
		}

		records[value] = key
	}

	return result
}
