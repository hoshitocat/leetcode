package leetcode

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Run20230617() {
	// https: //leetcode.com/problems/add-two-numbers/

	cases := []struct {
		name string
		want []int
		nums []*ListNode
	}{
		{
			name: "test case 1",
			want: []int{7, 0, 8},
			nums: []*ListNode{
				{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "test case 2",
			want: []int{0},
			nums: []*ListNode{
				{Val: 0},
				{Val: 0},
			},
		},
		{
			name: "test case 3",
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
			nums: []*ListNode{
				{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		node := addTwoNumbers(c.nums[0], c.nums[1])
		var got []int
		for node != nil {
			got = append(got, node.Val)
			node = node.Next
		}

		if diff := cmp.Diff(got, c.want); diff != "" {
			fmt.Printf("failed %s: want is %d but got is %d: %s\n", c.name, c.want, got, diff)
		} else {
			fmt.Printf("%s succeeded\n", c.name)
		}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode
	var currentNode *ListNode
	var nextNum int
	for {
		var val int
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		val += nextNum
		nextNum = val / 10
		val %= 10

		node := ListNode{
			Val:  val,
			Next: nil,
		}
		if root == nil {
			root = &node
			currentNode = &node
		} else {
			currentNode.Next = &node
			currentNode = &node
		}

		if l1 == nil && l2 == nil {
			if nextNum != 0 {
				currentNode.Next = &ListNode{Val: nextNum}
			}
			break
		}
	}

	return root
}
