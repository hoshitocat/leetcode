// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
package main

import (
	"github.com/davecgh/go-spew/spew"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]
	var profit int
	for _, p := range prices[1:] {
		if p < min {
			min = p
		}
		diff := p - min
		if profit < diff {
			profit = diff
		}
	}

	return profit
}

func main() {
	cases := [][]int{
		{7,1,5,3,6,4},
		{7,6,4,3,1},
	}
	for _, c := range cases {
		spew.Dump(maxProfit(c))
	}
}
