type carSorter struct {
	positions []int
	speed     []int
}

func (c carSorter) Len() int {
	return len(c.positions)
}

func (c carSorter) Less(i, j int) bool {
	return c.positions[i] < c.positions[j]
}

func (c carSorter) Swap(i, j int) {
	c.positions[i], c.positions[j] = c.positions[j], c.positions[i]
	c.speed[i], c.speed[j] = c.speed[j], c.speed[i]
}

func carFleet(target int, position []int, speed []int) int {

	// lets sort the positions first
	sort.Sort(carSorter{positions: position, speed: speed})

	fleetsAt := make([]int, len(position))
	top := -1

	lp := len(position)
	// loop through the positions
	for i := range lp {
		for top >= 0 &&
			(float64(target-position[i])/float64(speed[i])) >= (float64(target-position[fleetsAt[top]])/float64(speed[fleetsAt[top]])) {
			top--
		}
		top++
		fleetsAt[top] = i
	}
	return top + 1
}