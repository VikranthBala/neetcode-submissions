func getID(s string) [26]byte {
	var v [26]byte
	for i := 0; i < len(s); i++ {
		v[s[i]-'a']++
	}
	return v
}

func groupAnagrams(strs []string) [][]string {
	id := make(map[[26]byte][]string)
	for _, str := range strs {
		strID := getID(str)
		id[strID] = append(id[strID], str)
	}
	ans := make([][]string, 0, len(id))
	for _, v := range id {
		ans = append(ans, v)
	}
	return ans
}