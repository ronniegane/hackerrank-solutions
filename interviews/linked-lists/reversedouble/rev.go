package reversedouble

// DoublyLinkedListNode is a single node in a doubly-linked list
type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

// reverse a doubly linked list, returning a pointer to the new "head"
// of the list
func reverse(head *DoublyLinkedListNode) *DoublyLinkedListNode {
	// reverse this node
	head.next, head.prev = head.prev, head.next

	// if we've reached the end of the list, return the head
	if head.prev == nil {
		return head
	}
	// else recurse to the next node (go to "prev" as it was "next" in the original list)
	return reverse(head.prev)
}
