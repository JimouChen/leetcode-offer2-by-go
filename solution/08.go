package solution

type CQueue struct {
	stack1, stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	len1, len2 := len(this.stack1), len(this.stack2)
	if len1 == 0 && len2 == 0 {
		return -1
	} else if len2 != 0 {
		res := this.stack2[len2-1]
		this.stack2 = append(this.stack2[:len2-1])
		return res
	} else {
		for len(this.stack1) > 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = append(this.stack1[:len(this.stack1)-1])
		}
		res := this.stack2[len(this.stack2)-1]
		this.stack2 = append(this.stack2[:len(this.stack2)-1])
		return res
	}

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
