func twoSum(nums []int, target int) []int {
	ind := make(map[int]int, len(nums))
	for i := range len(nums) {
		if v, ok := ind[target-nums[i]]; ok {
			return []int{v, i}
		}
        ind[nums[i]] = i
	}
	return []int{}
}