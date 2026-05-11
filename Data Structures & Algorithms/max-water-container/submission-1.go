func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	area := 0
	for l < r {
		area = max(area, min(heights[l], heights[r])*(r-l))
		if heights[l] < heights[r] {
			l++
			continue
		}
		r--
	}
	return area
}