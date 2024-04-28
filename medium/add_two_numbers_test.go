package medium

import (
	m "leetcode/pkg/model"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {

	// Arange
	list111 := &m.ListNode{Val: 2, Next: nil}
	list11 := &m.ListNode{Val: 4, Next: list111}
	list1 := &m.ListNode{Val: 3, Next: list11}

	list222 := &m.ListNode{Val: 5, Next: nil}
	list22 := &m.ListNode{Val: 6, Next: list222}
	list2 := &m.ListNode{Val: 4, Next: list22}

	// Act
	result := addTwoNumbers(list1, list2)

	// Assert
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, *result.Next.Next, *list11)
	}

}

func TestAddTwoNumbers2(t *testing.T) {

	// Arange
	list1111111 := &m.ListNode{Val: 9, Next: nil}
	list111111 := &m.ListNode{Val: 9, Next: list1111111}
	list11111 := &m.ListNode{Val: 9, Next: list111111}
	list1111 := &m.ListNode{Val: 9, Next: list11111}
	list111 := &m.ListNode{Val: 9, Next: list1111}
	list11 := &m.ListNode{Val: 9, Next: list111}
	list1 := &m.ListNode{Val: 9, Next: list11}

	list22222 := &m.ListNode{Val: 9, Next: nil}
	list2222 := &m.ListNode{Val: 9, Next: list22222}
	list222 := &m.ListNode{Val: 9, Next: list2222}
	list22 := &m.ListNode{Val: 9, Next: list222}
	list2 := &m.ListNode{Val: 9, Next: list22}

	// Act
	result := addTwoNumbers(list1, list2)

	// Assert
	if result.Val != 8 || result.Next.Val != 9 || result.Next.Next.Val != 9 {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, *result.Next.Next, *list11)
	}

}

func TestAddTwoNumbers3(t *testing.T) {

	// Arange
	list111 := &m.ListNode{Val: 2, Next: nil}
	list11 := &m.ListNode{Val: 4, Next: list111}
	list1 := &m.ListNode{Val: 9, Next: list11}

	list22222 := &m.ListNode{Val: 5, Next: nil}
	list222 := &m.ListNode{Val: 6, Next: list22222}
	list22 := &m.ListNode{Val: 4, Next: list222}
	list2 := &m.ListNode{Val: 9, Next: list22}

	// Act
	result := addTwoNumbers(list1, list2)

	// Assert
	if result.Val != 8 || result.Next.Val != 9 {
		t.Fatalf(`Result not as expected. Got %v, and it should be %v`, *result.Next.Next, *list11)
	}

}
