type MyHashMap struct {
	arr [1000001]int
}

func Constructor() MyHashMap {
    arr := [1000001]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
	return MyHashMap{arr: arr}
}

func (this *MyHashMap) Put(key int, value int) {
    this.arr[key] = value
}

func (this *MyHashMap) Get(key int) int {
    return this.arr[key]
}

func (this *MyHashMap) Remove(key int) {
	this.arr[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */