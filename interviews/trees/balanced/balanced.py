"""
given a tree of nodes, can we split it into three
trees with identical sums, by doing the following:
1) add a single node with value of our choosing
2) cut two edges

Possible approach: work out tree sums for each node
in the tree. we require a certain pattern 

could brute-force cutting all combinations of 2 edges
and working out the sum of the resulting trees, to see
if any of them have two identical sums (then add node to 
balance third one)

no limit listed on edges, but up to 2000 nodes and each node has
one upward-travelling edge besides the root node, so that would be 1999 edges.

2000 choose 2 edges (order unimportant) is
2000! / (2!*1998!) = 2000 * 1999 / 2 = 1999000 

n ~= 2*10e6 is pretty large for brute force, especially when summing
will take appreciable time

input is a list of node values and a list of edges,
so you may have to build the tree structure first.
"""
