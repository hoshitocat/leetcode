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
	dp := make([][2]int, len(prices))
	dp[0][0] = 0          //0 stands for hold cash on day 0
	dp[0][1] = -prices[0] //1 stands for hold stock on day 0

	for i := 1; i < len(prices); i++ {
		//if the decision is to hold cash on day i, there are 2 options
		//option 1: sell the stock held from previous day
		//option 2: do nothing, keep on holding cash like previous day
		//max profit of day i with cash is max(option1, option2)
		dp[i][0] = maxInt(prices[i]+dp[i-1][1]-fee, dp[i-1][0])

		//if the decision is to hold stock on day i, there are also 2 options
		//option 1: buy stock, based on previous day's cash profit
		//option 2: do nothing, keep on holding stock like previous day
		//max profit of day i with stock is max(option1, option2)
		dp[i][1] = maxInt(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	return dp[len(prices)-1][0]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
