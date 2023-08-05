package medium

import (
	m "leetcode/pkg/model"
	"testing"
)

func TestRemoveNthNode1(t *testing.T) {

	// Arange
	list11111 := &m.ListNode{Val: 5, Next: nil}
	list1111 := &m.ListNode{Val: 4, Next: list11111}
	list111 := &m.ListNode{Val: 3, Next: list1111}
	list11 := &m.ListNode{Val: 2, Next: list111}
	list1 := &m.ListNode{Val: 1, Next: list11}

	// Act
	result := removeNthFromEnd(list1, 2)
	should := 5

	// Assert
	if result.Next.Next.Next.Val != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result.Next.Next.Next.Val, should)
	}
}

func TestRemoveNthNode2(t *testing.T) {

	// Arange
	list1 := &m.ListNode{Val: 1, Next: nil}

	// Act
	result := removeNthFromEnd(list1, 1)
	var should m.ListNode

	// Assert
	if result != nil {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result.Next.Next.Next.Val, should)
	}
}

func TestRemoveNthNode3(t *testing.T) {

	// Arange

	list11 := &m.ListNode{Val: 2, Next: nil}
	list1 := &m.ListNode{Val: 1, Next: list11}

	// Act
	result := removeNthFromEnd(list1, 1)
	should := 1

	// Assert
	if result.Val != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result.Val, should)
	}
}

func TestRemoveNthNode4(t *testing.T) {

	// Arange

	list11 := &m.ListNode{Val: 2, Next: nil}
	list1 := &m.ListNode{Val: 1, Next: list11}

	// Act
	result := removeNthFromEnd(list1, 2)
	should := 2

	// Assert
	if result.Val != should {
		t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result.Val, should)
	}
}
