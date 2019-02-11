"""
Detect a cycle in a singly- linked list. Note that the head pointer may be 'None' if the list is empty.

A Node is defined as: 
 
    class Node(object):
        def __init__(self, data = None, next_node = None):
            self.data = data
            self.next = next_node
"""



def has_cycle(head):
    # attempt 1: using a set to keep track of seen nodes
    seen = set()
    while head != None:
        if head in seen:
            # Have visited this node before
            return True
            # Mark this node as seen
        seen.add(head)
        # go to the next node
        head = head.next

    # Have reached the end of the list
    return False

"""
Floyd's tortise-and-hare cycle detecting algorithm
https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_Tortoise_and_Hare
no set needed, less memory, but potentially more evaluations

two pointers: one "fast" moves through list two at a time,
one "slow" moves through one at a time.

If one reaches the end of the list, no cycle.

If they ever point to the same value, there is a cycle.
"""

def has_cycle_floyd(head):
    # Using Floyd's tortoise and hare algorithm
    fast = head
    slow = head
    while fast != None and fast.next != None:
        fast = fast.next.next
        slow = slow.next
        if fast == slow:
            return True # cycle found
    # If either the fast pointer or its child is None then 
    # we have found the end of the list
    # don't need to check slow pointer, fast will get there first
    return False

"""
Brent's Algorithm
https://en.wikipedia.org/wiki/Cycle_detection#Brent's_algorithm

this also finds the length of the cycle directly
"""
