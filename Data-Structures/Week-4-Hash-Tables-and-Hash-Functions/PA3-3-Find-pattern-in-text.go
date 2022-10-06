package main

import "fmt"

func main() {
	var pattern, txt string
	fmt.Scan(&pattern, &txt)
	// q -> A prime number
	Search(txt, pattern, 101)
}

const d = int(256)

// Search searches given patterns in txt and returns the matched ones. Returns
// empty string slice if there is no match.
// https://www.geeksforgeeks.org/rabin-karp-algorithm-for-pattern-searching/
func Search(txt string, pattern string, q int) {
	M := len(pattern)
	N := len(txt)
	p := 0 // hash value for pattern
	t := 0 // hash value for txt
	var j int
	h := 1

	// The value of h would be "pow(d, M-1)%q"
	for i := 0; i < M-1; i++ {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first
	// window of text

	for i := 0; i < M; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(txt[i])) % q
	}

	// Slide the pattern over text one by one
	for i := 0; i <= N-M; i++ {

		// Check the hash values of current window of text
		// and pattern. If the hash values match then only
		// check for characters one by one
		if p == t {
			/* Check for characters one by one */
			for j = 0; j < M; j++ {
				if txt[i+j] != pattern[j] {
					break
				}
			}

			// if p == t and pat[0...M-1] = txt[i, i+1,
			// ...i+M-1]
			if j == M {
				// Pattern found at index
				fmt.Printf("%d ", i)
			}
		}

		// Calculate hash value for next window of text:
		// Remove leading digit, add trailing digit
		if i < N-M {
			t = (d*(t-int(txt[i])*h) + int(txt[i+M])) % q

			// We might get negative value of t, converting
			// it to positive
			if t < 0 {
				t += q
			}
		}
	}
}
