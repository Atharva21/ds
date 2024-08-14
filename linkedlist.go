package ds

import "fmt"

type ListNode[T any] struct {
	Data T
	Next *ListNode[T]
}

type LinkedList[T any] struct {
	Head *ListNode[T]
	Tail *ListNode[T]
	Size uint
}

func (ll *LinkedList[T]) Add(data T) *ListNode[T] {
	defer func() {
		ll.Size++
	}()
	node := &ListNode[T]{Data: data}
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = ll.Head
		return node
	}
	ll.Tail.Next = node
	ll.Tail = ll.Tail.Next
	return node
}

func (ll *LinkedList[T]) Remove(node *ListNode[T]) error {
	if ll.Head == nil {
		return fmt.Errorf("list is empty")
	}
	defer func() {
		ll.Size--
	}()
	if ll.Head == node {
		ll.Head = ll.Head.Next
		return nil
	}
	for cur := ll.Head; cur != nil; cur = cur.Next {
		if cur.Next == node {
			cur.Next = cur.Next.Next
			return nil
		}
	}
	return fmt.Errorf("node not found")
}

func (ll *LinkedList[T]) Println() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Data)
	}
	fmt.Println()
}
