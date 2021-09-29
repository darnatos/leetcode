package LRUCache

type DLinkedNode struct {
	key   int
	value int
	pre   *DLinkedNode
	post  *DLinkedNode
}

type LRUCache struct {
	count    int
	capacity int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	count := 0
	cache := make(map[int]*DLinkedNode, capacity)
	head, tail := new(DLinkedNode), new(DLinkedNode)
	head.post = tail
	tail.pre = head
	return LRUCache{
		count:    count,
		capacity: capacity,
		cache:    cache,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}
	this.moveToHead(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}

	this.count++
	if this.count > this.capacity {
		tail := this.popTail()
		delete(this.cache, tail.key)
		this.count--
	}
	node := &DLinkedNode{
		key:   key,
		value: value,
	}
	this.cache[key] = node
	this.addNode(node)
}

func (this *LRUCache) addNode(node *DLinkedNode) {
	node.pre, node.post = this.head, this.head.post
	this.head.post, this.head.post.pre = node, node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	pre, post := node.pre, node.post
	pre.post, post.pre = post, pre
	node = nil
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *DLinkedNode {
	pre := this.tail.pre
	this.removeNode(pre)
	return pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
