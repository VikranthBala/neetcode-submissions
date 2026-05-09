func twoSum(nums []int, target int) []int {
    dm := map[int]int{}
    for i, n := range nums {
        if v, ok := dm[n];ok {
            return []int{v,i}
        }
        dm[target-n]=i
    }
    return []int{}
}
