func productExceptSelf(nums []int) []int {
	lnums := len(nums)
	lp := make([]int, lnums)
	rp := make([]int, lnums)

	lp[0] = 1
	for i := 1; i < lnums; i++ {
		lp[i] = lp[i-1] * nums[i-1]
	}

	rp[lnums-1] = 1
	for i := lnums - 2; i >= 0; i-- {
		rp[i] = rp[i+1] * nums[i+1]
	}

	res := make([]int, lnums)
	for i := range lnums {
		res[i] = lp[i] * rp[i]
	}

	return res
}