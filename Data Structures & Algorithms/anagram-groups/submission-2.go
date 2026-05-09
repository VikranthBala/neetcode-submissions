// simplified
func groupAnagrams(strs []string) [][]string {
	m := map[[26]rune][]string{}
	for _, s := range strs {
		k := [26]rune{}
		for _, c := range s {
			k[c-'a']++
		}
		m[k] = append(m[k], s)
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}