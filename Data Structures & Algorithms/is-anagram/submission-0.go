func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    s1,t1 := map[byte]int{},map[byte]int{}
    for i := range s {
        s1[s[i]]++
        t1[t[i]]++
    }
    for k,v := range s1 {
        if v != t1[k] {
            return false
        }
    }
    return true
}
