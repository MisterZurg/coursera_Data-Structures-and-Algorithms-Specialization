# Dynamic Arrays and Amortized Analysis
## Question 1
> Consider the following program:
> ```py
> for i from 1 to 12:
>   MakeSet(i)
> Union(2, 10)
> Union(7, 5)
> Union(6, 1)
> Union(3, 4)
> Union(5, 11)
> Union(7, 8)
> Union(7, 3)
> Union(12, 2)
> Union(9, 6)
> print(Find(6))
> print(Find(3))
> print(Find(11))
> print(Find(9))
> ```
> Assume that the disjoint sets data structure is implemented as an array `smallest[1 ... 12]:smallest[i]` is equal to the smallest element in the set containing `i`.
>
> What is the output of the following program? As an answer, enter four integers separated by spaces.


```
1 3 3 1
```

## Question 2
> Consider the following program:
> ```py
> for i from 1 to 12:
>     MakeSet(i)
> Union(2, 10)
> Union(7, 5)
> Union(6, 1)
> Union(3, 4)
> Union(5, 11)
> Union(7, 8)
> Union(7, 3)
> Union(12, 2)
> Union(9, 6)
> ```
> Assume that the disjoint sets data structure is implemented as disjoint trees with union by rank heuristic.
>
> Compute the product of the heights of the resulting trees after executing the code. For example, for a forest consisting of four trees of height 1, 2, 3, 1 the answer would be 6. (Recall that the height of a tree is the number of edges on a longest path from the root to a leaf. In particular, the height of a tree consisting of just one node is equal to 0.)


```
2 # There will be 3 trees of height 1, 1, and 2.
```

## Question 3
> Consider the following program:
> ```py
> for i from 1 to n:
>   MakeSet(i)
> for i from 1 to n-1:
>   Union(i, i + 1)
> ```
> Assume that the disjoint sets data structure is implemented as disjoint trees with union by rank heuristic.
>
> What is the number of trees in the forest and the maximum height of a tree in this forest after executing this code? (Recall that the height of a tree is the number of edges on a longest path from the root to a leaf. In particular, the height of a tree consisting of just one node is equal to 0.)

```
✔ One tree of height 1.
❌ n/2 trees, the maximum height is 2.
❌ Two trees, both of height 1.
❌ log_2(n) trees, the maximum height is 1.
❌ n trees, the maximum height is 1.
❌ One tree of height log_2(n).
```

## Question 4
> Consider the following program:
> ```py
> for i from 1 to 60:
>   MakeSet(i)
> for i from 1 to 30:
>   Union(i, 2*i)
> for i from 1 to 20:
>   Union(i, 3*i)
> for i from 1 to 12:
>   Union(i, 5*i)
> for i from 1 to 60:
>   Find(i)
> ```
> Assume that the disjoint sets data structure is implemented as disjoint trees with union by rank heuristic and with path compression heuristic.
>
> Compute the maximum height of a tree in the resulting forest. (Recall that the height of a tree is the number of edges on a longest path from the root to a leaf. In particular, the height of a tree consisting of just one node is equal to 0.)

```
1 
# There is at least one tree of height 1 in the forest. 
# Also, all trees have height at most 1, since the last for-loop calls Find() for all 60 elements. 
# Since path compression is used, each non-root node will be attached directly to the corresponding root in this loop,
# and hence all the trees will have height at most 1.
```
