// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"github.com/davecgh/go-spew/spew"
)

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func maxProfit(prices []int) int {
	var tradeFn func(int) (int, int)
	tradeFn = func(day int) (int, int) {
		if day == 0 {
			return -prices[0], 0
		}
		prevHold, prevNotHold := tradeFn(day-1)
		hold := max(prevHold, prevNotHold - prices[day])
		notHold := max(prevNotHold, prevHold + prices[day])
		return hold, notHold
	}

	_, profit := tradeFn(len(prices) - 1)
	return profit
}

func main() {
	cases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}

	for _, c := range cases {
		spew.Dump(maxProfit(c))
	}
}
