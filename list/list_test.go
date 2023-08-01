package list

import "testing"

/*
	// func TestT_M(t *testing.T) {
	func TestF(t *testing.T) {
		t.Log("Similar to fmt.Println() and concurrently safe")
		t.Fail() // Will show a test case that has failed in the results
		t.FailNow() // t.Fail() + safely exits without continuing
		t.Error("t.Log() + t.Fail()")
		t.Fatal("t.Log() + t.FailNow()")
	}
*/

func checkList(list *List, arr []any, t *testing.T) {
	if list.size != len(arr) {
		t.Errorf("Incorrect size: expected %v, but found %v\n", len(arr), list.size)
	}
	for e, i := list.head, 0; e != nil; e, i = e.next, i+1 {
		if e.Val != arr[i] {
			t.Errorf("On position #%v expected %v, but found %v\n", i+1, arr[i], e.Val)
		}
	}
}

func checkTail(list *List, node *Node, t *testing.T) {
	if list.tail == nil && node != nil {
		t.Errorf("Nil tail reference while it should not be")
	}
	if list.tail != nil && node == nil {
		t.Errorf("No nil tail reference while it should be")
	}
	if list.tail.Val != node.Val {
		t.Errorf("Wrong tail address: Expected %v, but found %v", node.Val, list.tail.Val)
	}
}
func checkHead(list *List, node *Node, t *testing.T) {
	if list.head == nil && node != nil {
		t.Errorf("Nil head reference while it should not be")
	}
	if list.head != nil && node == nil {
		t.Errorf("No nil head reference while it should be")
	}
	if list.head.Val != node.Val {
		t.Errorf("Wrong head address: Expected %v, but found %v", node.Val, list.head.Val)
	}
}

func TestList_PushBack_PushFront(t *testing.T) {
	list1 := ListInit()
	list2 := ListInit()
	list3 := ListInit()

	list1.PushFront(1)
	list1.PushFront(2)
	list1.PushFront(3)

	checkList(list1, []any{3, 2, 1}, t)
	checkHead(list1, &Node{Val: 3}, t)
	checkTail(list1, &Node{Val: 1}, t)

	list2.PushBack(4)
	list2.PushBack(5)
	list2.PushBack(6)

	checkList(list2, []any{4, 5, 6}, t)
	checkHead(list2, &Node{Val: 4}, t)
	checkTail(list2, &Node{Val: 6}, t)

	list3.PushFront(0)
	list3.PushBack(1)
	list3.PushFront(2)
	list3.PushBack(3)

	checkList(list3, []any{2, 0, 1, 3}, t)
}

func TestList_InsertBefore(t *testing.T) {
	list := ListInit()
	i1 := list.PushBack(1)
	i2 := list.PushBack(2)
	list.InsertBefore(0, i1)
	list.InsertBefore(3, i2)
	checkHead(list, &Node{Val: 0}, t)
	checkTail(list, &Node{Val: 2}, t)
	checkList(list, []any{0, 1, 3, 2}, t)
}

func TestList_InsertAfter(t *testing.T) {
	list := ListInit()
	i1 := list.PushBack(2)
	i2 := list.PushFront(1)
	list.InsertAfter(0, i2)
	i3 := list.InsertAfter(3, i1)
	checkHead(list, i2, t)
	checkTail(list, i3, t)
	checkList(list, []any{1, 0, 2, 3}, t)
}

func TestList_MoveAfter(t *testing.T) {
	list := ListInit()
	i1 := list.PushBack(1)
	i2 := list.InsertAfter(2, i1)
	list.MoveAfter(i2, i1)
	checkHead(list, i1, t)
	checkTail(list, i2, t)
	
	list.MoveAfter(i1, i1)
	list.MoveAfter(i2, i2)
	i3 := list.PushBack(3)
	checkHead(list, i1, t)
	checkTail(list, i3, t)
	checkList(list, []any{1, 2, 3}, t)
}

func TestList_MoveBefore(t *testing.T) {
	list := ListInit()
	i1 := list.PushBack(1)
	i2 := list.InsertAfter(2, i1)
	list.MoveBefore(i2, i1)
	checkHead(list, i2, t)
	checkTail(list, i1, t)
	checkList(list, []any{2, 1}, t)
	
	list.MoveBefore(i1, i1)
	list.MoveBefore(i2, i2)
	i3 := list.PushFront(3)
	checkHead(list, i3, t)
	checkTail(list, i1, t)
	checkList(list, []any{3, 2, 1}, t)
}
