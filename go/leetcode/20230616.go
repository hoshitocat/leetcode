package leetcode

import (
	"fmt"
)

func Run20230616() {
	// NOTE: https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/
	cases := []struct {
		name string
		want int
		nums []int
	}{
		{
			name: "test case 1",
			want: 1,
			nums: []int{2, 1, 3},
		},
		{
			name: "test case 2",
			want: 5,
			nums: []int{3, 4, 5, 1, 2},
		},
		{
			name: "test case 3",
			want: 0,
			nums: []int{1, 2, 3},
		},
	}

	for _, c := range cases {
		got := numOfWays(c.nums)

		if got != c.want {
			fmt.Printf("failed %s: want is %d but got is %d\n", c.name, c.want, got)
		}
	}
}

var (
	modulo = int(1e9) + 7
)

func numOfWays(nums []int) int {
	triangle := pascalTriangle(len(nums) + 1)

	ways := calcNumOfWays(triangle, nums) - 1
	return ways % modulo
}

func calcNumOfWays(pascalTriangle [][]int, nums []int) int {
	if len(nums) <= 2 {
		return 1
	}

	left, right := getLessAndGreaterThan(nums)
	waysLeft := calcNumOfWays(pascalTriangle, left) % modulo
	waysRight := calcNumOfWays(pascalTriangle, right) % modulo

	permutations := pascalTriangle[len(left)+len(right)][len(right)] % modulo
	totalPermutations := permutations % modulo
	totalPermutations = (totalPermutations * waysLeft) % modulo
	totalPermutations = (totalPermutations * waysRight) % modulo

	return totalPermutations
}

func getLessAndGreaterThan(numbers []int) ([]int, []int) {
	left := []int{}
	right := []int{}

	number := numbers[0]
	for _, n := range numbers {
		if n < number {
			left = append(left, n)
			continue
		}

		if n > number {
			right = append(right, n)
		}
	}

	return left, right
}

func pascalTriangle(lenght int) [][]int {
	triangle := make([][]int, lenght)

	for i := 0; i < lenght; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = (triangle[i-1][j-1] + triangle[i-1][j]) % modulo
		}
	}

	return triangle
}
