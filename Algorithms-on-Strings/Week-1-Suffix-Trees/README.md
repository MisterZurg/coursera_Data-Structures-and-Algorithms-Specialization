# Trie Construction - Pseudocode

The pseudocode for constructing a trie from a collection of patterns:
```
TrieConstruction(Patterns)
    Trie ← a graph consisting of a single node root
    for each string Pattern in Patterns
        currentNode ← root
        for i ← 1 to |Pattern|
            currentSymbol ← i-th symbol of Pattern
            if there is an outgoing edge from currentNode with label currentSymbol
                currentNode ← ending node of this edge
            else
                add a new node newNode to a Trie
                add a new edge  from currentNode to newNode with label currentSymbol
                currentNode ← newNode
    return Trie
```


The pseudocode for matching a collection of patterns against the text using a trie:
TODO Add
