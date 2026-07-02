type ListNode struct {
	key int
	next *ListNode
}

type MyHashSet struct {
	set []*ListNode
}

func Constructor() MyHashSet {
    set := make([]*ListNode, 10000)
	for i := range set {
		set[i] = &ListNode{key: 0}
	}
	return MyHashSet{set: set}
}

func (this *MyHashSet) hash(key int) int {
	return key % len(this.set)
}

func (this *MyHashSet) Add(key int) {
	cur := this.set[this.hash(key)]
	for cur.next != nil {
		if cur.next.key == key {
			return
		}
		cur = cur.next
	}
	cur.next = &ListNode{key: key}
}

func (this *MyHashSet) Remove(key int) {
	cur := this.set[this.hash(key)]

	for cur.next != nil {
		if cur.next.key == key {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
    cur := this.set[this.hash(key)]

	for cur.next != nil {
		if cur.next.key == key {
			return true
		}
		cur = cur.next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 