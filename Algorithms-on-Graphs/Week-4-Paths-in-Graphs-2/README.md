# Algorithms on Graphs at Coursera: Paths in Graphs Programming Challenges (Week 4)
## 1 Computing the Minimum Cost of a Flight
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> Task. Given an directed graph with positive edge weights and with $n$ vertices and $m$ edges as well as two vertices $u$ and $v$, compute the weight of a shortest path between $u$ and $v$ (that is, the minimum total weight of a path from $u$ to $v$).

**Input Format.** A graph is given in the standard format. The next line contains two vertices $u$ and $v$.

**Constraints.** $1 ≤ n ≤ 10^4, 0 ≤ m ≤ 10^5, u \ne v, 1 ≤ u, v ≤ n$, edge weights are non-negative integers not exceeding $10^8$.

**Output Format.** Output the minimum weight of a path from $u$ to $v$, or $−1$ if there is no path.

**Sample 1.**

Input1:
```
4 4
1 2 1
4 1 2
2 3 2
1 3 5
1 3
```
Output:
```
3
```
**Sample 2.**

Input:
```
5 9
1 2 4
1 3 2
2 3 2
3 2 1
2 4 2
3 5 4
5 4 1
2 5 3
3 4 4
1 5
```
Output:
```
6
```
**Sample 3.**

Input:
```
3 3
1 2 7
1 3 5
2 3 2
3 2
```
Output:
```
-1
```
## 2 Detecting Anomalies in Currency Exchange Rates
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> Task. Given a directed graph with possibly negative edge weights and with $n$ vertices and $m$ edges, check whether it contains a cycle of negative weight.

**Input Format.** A graph is given in the standard format.

**Constraints.** $1 ≤ n ≤ 10^3, 0 ≤ m ≤ 10^4$, edge weights are integers of absolute value at most $10^3$.

**Output Format.** Output $1$ if the graph contains a cycle of negative weight and $0$ otherwise.

**Sample.**

Input:
```
4 4
1 2 -5
4 1 2
2 3 2
3 1 1
```
Output:
```
1
```
## 3 Exchanging Money Optimally
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> Task. Given a directed graph with possibly negative edge weights and with $n$ vertices and $m$ edges as well as its vertex $s$, compute the length of shortest paths from $s$ to all other vertices of the graph.

**Input Format.** A graph is given in the standard format.

**Constraints.** $1 ≤ n ≤ 10^3, 0 ≤ m ≤ 10^4, 1 ≤ s ≤ n$, edge weights are integers of absolute value at most $10^9$.

**Output Format.** For all vertices $i$ from $1$ to $n$ output the following on a separate line:
- $“*”$, if there is no path from $s$ to $u$;
- $“-”$, if there is a path from $s$ to $u$, but there is no shortest path from $s$ to $u$ (that is, the distance from $s$ to $u$ is $-\inf$);
- the length of a shortest path otherwise.

**Sample 1.**

Input:
```
6 7
1 2 10
2 3 5
1 3 100
3 5 7
5 4 10
4 3 -18
6 1 -1
1
```
Output:
```
0
10
-
-
-
*
```
**Sample 2.**

Input:
```
5 4
1 2 1
4 1 2
2 3 2
3 1 -5
4
```
Output:
```
-
-
-
0
*
```