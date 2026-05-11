func trap(height []int) int {
	lmax := 0
	rmax := 0

	l := 0
	r := len(height) - 1

	trapped := 0

	for l < r {
		// this means left side is the deciding side
		if height[l] < height[r] {
			// check if the current height is the maxHeight or not
			if height[l] > lmax {
				lmax = height[l] // update this to the current max height
			}
			if lmax-height[l] > 0 {
				trapped += lmax - height[l]
			}
			l++
		} else {
			// check if the current height is the maxHeight or not
			if height[r] > rmax {
				rmax = height[r] // update this to the current max height
			}
			if rmax-height[r] > 0 {
				trapped += rmax - height[r]
			}
			r--
		}
	}
	return trapped
}