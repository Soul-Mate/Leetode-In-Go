/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (32.52%)
 * Total Accepted:    33.1K
 * Total Submissions: 101.7K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p1, p2 *ListNode

	dummy := &ListNode{
		Next: head,
	}

	p1 = dummy

	p2 = dummy

	for i := 1; i <= n+1; i++ {
		p1 = p1.Next
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		// Next: &ListNode{
		// 	Val: 2,
		// 	// Next: &ListNode{
		// 	// 	Val: 3,
		// 	// 	Next: &ListNode{
		// 	// 		Val: 4,
		// 	// 		Next: &ListNode{
		// 	// 			Val: 5,
		// 	// 		},
		// 	// 	},
		// 	// },
		// },
	}

	head = removeNthFromEnd(head, 1)

	for ; head != nil; head = head.Next {
		println(head.Val)
	}
}
