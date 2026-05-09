func maxArea(heights []int) int {
	ma := -1
	l := 0
	r := len(heights) - 1

	for l < r {
		// find mininum
		a := r - l
		if heights[r] < heights[l] {
			a *= heights[r]
			r--
		} else {
			a *= heights[l]
			l++
		}
		ma = max(a, ma)
	}
	return ma
}