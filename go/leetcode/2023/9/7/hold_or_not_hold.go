// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&envId=top-interview-150

package main

import (
	"math"

	"github.com/davecgh/go-spew/spew"
)

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func maxProfit(prices []int) int {
	hold,notHold := math.MinInt32, 0
	for _, p := range prices {
		prevHold, prevNotHold := hold, notHold
		hold = max(prevHold, prevNotHold - p)
		notHold = max(prevNotHold, prevHold + p)
	}

	return notHold
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
