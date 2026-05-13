func evalRPN(tokens []string) int {

	st := make([]int, len(tokens))
	top := -1

	for _, c := range tokens {
		switch c {
		case "+":
			// pop the first variable
			l := st[top]
			top--
			r := st[top]
			st[top] = l + r
		case "-":
			l := st[top]
			top--
			r := st[top]
			st[top] = r - l
		case "*":
			l := st[top]
			top--
			r := st[top]
			st[top] = l * r
		case "/":
			l := st[top]
			top--
			r := st[top]
			st[top] = r / l
		default:
			// this means the number is not an operand, so we will add it to stack
			top++
			st[top], _ = strconv.Atoi(c)
		}
	}
	return st[top]
}