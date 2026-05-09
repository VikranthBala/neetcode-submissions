func topKFrequent(nums []int, k int) []int {
	// field of interest
	fmap := make(map[int]int, 2000)
	for _, n := range nums {
		fmap[n]++
	}
	bmap := make([][]int, len(nums)+1)
	for n, f := range fmap {
		bmap[f] = append(bmap[f], n)
	}

	//do a reverse lookup
	res := []int{}
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