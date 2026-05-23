package main

import "fmt"

func longestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	charCount := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		// shrink window if more than k distinct chars
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		// update max length
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	s := "eceba"

	k := 2
	fmt.Println(longestSubstringKDistinct(s, k)) // Output: 3

	fmt.Println(longestSubstringKDistinct("aaabcbbbbb", 2)) // Output: 7 ("bcbbbbb")
}
