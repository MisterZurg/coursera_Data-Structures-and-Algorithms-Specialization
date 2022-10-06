# Dynamic Arrays and Amortized Analysis
## Question 1
> What is the size of the array needed to store integer keys with up to 12 digits using direct addressing?
```
❌ 12  It is much more.
✔ 10^12
2^12 # 1000000000 has only 10 digits, but 2^12 = 40962, so there is no cell with index 1000000000 in the array of size 2^12
```
## Question 2
> What is the maximum possible chain length for a hash function h(x) = x mod 1000 used with a hash table of size 1000 for a universe of all integers with at most 12 digits?
```
❌ 1
✔  10^9 # When the values of the last 33 digits are fixed, there are 10^9
❌ 10^12 numbers with at most 12 digits.
```
## Question 3
> You want to hash integers from 0 up to 1000000. What can be a good choice of pp for the universal family?
```
✔  1000003    # This is a prime number bigger than 10000001000000.
❌ 1000002 # p must be prime, but 1000002 = 2 × 500001.
❌ 999997
```

## Question 4
> How can one build a universal family of hash functions for integers between -1000000−1000000 (minus one million) and 10000001000000 (one million)?
```
✔ First, add 1000000 to each integer and get the range of integers between 0 and 2000000. Then use the universal family for integers with p = 2000003.
❌ First, add 1000000 to each integer. Then use the universal family for integers with p = 1000003.
❌ Take the universal family for integers with p = 1000003.
```