package double

// DoublyLinkedListNode represents one node in a doubly-linked list
type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

// SortedInsert creates a new node with the data provided
// and inserts it into the correct position in a sorted,
// doubly-linked list.
func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	node := DoublyLinkedListNode{data, nil, nil}
	// Special case if data belongs at head of list
	if data < head.data {
		head.prev = &node
		node.next = head
		return &node
	}
	current := head
	next := head.next
	// Find the spot in the list where the data sits in order
	for data < next.data {
		current = next
		next = current.next
	}
	// Update references so current -> node -> next
	current.next = &node
	next.prev = &node
	node.prev = current
	node.next = next

	return head
}
