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
	
}
