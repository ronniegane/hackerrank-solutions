"""
Huffman coding for lossless compression
each character is assigned a code
more frequently-appearing characters get shorter codes,
so the length of the coded message is minimised.

only leaves of the tree contain a letter and a frequency count.
internal nodes contail null, and the sum of frequencies of it's children
ie N.freq = N.left.freq + N.right.freq

having no internal nodes with chars ensures that no code is a prefix
of any other code (preventing ambiguity)
"""

"""class Node:
    def __init__(self, freq,data):
        self.freq= freq
        self.data=data
        self.left = None
        self.right = None
"""        

# Enter your code here. Read input from STDIN. Print output to STDOUT
# root is the root of a huffman tree
# s is the input string
def decodeHuff(root, s):
	# process string left to right
    current = root
    output = ""
    for c in s:
        # go left or right depending on whether c is 0 or 1
        if c == '0':
            current = current.left
        else:
            current = current.right
        # if this is a leaf node, store the character and reset to the root
        if current.left == None and current.right == None:
            output += current.data
            current = root
    
    print(output)
