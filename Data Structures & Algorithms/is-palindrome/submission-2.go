func isAlphaNumeric(r byte) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	if r >= 97 && r <= 122 {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlphaNumeric(s[l]) {
			l++
			continue
		}
		if !isAlphaNumeric(s[r]) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}