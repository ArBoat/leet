type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{
		q: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	n := len(this.q)
	this.q = append(this.q, x)
	for ; n > 0; n-- {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

func (this *MyStack) Pop() int {
	res := this.q[0]
	this.q = this.q[1:]
	return res
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */