type wordSig struct {
	a int
	b int
	c int
	d int
	e int
	f int
	g int
	h int
	i int
	j int
	k int
	l int
	m int
	n int
	o int
	p int
	q int
	r int
	s int
	t int
	u int
	v int
	w int
	x int
	y int
	z int
}

// brute force
func groupAnagrams(strs []string) [][]string {

	//word signature
	ws := map[wordSig][]string{}

	// for any word which has same letters will have same sum
	for i := 0; i < len(strs); i++ {
		st := wordSig{}
		for j, _ := range strs[i] {
			switch strs[i][j] {
			case 'a':
				st.a++
			case 'b':
				st.b++
			case 'c':
				st.c++
			case 'd':
				st.d++
			case 'e':
				st.e++
			case 'f':
				st.f++
			case 'g':
				st.g++
			case 'h':
				st.h++
			case 'i':
				st.i++
			case 'j':
				st.j++
			case 'k':
				st.k++
			case 'l':
				st.l++
			case 'm':
				st.m++
			case 'n':
				st.n++
			case 'o':
				st.o++
			case 'p':
				st.p++
			case 'q':
				st.q++
			case 'r':
				st.r++
			case 's':
				st.s++
			case 't':
				st.t++
			case 'u':
				st.u++
			case 'v':
				st.v++
			case 'w':
				st.w++
			case 'x':
				st.x++
			case 'y':
				st.y++
			case 'z':
				st.z++
			}
		}
		ws[st] = append(ws[st], strs[i])
	}
	v := [][]string{}
	for _, sv := range ws {
		v = append(v, sv)
	}
	return v
}