package main

import (
	"fmt"

	"github.com/Jonny-Burkholder/data-structures/linkedlist"
)

func main() {
	l := linkedlist.NewLinkedList()
	l.Push("one")
	l.Push("two")
	l.Unshift(0)

	fmt.Println(l)

	fmt.Println(l.Pop())

	data, _ := l.Shift()

	fmt.Println(data)

	fmt.Println(l)
	fmt.Println(l.Head)
	fmt.Println(l.Tail)
}
