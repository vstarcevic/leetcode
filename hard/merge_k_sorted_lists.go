package hard

// Merge k Sorted Lists
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.

// Example 1:
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

// Example 2:
// Input: lists = []
// Output: []

// Example 3:
// Input: lists = [[]]
// Output: []

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	var elements = make([]int, 0)
	for _, v := range lists {
		for v != nil {
			elements = append(elements, v.Val)
			v = v.Next
		}
	}

	sort.Ints(elements)

	if len(elements) == 0 {
		return nil
	}

	result := &ListNode{}
	current := result

	for i, v := range elements {
		if i == 0 {
			result.Val = v
		} else {
			tmp := &ListNode{v, nil}
			current.Next = tmp
			current = current.Next
		}
	}

	return result

}
