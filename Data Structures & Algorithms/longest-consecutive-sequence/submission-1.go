func longestConsecutive(nums []int) int {
	kmap := make(map[int]struct{}, len(nums))
	seen := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		kmap[num] = struct{}{}
	}

	maxCon := 0
	for k := range kmap {
		clen := 0
		if _, ok := seen[k]; ok {
			fmt.Println("seen: ", k)
			continue
		}

		// that particular number -> ex: 2
		clen++
		// mark it as seen
		seen[k] = struct{}{}

		// now look at left & right of that number, starting with diff 1
		ldiff, rdiff := 1, 1
		// ok := true

		_, lok := kmap[k-ldiff]
		_, rok := kmap[k+rdiff]

		for lok {
			if _, ok := seen[k-ldiff]; !ok {
				seen[k-ldiff] = struct{}{}
				clen++
				ldiff--
				_, lok = kmap[k-ldiff]
			}
			lok = false
		}

		for rok {
			seen[k+rdiff] = struct{}{}
			clen++
			rdiff++
			_, rok = kmap[k+rdiff]
		}

		maxCon = max(maxCon, clen)
	}
	return maxCon
}