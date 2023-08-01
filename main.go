package main

import (
	"example/main/list"
	"fmt"
)

func main() {
	list := list.ListInit()
	list.PushBack(5)
	node := list.PushFront(7)
	list.PushBack(6)
	mark := list.InsertAfter(-1, node)
	list.InsertBefore(1, node)
	list.MoveToFront(mark)
	for it := list.Front(); it != nil; it = it.Next() {
		fmt.Printf("%v ", it.Val)
	}
}
