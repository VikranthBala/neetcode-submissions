func dailyTemperatures(temperatures []int) []int {

	lt := len(temperatures)
	// result
	result := make([]int, lt)

	// simple stack to maintain indices
	inds := make([]int, lt)
	ltop := -1

	// now loop over temperatures

	for i, temp := range temperatures {
		// if ltops is -1 then that means nothing in the stack, so directly add the current temp
		if ltop == -1 {
			ltop++
			inds[ltop] = i
		} else {
			// this means that some temperature is already in the stack
			// check if the top temp in stack is greater than or less than current temp

			// if the top temp of the stack is greater than current temp,
			// add it to the stack, as its not a warmer day
			if temperatures[inds[ltop]] >= temp {
				ltop++
				inds[ltop] = i
			} else {
				// i.e, our stack is lesser than temp, i.e, a warmer day
				topTemp := temperatures[inds[ltop]]
				for topTemp < temp {
					// start removing elements from the stack if the current temp is warmer

					// start popping
					result[inds[ltop]] = i - inds[ltop]
					ltop--
					// this means no more elements present in stack, so exit from the for loop
					if ltop == -1 {
						break
					}
					topTemp = temperatures[inds[ltop]]
				}

				// once all the elements are removed we need to add the curren temp
				ltop++
				inds[ltop] = i
			}
		}
	}
	return result
}