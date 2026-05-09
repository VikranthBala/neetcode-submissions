func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// as only lower case english letters
	lcs := make([]int, 26)
	for i := 0; i < len(s); i++ {
		// x := int()
		lcs[s[i]-'a']++
		lcs[t[i]-'a']--
	}
	for i := range 26 {
		if lcs[i] != 0 {
			return false
		}
	}
	return true
}