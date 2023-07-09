package leetcode

import "fmt"

func Run20230707() {
	// https://leetcode.com/problems/maximize-the-confusion-of-an-exam/
	cases := []struct {
		name      string
		answerKey string
		k         int
		want      int
	}{
		{
			name:      "example 1",
			answerKey: "TTFF",
			k:         2,
			want:      4,
		},
		{
			name:      "example 2",
			answerKey: "TFFT",
			k:         1,
			want:      3,
		},
		{
			name:      "example 3",
			answerKey: "TTFTTFTT",
			k:         1,
			want:      5,
		},
	}

	for _, c := range cases {
		got := maxConsecutiveAnswers(c.answerKey, c.k)
		if got != c.want {
			fmt.Printf("%s: got %d, want %d\n", c.name, got, c.want)
		}
	}
}

func maxConsecutiveAnswers(answerKey string, k int) int {
}
