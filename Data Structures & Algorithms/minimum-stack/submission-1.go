type MinStack struct {
	arr []int
	minArr []int
}

func Constructor() MinStack {
	return MinStack{
		arr: []int{},
		minArr: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	minVal := val
	if len(this.minArr) > 0 {
		if top := this.minArr[len(this.minArr)-1]; top < val {
			minVal = top
		}
	}
	this.minArr = append(this.minArr, minVal)
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}
