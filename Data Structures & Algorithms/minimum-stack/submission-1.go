type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minstack: []int{},
	}
}

func (this *MinStack) Push(val int) {

	this.stack = append(this.stack, val)

	if len(this.minstack) == 0 || val <= this.GetMin() {
		this.minstack = append(this.minstack, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()

	this.stack = this.stack[:len(this.stack)-1]

	if top == this.GetMin() {
		this.minstack =
			this.minstack[:len(this.minstack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
