import "slices"
type MinStack struct {
	arr []int
}

func Constructor() MinStack {
	return MinStack{
		arr: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return slices.Min(this.arr)
}
