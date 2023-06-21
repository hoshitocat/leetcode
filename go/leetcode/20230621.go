package leetcode

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Run20230621() {
	// https://leetcode.com/problems/k-radius-subarray-averages/
	cases := []struct {
		name  string
		wants []int
		nums  []int
		k     int
	}{
		{
			name: "case1",
			wants: []int{
				-1, -1, -1, 5, 4, 4, -1, -1, -1,
			},
			nums: []int{7, 4, 3, 9, 1, 8, 5, 2, 6},
			k:    3,
		},
		{
			name: "case2",
			wants: []int{
				100000,
			},
			nums: []int{100000},
			k:    0,
		},
		{
			name: "case3",
			wants: []int{
				-1,
			},
			nums: []int{8},
			k:    100000,
		},
	}

	for _, ca := range cases {
		got := getAverages(ca.nums, ca.k)

		if diff := cmp.Diff(ca.wants, got); diff != "" {
			fmt.Printf("wants %v but got %v: diff = %s\n", ca.wants, got, diff)
		}
	}
}

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}

	div := k*2 + 1
	size := len(nums)
	const defaultAvg = -1
	result := make([]int, size)
	var sum int
	for i := 0; i < size; i++ {
		begin := i - k
		end := (i + k) + 1
		if begin < 0 || size < end {
			result[i] = defaultAvg
			continue
		}

		if sum == 0 {
			sum = sumOfSlice(nums[begin:end])
		} else {
			sum -= nums[begin-1]
			sum += nums[end-1]
		}
		result[i] = sum / div
	}

	return result
}

func sumOfSlice(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	return sum
}
