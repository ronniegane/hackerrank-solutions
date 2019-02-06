package double

import "testing"

var q, r, s, t, u, v, w, x, y, z DoublyLinkedListNode

// Set up lists for testing
func setup() {

}

// Testing equality of doubly linked lists
func equalList(a, b *DoublyLinkedListNode) bool {
	if a == nil && b == nil {
		// Have reached end of list
		return true
	}
	if a == nil && b != nil || a != nil && b == nil {
		// lists are different lengths
		return false
	}
	if a.data != b.data {
		// discrepancy found
		return false
	}
	// test next node for equality
	return equalList(a.next, b.next)
}

func TestShouldBeTrue(t *testing.T) {
	// Create identical lists to test equality code
	var a, b, c, d, e, f DoublyLinkedListNode
	a.data = 2
	a.next = &b
	b.data = 4
	b.prev = &a
	b.next = &c
	c.data = 8
	c.prev = &b

	d.data = 2
	d.next = &e
	e.data = 4
	e.prev = &d
	e.next = &f
	f.data = 8
	f.prev = &e

	if !equalList(&a, &d) {
		t.Errorf("Lists not identical")
	}

}

func TestShouldBeFalse(t *testing.T) {
	// Create identical lists to test equality code
	var a, b, c, d, e, f DoublyLinkedListNode
	a.data = 2
	a.next = &b
	b.data = 4
	b.prev = &a
	b.next = &c
	c.data = 8
	c.prev = &b

	d.data = 2
	d.next = &e
	e.data = 4
	e.prev = &d
	e.next = &f
	f.data = 7
	f.prev = &e

	if equalList(&a, &d) {
		t.Errorf("Different lists should not pass equality check")
	}

}
