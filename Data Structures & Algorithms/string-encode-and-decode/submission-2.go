type Solution struct{}

func (s *Solution) Encode(strs []string) string {
// as we just need to join the strings
	// and the input contains all possible ascii character, cannot use generic values as joiner

	// lets do a length-based encoding
	var encoded strings.Builder
	for _, str := range strs {
		encoded.WriteString(strconv.Itoa(len(str)))
		encoded.WriteByte('#')
		encoded.WriteString(str)
	}
	return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
	decoded := []string{}
	ind := 0
	slen := 0
	for ind < len(encoded) {
		for i := ind; i < len(encoded); i++ {
			if encoded[i] == '#' {
				// get the length from the start
				slen, _ = strconv.Atoi(string(encoded[ind:i]))

				// move ind by i+1, we need to go to the actual string
				ind = i + 1
				break
			}
		}
		// append from ind, to ind+slen to the end
		decoded = append(decoded, encoded[ind:ind+slen])
		ind = ind + slen
	}
	return decoded
}

