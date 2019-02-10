package reverseprint

import "fmt"

// SinglyLinkedListNode is a single node in a list
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// print out the given list in reverse order
// you are passed the head of the node
func reversePrint(llist *SinglyLinkedListNode) {
	// recursive approach
	if llist == nil {
		return
	}
	// recurse before printing ensures the prints will occur
	// in reverse order
	reversePrint(llist.next)
	fmt.Println(llist.data)
}
