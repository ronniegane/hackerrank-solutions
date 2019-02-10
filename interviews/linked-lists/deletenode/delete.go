package delete

// Complete the deleteNode function below.

// SinglyLinkedListNode is a single node in a list
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func deleteNode(head *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// Special case 1: deleting the head of the list
	if position == 0 {
		return head.next
	}

	// iterate through list to find position to delete
	marker := head
	for position > 1 {
		marker = marker.next
		position--
	}

	// delete the next node
	if marker.next.next == nil {
		// special case 2: deleting end of the list
		marker.next = nil
	} else {
		// common case: just skip this node
		marker.next = marker.next.next
	}

	return head
}
