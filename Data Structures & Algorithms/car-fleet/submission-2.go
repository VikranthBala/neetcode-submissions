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

// simpler approach
func carFleet(target int, position []int, speed []int) int {

	// lets sort the positions first
	sort.Sort(carSorter{positions: position, speed: speed})

	// loop through the positions
	fleets := 0
	slowestTime := 0.0

	for i := len(position) - 1; i >= 0; i-- {

		time :=
			float64(target-position[i]) /
				float64(speed[i])

		if time > slowestTime {
			fleets++
			slowestTime = time
		}
	}

	return fleets
}