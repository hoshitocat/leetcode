package leetcode

import (
	"fmt"
	"strconv"
)

func Run230612() {
	summaryRanges([]int{0, 1, 2, 4, 5, 7})
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var head int
	var result []string
	for i := range nums {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}

		if head == i {
			result = append(result, strconv.Itoa(nums[i]))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", nums[head], nums[i]))
		}

		head = i + 1
	}

	return result
}
