package LFUCache

type Node struct {
	key   int
	value int
	freq  int
	prev  *Node
	next  *Node
}

type DLinkedList struct {
	sentinel *Node
	size     int
}

type LinkedList interface {
	Append(node Node)
	Pop(node Node)
}

func (this *DLinkedList) Append(node *Node) {
	node.next = this.sentinel.next
	node.prev = this.sentinel
	node.next.prev = node
	this.sentinel.next = node
	this.size++
}

func (this *DLinkedList) Pop(node *Node) *Node {
	if this.size == 0 {
		return node
	}
	if node == nil {
		node = this.sentinel.prev
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	this.size--

	return node
}

func newDLinkedList() *DLinkedList {
	sentinel := new(Node)
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &DLinkedList{sentinel: sentinel, size: 0}
}

type LFUCache struct {
	node     map[int]*Node
	freq     map[int]*DLinkedList
	minFreq  int
	capacity int
}

func Constructor(capacity int) LFUCache {
	node := make(map[int]*Node, capacity+1)
	freq := make(map[int]*DLinkedList)
	freq[1] = newDLinkedList()
	cache := LFUCache{
		node:     node,
		freq:     freq,
		minFreq:  0,
		capacity: capacity,
	}
	return cache
}

func (this *LFUCache) update(node *Node) {
	this.freq[node.freq].Pop(node)

	if node.freq == this.minFreq && this.freq[node.freq].size == 0 {
		this.minFreq++
	}
	node.freq++
	if _, ok := this.freq[node.freq]; !ok {
		this.freq[node.freq] = newDLinkedList()
	}
	this.freq[node.freq].Append(node)
}
func (this *LFUCache) Get(key int) int {
	node, ok := this.node[key]
	if !ok {
		return -1
	}
	this.update(node)
	return node.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	node, ok := this.node[key]
	if ok {
		(*node).value = value
		this.node[key] = node
		this.Get(key)
		return
	}
	if len(this.node) >= this.capacity {
		node := this.freq[this.minFreq].Pop(nil)
		delete(this.node, node.key)
	}
	node = &Node{key: key, value: value, freq: 1}
	this.node[key] = node
	this.freq[1].Append(node)
	this.minFreq = 1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
