type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	//可能不合法
	if len(this.out) == 0 {
		return -1
	}
	x := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return x
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */