// simplified solutions
func longestConsecutive(nums []int) int {
	// just get the numbers in the nums array
	// removes duplicates
	inp := make(map[int]struct{}, len(nums))
	for i := range nums {
		inp[nums[i]] = struct{}{}
	}

	maxLen := 0
	// think like this, if n-1 exists then this is not the start of the sequence
	// else this is the start of the sequence & only look forward
	for k := range inp {
		// if k-1 exists then this is not the start of the sequence
		clen := 1
		if _, ok := inp[k-1]; !ok {
			_, present := inp[k+clen]
			for present {
				clen++
				_, present = inp[k+clen]
			}
		}
		maxLen = max(maxLen, clen)
	}

	return maxLen
}