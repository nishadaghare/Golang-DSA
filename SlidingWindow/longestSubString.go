package main

import "fmt"

func longestSubstringKDistinct(s string, k int) int {

	// If string is empty OR k is 0
	if len(s) == 0 || k == 0 {
		return 0
	}

	// Store character counts
	charCount := make(map[byte]int)

	// Start of window
	windowStart := 0

	// Store maximum length found
	maxLength := 0

	// Move windowEnd one by one
	for windowEnd := 0; windowEnd < len(s); windowEnd++ {

		// Current character
		currentChar := s[windowEnd]

		// Add character into map
		charCount[currentChar]++

		// If distinct characters become more than k
		// shrink window from left side
		for len(charCount) > k {

			// Left character of window
			leftChar := s[windowStart]

			// Reduce count
			charCount[leftChar]--

			// Remove completely if count becomes 0
			if charCount[leftChar] == 0 {
				delete(charCount, leftChar)
			}

			// Move window start forward
			windowStart++
		}

		// Current valid window size
		currentWindowLength := windowEnd - windowStart + 1

		// Update maximum length
		if currentWindowLength > maxLength {
			maxLength = currentWindowLength
		}
	}

	return maxLength
}

func main() {

	// Example 1
	s1 := "eceba"
	k1 := 2

	answer1 := longestSubstringKDistinct(s1, k1)

	fmt.Println("String :", s1)
	fmt.Println("K      :", k1)
	fmt.Println("Answer :", answer1)

	fmt.Println()

	// Example 2
	s2 := "aaabcbbbbb"
	k2 := 2

	answer2 := longestSubstringKDistinct(s2, k2)

	fmt.Println("String :", s2)
	fmt.Println("K      :", k2)
	fmt.Println("Answer :", answer2)
}

// String: e c e b a
//          ↑
//       windowEnd moves →

// Step 1:
// [e]

// Step 2:
// [e c]

// Step 3:
// [e c e]

// Step 4:
// [e c e b] ❌

// Shrink:

// [c e b] ❌

// Shrink again:

// [e b] ✅

// Step 5:

// [e b a] ❌

// Shrink:

// [b a] ✅
