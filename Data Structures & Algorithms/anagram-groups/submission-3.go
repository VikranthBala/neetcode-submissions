func getID(s string) [26]int {
	v := [26]int{}
	for _, c := range s {
		v[c-'a']++
	}
	return v
}

func groupAnagrams(strs []string) [][]string {
	id := make(map[[26]int][]string)
	for _, str := range strs {
		strID := getID(str)
		id[strID] = append(id[strID], str)
	}
	ans := make([][]string, 0)
	for _, v := range id {
		ans = append(ans, v)
	}
	return ans
}