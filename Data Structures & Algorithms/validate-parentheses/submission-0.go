func isValid(s string) bool {
	i := -1
	st := make([]rune, len(s))

	lookup := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	for _, c := range s {
		if i < 0 {
			i++
			st[i] = c
			continue
		} else {
			if ex, ok := lookup[c]; ok {
				if st[i] != ex {
					return false
				}
				i--
				continue
			}
			i++
			st[i] = c
		}
	}
	return i == -1
}