package main

import (
	"example/main/list"
	"fmt"
)

func print(list *list.List) {
	for it := list.Front(); it != nil; it = it.Next() {
		fmt.Printf("%v ", it.Val)
	}
	fmt.Println()
}

func main() {
	list := list.ListInit()
	it1 := list.PushBack(7)
	it2 := list.PushFront(4)
	it3 := list.PushBack(8)
	it4 := list.PushFront(2)
	print(list)
	i := 1
	j := -1
	for it := list.Front(); it != nil; it = it.Next() {
		if it.Val == 8 {
			list.InsertAfter(j, it)
			list.InsertBefore(i, it)
		}
	}
	print(list)
	list.MoveAfter(it1, it3)
	print(list)
	list.MoveBefore(it2, it4)
	print(list)
	list.PushFront(-10)
	list.PushBack(10)
	print(list)
}
