type MyHashSet struct {
	arr [1000001]bool
}

func Constructor() MyHashSet {
    return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
    this.arr[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.arr[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.arr[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 