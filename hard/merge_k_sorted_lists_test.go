package hard

import (
	"testing"
)

func TestMergeKSortedLists(t *testing.T) {

	// Arange
	list111 := &ListNode{Val: 2, Next: nil}
	list11 := &ListNode{Val: 4, Next: list111}
	list1 := &ListNode{Val: 3, Next: list11}

	list222 := &ListNode{Val: 5, Next: nil}
	list22 := &ListNode{Val: 6, Next: list222}
	list2 := &ListNode{Val: 4, Next: list22}

	lists := []*ListNode{list1, list2}

	// Act
	result := mergeKLists(lists)

	// Assert
	if result.Val != 2 && result.Next.Next.Next.Next.Val != 6 {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, *result.Next.Next, *list11)
	}

}

func TestMergeKSortedEmptyLists(t *testing.T) {

	lists := []*ListNode{}

	// Act
	result := mergeKLists(lists)

	// Assert
	if result != nil {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, *result, *result)
	}

}
