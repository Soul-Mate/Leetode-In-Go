/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (33.17%)
 * Total Accepted:    95.4K
 * Total Submissions: 287.8K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if (l1.Val == 0 && l1.Next == nil) && (l2.Val == 0 && l2.Next == nil) {
		return l1
	}

	if l1.Val == 0 && l1.Next == nil {
		return l2
	}

	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	var x, y, carry int

	var head, rear *ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}

		sum := x + y + carry

		carry = sum / 10

		if head == nil {
			rear = &ListNode{
				Val: sum % 10,
			}

			head = rear
		} else {
			temp := &ListNode{
				Val: sum % 10,
			}

			rear.Next = temp
			rear = temp
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		rear.Next = &ListNode{Val: carry}
	}

	return head
}

/*
[0,8,6,5,6,8,3,5,7]
[6,7,8,0,8,5,8,9,7]
*/
func main() {
	l1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	for head := result; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
