package easy

import (
	"testing"
)

func TestMergeTwoSorted(t *testing.T) {

	// Arange
	list111 := &ListNode{Val: 4, Next: nil}
	list11 := &ListNode{Val: 2, Next: list111}
	list1 := &ListNode{Val: 1, Next: list11}

	list222 := &ListNode{Val: 4, Next: nil}
	list22 := &ListNode{Val: 3, Next: list222}
	list2 := &ListNode{Val: 1, Next: list22}

	// Act
	result := mergeTwoLists(list1, list2)

	// Assert
	if result.Next.Next != list11 {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, *result.Next.Next, *list11)
	}

}
