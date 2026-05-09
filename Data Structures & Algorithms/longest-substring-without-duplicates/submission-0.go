func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	// this is pointing to the start of the sliding window
	left := 0
	// store rune, and its, last seen position
	m := map[rune]int{}
	for i, c := range s {
		// if we see a character, as already seen, check if it is within our window
		if ind, ok := m[c]; ok && ind >= left {
			left = ind + 1
		}
		m[c] = i

		// always update the maxLen, by the length of the window
		maxLen = max(maxLen, i-left+1)
	}
	return maxLen
}