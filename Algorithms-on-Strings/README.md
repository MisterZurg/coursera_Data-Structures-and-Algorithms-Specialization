# Algorithms on Strings
## Syllabus
**Suffix Trees**
> [!NOTE]
> How would you search for a longest repeat in a string in LINEAR time? In 1973, Peter Weiner came up with a surprising solution that was based on suffix trees, the key data structure in pattern matching. Computer scientists were so impressed with his algorithm that they called it the Algorithm of the Year. In this lesson, we will explore some key ideas for pattern matching that will - through a series of trials and errors - bring us to suffix trees.

**Burrow-Wheeler Transform and Suffix Arrays**
> [!NOTE]
> Although EXACT pattern matching with suffix trees is fast, it is not clear how to use suffix trees for APPROXIMATE pattern matching. In 1994, Michael Burrows and David Wheeler invented an ingenious algorithm for text compression that is now known as Burrows-Wheeler Transform. They knew nothing about genomics, and they could not have imagined that 15 years later their algorithm will become the workhorse of biologists searching for genomic mutations. But what text compression has to do with pattern matching??? In this lesson you will learn that the fate of an algorithm is often hard to predict – its applications may appear in a field that has nothing to do with the original plan of its inventors.

**Knuth-Morris-Pratt Algorithm**
> [!NOTE]
> Congratulations, you have now learned the key pattern matching concepts: tries, suffix trees, suffix arrays and even the Burrows-Wheeler transform! However, some of the results Pavel mentioned remain mysterious: e.g., how can we perform exact pattern matching in $O(|Text|)$ time rather than in $O(|Text|*|Pattern|)$ time as in the naïve brute force algorithm? How can it be that matching a 1000-nucleotide pattern against the human genome is nearly as fast as matching a 3-nucleotide pattern??? Also, even though Pavel showed how to quickly construct the suffix array given the suffix tree, he has not revealed the magic behind the fast algorithms for the suffix tree construction!In this module, Miсhael will address some algorithmic challenges that Pavel tried to hide from you :) such as the Knuth-Morris-Pratt algorithm for exact pattern matching and more efficient algorithms for suffix tree and suffix array construction.

**Constructing Suffix Arrays and Suffix Trees**
> [!NOTE]
> In this module we continue studying algorithmic challenges of the string algorithms. You will learn an $O(n log n)$ algorithm for suffix array construction and a linear time algorithm for construction of suffix tree from a suffix array. You will also implement these algorithms and the Knuth-Morris-Pratt algorithm in the last Programming Assignment in this course.