// Remove Nth Node From End of List
// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:
// 1 -> 2 -> 3 -> 4 -> 5
// 1 -> 2 -> 3 -> 5
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

// Example 2:
// Input: head = [1], n = 1
// Output: []

package medium

import (
	m "leetcode/pkg/model"
)

func removeNthFromEnd(head *m.ListNode, n int) *m.ListNode {

	if head == nil {
		return nil
	}
	counter := 1

	dummy := new(m.ListNode)
	dummy.Next = head

	tmp, tmpNth := head, dummy

	for tmp != nil {

		if counter > n {
			tmpNth = tmpNth.Next
		}

		if tmp.Next == nil {
			tmpNth.Next = tmpNth.Next.Next
		}

		counter++
		tmp = tmp.Next

	}

	return dummy.Next
}
