func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// as we just need count
	counts := map[byte]int{}
	for i := range s {
		counts[s[i]]++
		counts[t[i]]--
	}
	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}

