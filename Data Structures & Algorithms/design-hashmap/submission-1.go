type ListNode struct {
	key, val int
	next *ListNode
}
type MyHashMap struct {
	myMap [10000]*ListNode
}

func Constructor() MyHashMap {
    myMap := [10000]*ListNode{}
	for i := range myMap {
		myMap[i] = &ListNode{key: -1, val: -1}
	}
	return MyHashMap{myMap: myMap}
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.myMap)
}

func (this *MyHashMap) Put(key int, value int) {
    idx := this.hash(key)
	cur := this.myMap[idx]
	for cur.next != nil {
		if cur.next.key == key {
			cur.next.val = value
			return
		}
		cur = cur.next
	}
	cur.next = &ListNode{key: key, val: value}
}

func (this *MyHashMap) Get(key int) int {
    idx := this.hash(key)
	cur := this.myMap[idx]
	for cur.next != nil {
		if cur.next.key == key {
			return cur.next.val
		}
		cur = cur.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    idx := this.hash(key)
	cur := this.myMap[idx]
	for cur.next != nil {
		if cur.next.key == key {
			cur.next = cur.next.next
			return
		}
		cur=cur.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */