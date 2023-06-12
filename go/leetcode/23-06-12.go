package leetcode

import "fmt"

func Run230612() {
	summaryRanges([]int{0, 1, 2, 4, 5, 7})
}

func summaryRanges(nums []int) []string {
	var result []string
	var tmp []int
	for _, n := range nums {
		if len(tmp) == 0 {
			tmp = append(tmp, n)
			continue
		}

		lastIndex := len(tmp) - 1

		if tmp[lastIndex]+1 == n {
			tmp = append(tmp, n)
			continue
		}

		if len(tmp) == 1 {
			result = append(result, fmt.Sprintf("%d", tmp[0]))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", tmp[0], tmp[lastIndex]))
		}

		tmp = []int{n}
	}

	if len(tmp) == 0 {
		return nil
	}

	if len(tmp) == 1 {
		result = append(result, fmt.Sprintf("%d", tmp[0]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", tmp[0], tmp[len(tmp)-1]))
	}

	return result
}
