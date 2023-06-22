package leetcode

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Run20230622() {
	// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
	cases := []struct {
		name   string
		wants  int
		prices []int
		fee    int
	}{
		{
			name:   "case1",
			wants:  8,
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
		},
		{
			name:   "case2",
			wants:  6,
			prices: []int{1, 3, 7, 5, 10, 3},
			fee:    3,
		},
	}

	for _, ca := range cases {
		got := maxProfit(ca.prices, ca.fee)

		if diff := cmp.Diff(ca.wants, got); diff != "" {
			fmt.Printf("wants %v but got %v: diff = %s\n", ca.wants, got, diff)
		}
	}
}

func maxProfit(prices []int, fee int) int {
	var maxProfit int
}
