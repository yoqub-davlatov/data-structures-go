package list

// Node is a building block of a list.
type Node struct {
	// Val carries an object that current node holds.
	Val any
	// next points to the next object of the list
	// that is located right after the current node.
	next *Node
	// prev points to the previous object of the list
	// that is located right before the current node.
	prev *Node
}

// Returns a pointer to the next element in the list.
// If node is the tail, Next returns nil.
// Panics when node is nil.
func (node *Node) Next() *Node {
	if node == nil {
		panic("Access to nil pointer")
	}
	return node.next
}

// Returns a pointer to the previous element in the list.
// If node is the head, Prev returns nil.
// Panics when node is nil.
func (node *Node) Prev() *Node {
	if node == nil {
		panic("Access to nil pointer")
	}
	return node.prev
}

// List is a struct that represents doubly linked list.
type List struct {
	head *Node
	tail *Node
	size int
}

// Initializes a list.
func ListInit() *List {
	return &List{}
}

// Returns the current size of the list in O(1),
// i.e. number of elements in the list.
func (list *List) Size() int {
	return list.size
}

// Adds a new element in the end of the list in O(1) time.
// Returns the address of the block in the memory that was
// allocated for the object.
func (list *List) PushBack(val any) *Node {
	list.size++
	newNode := &Node{Val: val, next: nil, prev: list.tail}
	if list.head == nil {
		list.head = newNode
	} else {
		list.tail.next = newNode
	}
	list.tail = newNode
	return newNode
}

// Returns a pointer to the tail of the list.
func (list *List) Back() *Node {
	return list.tail
}

// Adds a new element in the head of the list in O(1) time.
// Returns the address of the block of memory that was allocated
// for the new object.
func (list *List) PushFront(val any) *Node {
	list.size++
	newNode := &Node{Val: val, next: list.head, prev: nil}
	if list.head == nil {
		list.tail = newNode
	} else {
		list.head.prev = newNode
	}
	list.head = newNode
	return newNode
}

// Returns a pointer to the head of the list.
func (list *List) Front() *Node {
	return list.head
}

// Inserts a new element before node.
// Panics when node is nil. Returns a pointer to a new
// allocated block of memory for the object
func (list *List) InsertBefore(val any, node *Node) *Node {
	list.size++
	if node == nil {
		panic("A value of 'node' variable must not be nil")
	}
	newNode := &Node{Val: val, prev: node.prev, next: node}
	if node.prev != nil {
		node.prev.next = newNode
	} else {
		list.head = newNode
	}
	node.prev = newNode
	return newNode
}

// Inserts a new element after node.
// Panics when node is nil. Returns a pointer to a new
// allocated block of memory for the object
func (list *List) InsertAfter(val any, node *Node) *Node {
	list.size++
	if node == nil {
		panic("A value of 'node' variable must not be nil")
	}
	newNode := &Node{Val: val, prev: node, next: node.next}
	if node.next != nil {
		node.next.prev = newNode
	} else {
		list.tail = newNode
	}
	node.next = newNode
	return newNode
}

// Places node after the node.
// Panics when mark is nil.
func (list *List) MoveAfter(node, mark *Node) {
	if node == nil {
		return
	}
	if mark == nil {
		panic("A value of 'mark' variable must not be nil")
	}
	if node == mark {
		return
	}
	list.erase(node)
	node.prev = mark
	if mark.next != nil {
		mark.next.prev = node
	} else {
		list.tail = node
	}
	node.next = mark.next
	mark.next = node
}

// Places node before the node.
// Panics when mark is nil.
func (list *List) MoveBefore(node, mark *Node) {
	if node == nil {
		return
	}
	if mark == nil {
		panic("A value of 'mark' variable must not be nil")
	}
	if node == mark {
		return
	}
	list.erase(node)
	node.next = mark
	if mark.prev == nil {
		list.head = node
	} else {
		mark.prev.next = node
	}
	node.prev = mark.prev
	mark.prev = node
}

// Moves node to the end of the list, thus making 
// it list's tail. Panics when node is nil
func (list *List) MoveToBack(node *Node) {
	if node == nil {
		panic("The value of 'node' should not be nil")
	}
	list.erase(node)
	if list.tail != nil {
		list.tail.next = node
	} else {
		list.head = node
	}
	node.prev = list.tail
	node.next = nil
	list.tail = node
}

// Moves node to the beginning of the list, thus making 
// it list's head. Panics when node is nil
func (list *List) MoveToFront(node *Node) {
	if node == nil {
		panic("The value of 'node' should not be nil")
	}
	list.erase(node)
	if list.head != nil {
		list.head.prev = node
	} else {
		list.tail = node
	}
	node.next = list.head
	node.prev = nil
	list.head = node
}

// removes node from the list
func (list *List) erase(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}
}

// Removes node from the list
func (list *List) Remove(node *Node) {
	list.size--
	list.erase(node)
}

// Removes all elements from the list
// that have objects equal to val
func (list *List) RemoveAll(val any) {
	for it := list.head; it != nil; it = it.next {
		if it.Val == val {
			list.erase(it)
		}
	}
}

func (list *List) Empty() bool {
	return list.size == 0
}