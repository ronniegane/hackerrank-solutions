# Find the point where two singly-linked lists merge.


# For your reference:
#
# SinglyLinkedListNode:
#     int data
#     SinglyLinkedListNode next


# head1 and head2 are singly-linked lists
# return the data point where they merge
def findMergeNode(head1, head2):
    # Thought: keep track of visited node in a set
    # return when either head1 or head2 is in that set
    if head1 != None :
        if head1 in seen:
            return head1.data
        seen.add(head1)
        # move to the next node
        head1 = head1.next
    if head2 != None :
        if head2 in seen:
            return head2.data
        seen.add(head2)
        head2 = head2.next
    
    # move to the next place in each
    return findMergeNode(head1, head2)


## Problem setter's solution
# iterate through both lists to get their length
# shift the pointer of larger list forward by the difference 
# between these lengths. This ensures that you will reach the merge point
# by moving forward both pointers evenly.
def listLen(head):
    count = 0
    current = head
    while current != None:
        count += 1
        current = current.next
    return count

def findMergeNodeLen(head1, head2):
    # iterate through both lists to get their length
    len1 = listLen(head1)
    len2 = listLen(head2)
    diff = abs(len1 - len2)

    # shift the pointer of larger list forward by the difference 
    # between these lengths. This ensures that you will reach the merge point
    # by moving forward both pointers evenly.
    if len1 >= len2:
        while diff > 0:
            head1 = head1.next
            diff -=1
    else:
        while diff > 0:
            head2 = head2.next
            diff -=1

    # now both pointers are the same distance away from the 
    # merge point, so increment both simultaneously until
    # we find it
    while head1 != head2:
        head1 = head1.next
        head2 = head2.next
    
    # when the nodes are equal we have found the merge point
    return head1.data

## Clever clogs answer
# https://stackoverflow.com/questions/1594061/check-if-two-linked-lists-merge-if-so-where/14956113#14956113
# Iterate through both lists evenly
# if you reach the end of one list, move to the start of the _other_ list
# guarantees you reach the intersection point at the same time
# after one loop pass
# eg if list1  = A -> D and list2 = B -> C -> D
# then paths are A D B C D and B C D A D, path length is the same
def findMergeNodeClever(head1, head2):
    p1 = head1
    p2 = head2
    while p1 != p2:
        if p1.next == None:
            p1 = head2
        else:
            p1 = p1.next
        if p2.next == None:
            p2 = head1
        else:
            p2 = p2.next
    
    return p1.data

### Below alogrithm is WRONG. ###
# Nodes are not equivalent if only the data is the same
seen = set()
def findMergeNodeSet(head1, head2):
    # Thought: keep track of visited node data in a set
    # return when either head1.data or head2.data is in that set
    if head1 != None :
        if head1.data in seen:
            return head1.data
        seen.add(head1.data)
        # move to the next node
        head1 = head1.next
    if head2 != None :
        if head2.data in seen:
            return head2.data
        seen.add(head2.data)
        head2 = head2.next
    
    # move to the next place in each
    return findMergeNode(head1, head2)
