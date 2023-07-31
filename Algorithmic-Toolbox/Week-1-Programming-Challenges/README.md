# Algorithmic Toolbox at Coursera: Programming Challenges (Week 1)
## 1 Sum of Two Digits
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> Compute the sum of two single digit numbers.
> 
> **Input:** Two single digit numbers.
>
> **Output:** The sum of these numbers.

We start from this ridiculously simple problem to show you the
pipeline of reading the problem statement, designing an algorithm, implementing
it, testing and debugging your program, and submitting it to
the grading system.

Input format. Integers a and b on the same line (separated by a space).

Output format. The sum of a and b.

Constraints. $0 \le a;b \le 9$.

**Sample.**

Input:
```
9 7
```
Output:
```
16
```

> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> Implementing an Algorithm
> ```
> SumOfTwoDigits(a, b):
> return a + b
> ```

## 2 Maximum Pairwise Product
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> Find the maximum product of two distinct numbers in a sequence of non-negative integers.
> 
> **Input:** A sequence of non-negative integers.
> 
> **Output:** The maximum value that can be obtained by multiplying two different elements from the sequence.

**Input format.** The first line contains an integer n. The next line contains n non-negative integers $a_1; \dots ;a_n$ (separated by spaces).

**Output format.** The maximum pairwise product.

**Constraints.** $2 \le n \le 2 * 10^5; 0 \le a_1; \dots ;a_n \le 2 * 10^5$.

**Sample 1.**

Input:
```
3
1 2 3
```
Output:
```
6
```
**Sample 2.**

Input:
```
10
7 5 14 2 8 8 10 1 2 3
```
Output:
```
140
```

> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/example.svg">
>   <img alt="Example" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/example.svg">
> </picture><br>
>
> Fast and Correct algorithm
> ```
> MaxPairwiseProductFast(A[1 : : :n]):
> index_1 ← 1
> for i from 2 to n:
>   if A[i] > A[index1]:
>       index_1 ← i
> if index_1 = 1:
>   index_2 ← 2
> else:
>   index_2 ← 1
> for i from 1 to n:
>   if A[i] != A[index_1] and A[i] > A[index_2]:
>       index_2 ← i
> return A[index1] * A[index2]
> ```
> A More Compact Algorithm
> ```
> MaxPairwiseProductBySorting(A[1 : : :n]):
> Sort(A)
> return A[n - 1] * A[n]
> ```
