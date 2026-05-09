func topKFrequent(nums []int, k int) []int {
	lnums := len(nums)

	// calculate the frequencies of each number
	fmap := make(map[int]int, lnums)
	for _, n := range nums {
		fmap[n]++
	}

	minfq := math.MaxInt
	maxfq := 0

	// get the frequency range, as we dont want to keep the buckets for lnums
	for _, f := range fmap {
		minfq = min(minfq, f)
		maxfq = max(maxfq, f)
	}

	// bucket array length, as this needs to only have frequencies, between min, max
	lbucket := maxfq - minfq + 1
	bmap := make([][]int, lbucket)

	for n, f := range fmap {
		// index must be adjusted, as we have from minfq to maxfq
		index := f - minfq
		bmap[index] = append(bmap[index], n)
	}

	//do a reverse lookup
	res := make([]int, 0, k)
	for i := len(bmap) - 1; i >= 0 && k > 0; i-- {
		for _, n := range bmap[i] {
			res = append(res, n)
			k = k - 1

			if k == 0 {
				return res
			}
		}
	}
	return res
}