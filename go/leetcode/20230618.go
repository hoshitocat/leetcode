package leetcode

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Run20230618() {
	// https://leetcode.com/problems/pascals-triangle-ii/
	cases := []struct {
		name  string
		input int
		want  []int
	}{
		{name: "case1", input: 3, want: []int{1, 3, 3, 1}},
		{name: "case2", input: 0, want: []int{1}},
		{name: "case3", input: 1, want: []int{1, 1}},
		{name: "case4", input: 21, want: []int{1, 21, 210, 1330, 5985, 20349, 54264, 116280, 203490, 293930, 352716, 352716, 293930, 203490, 116280, 54264, 20349, 5985, 1330, 210, 21, 1}},
	}

	for _, ca := range cases {
		got := getRow(ca.input)
		if diff := cmp.Diff(got, ca.want); diff != "" {
			fmt.Printf("want %v but got %v: diff = %s\n", ca.want, got, diff)
		}
	}
}

func getRow(rowIndex int) []int {
	var res []int

	for i := 0; i <= rowIndex; i++ {
		temp := make([]int, len(res))
		copy(temp, res)
		res = append(res, 1)
		for j := 1; j < len(res)-1; j++ {
			res[j] = temp[j-1] + temp[j]
		}
	}
	return res
}
