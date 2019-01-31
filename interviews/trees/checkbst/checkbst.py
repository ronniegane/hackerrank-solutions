""" Node is defined as
class node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None
"""


"""
checkBST should return true if the tree passed to it is a binary search tree,
that is, for each node N, all values in the subtree N.left are
<= N, and all values in the subtree n.right are >= N.
"""
def checkBST(root):
    return inOrderWalk(root)

def inOrderWalk(root):
    # possible approach: do an in-order walk to build a sorted list of values
    # as you add to the list, check that it remains sorted
    # i.e. the element you add is >= the last element
    
    # we are assured that our data is >= 0, so passing in -1 as 
    # the first seen value should be OK. -infinity would be better
    # or we could have a check if data < 0 then false
    valid, _ = walk(root, -1)
    return valid

# walk does an in-order walk through the tree and checks
# that the values it encounters maintain order 
def walk(node, last):
    # left first if it exists
    if node.left != None:
        valid, last = walk(node.left, last)
        if not valid:
            return False, None
    # then this node
    # we check for equality because every node should be distinc 
    if node.data <= last:
        return False, None
    last = node.data
    # then right if it exists
    if node.right != None:
        return walk(node.right, last)
    return True, last




def recurseMaxMin(root):
    # another approach: recurse down the tree with a helper function
    # isBST(root, max, min) which passes down allowed min and max values at each step
    return False
