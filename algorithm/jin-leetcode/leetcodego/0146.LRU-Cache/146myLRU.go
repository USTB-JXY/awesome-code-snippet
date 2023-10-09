package leetcode

type Node struct {
	Key, Val   int
	Prev, Next *Node
}
type LRUCache struct {
	Size       int
	Head, Tail *Node
	nodeMap    map[int]*Node
}

//	func Constructor(Sizeacity int) LRUCache {
//		return LRUCache{nodeMap: make(map[int]*Node), Size: Sizeacity}
//	}
//
//	func (this *LRUCache) Get(key int) int {
//		if node, ok := this.nodeMap[key]; ok {
//			this.Remove(node)
//			this.Add(node)
//			return node.Val
//		}
//		return -1
//	}
func Constructor(Sizeacity int) LRUCache {
	return LRUCache{nodeMap: make(map[int]*Node), Size: Sizeacity}
}
func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

//	func (this *LRUCache) Put(key int, value int) {
//		if node, ok := this.nodeMap[key]; ok {
//			node.Val = value
//			this.Remove(node)
//			this.Add(node)
//			return
//		} else {
//			node = &Node{Key: key, Val: value}
//			this.nodeMap[key] = node
//			this.Add(node)
//		}
//		if len(this.nodeMap) > this.Size {
//			delete(this.nodeMap, this.Tail.Key)
//			this.Remove(this.Tail)
//		}
//	}
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodeMap[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.nodeMap[key] = node
		this.Add(node)
		if len(this.nodeMap) > this.Size {
			delete(this.nodeMap, this.Tail.Key)
			this.Remove(this.Tail)
		}
	}
}

//	func (this *LRUCache) Add(node *Node) {
//		node.Prev = nil
//		node.Next = this.Head
//		if this.Head != nil {
//			this.Head.Prev = node
//		}
//		this.Head = node
//		if this.Tail == nil {
//			this.Tail = node
//			this.Tail.Next = nil
//		}
//	}
func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.Head
	if this.Head != nil {
		this.Head.Prev = node
	}
	this.Head = node
	if this.Tail == nil {
		this.Tail = node
		this.Tail.Next = nil
	}
}

//	func (this *LRUCache) Remove(node *Node) {
//		if node == this.Head {
//			this.Head = node.Next
//			node.Next = nil
//			return
//		}
//		if node == this.Tail {
//			this.Tail = node.Prev
//			node.Prev.Next = nil
//			node.Prev = nil
//			return
//		}
//		node.Prev.Next = node.Next
//		node.Next.Prev = node.Prev
//	}
func (this *LRUCache) Remove(node *Node) {

	if node == this.Head {
		this.Head = node.Next
		node.Next = nil
		return
	}
	if node == this.Tail {
		this.Tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
