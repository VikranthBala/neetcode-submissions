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
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}