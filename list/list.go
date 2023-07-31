package list

type Node struct {
	Val  any
	next *Node
	prev *Node
}

func (node *Node) Next() *Node {
	if node == nil {
		panic("Access to nil pointer")
	}
	return node.next
}

func (node *Node) Prev() *Node {
	if node == nil {
		panic("Access to nil pointer")
	}
	return node.prev
}

type List struct {
	head *Node
	tail *Node
	size int
}

func ListInit() *List {
	return &List{}
}

func (list *List) Size() int {
	return list.size
}

func (list *List) PushBack(val any) *Node {
	newNode := &Node{Val: val, next: nil, prev: list.tail}
	if list.head == nil {
		list.head = newNode
	} else {
		list.tail.next = newNode
	}
	list.tail = newNode
	return newNode
}

func (list *List) Back() *Node {
	return list.tail
}

func (list *List) PushFront(val any) *Node {
	newNode := &Node{Val: val, next: list.head, prev: nil}
	if list.head == nil {
		list.tail = newNode
	} else {
		list.head.prev = newNode
	}
	list.head = newNode
	return newNode
}

func (list *List) Front() *Node {
	return list.head
}

func (list *List) InsertBefore(val any, before *Node) *Node {
	if before == nil {
		panic("A value of 'before' variable must not be nil")
	}
	newNode := &Node{Val: val, prev: before.prev, next: before}
	if before.prev != nil {
		before.prev.next = newNode
	} else {
		list.head = newNode
	}
	before.prev = newNode
	return newNode
}

func (list *List) InsertAfter(val any, after *Node) *Node {
	if after == nil {
		panic("A value of 'after' variable must not be nil")
	}
	newNode := &Node{Val: val, prev: after, next: after.next}
	if after.next != nil {
		after.next.prev = newNode
	} else {
		list.tail = newNode
	}
	after.next = newNode
	return newNode
}

func (list *List) MoveAfter(node, after *Node) {
	if node == nil {
		return
	}
	if after == nil {
		panic("A value of 'after' variable must not be nil")
	}
	list.erase(node)
	node.prev = after
	if after.next != nil {
		after.next.prev = node
	} else {
		list.tail = node
	}
	node.next = after.next
	after.next = node
}

func (list *List) MoveBefore(node, before *Node) {
	if node == nil {
		return
	}
	if before == nil {
		panic("A value of 'before' variable must not be nil")
	}
	list.erase(node)
	node.next = before
	if before.prev == nil {
		list.head = node
	} else {
		before.prev.next = node
	} 
	node.prev = before.prev
	before.prev = node
}

func (list *List) erase(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}
