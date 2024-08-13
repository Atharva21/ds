package linkedlist

type ListNode[T any] struct {
	Data T
	Next *ListNode[T]
}
