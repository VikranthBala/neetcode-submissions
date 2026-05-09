func hasDuplicate(nums []int) bool {
    p := map[int]bool{}
    for _, i := range nums {
        if p[i] {
            return true
        }
        p[i] = true
    }
    return false
}
