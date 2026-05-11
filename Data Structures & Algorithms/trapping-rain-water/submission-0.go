func trap(height []int) int {
	lp := make([]int, len(height))
	rp := make([]int, len(height))
	maxH := 0
	for i := range height {
		lp[i] = max(maxH, height[i])
		maxH = max(maxH, height[i])
	}
	maxH = 0
	for i := len(height) - 1; i >= 0; i-- {
		rp[i] = max(maxH, height[i])
		maxH = max(maxH, height[i])
	}

	trapped := 0
	for i := range height {
		if min(lp[i], rp[i])-height[i] > 0 {
			trapped += min(lp[i], rp[i]) - height[i]
		}
	}
	return trapped
}