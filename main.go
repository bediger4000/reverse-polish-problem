package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// ItemType denotes number or operation
type ItemType int

// Possible values for ItemType
const (
	Number    ItemType = 0
	Operation ItemType = iota
)

// Item represents a number or an operation, an if the latter,
// what operation
type Item struct {
	Typ       ItemType
	Value     int
	Operation string
	Next      *Item
}

func main() {

	list := prepareList(os.Args[1:])
	result := eval(list)
	fmt.Printf("%d\n", result)
}

func eval(list []*Item) int {

	var stack *Item

	for _, node := range list {

		if node.Typ == Number {
			stack = stack.Push(node)
			continue
		}

		var left, right *Item

		stack, right = stack.Pop()
		stack, left = stack.Pop()

		var val int
		switch node.Operation {
		case "+":
			val = left.Value + right.Value
		case "-":
			val = left.Value - right.Value
		case "/":
			// Watch for div-by-zero
			val = left.Value / right.Value
		case "*":
			val = left.Value * right.Value
		}
		stack = stack.Push(&Item{Typ: Number, Value: val})
	}

	return stack.Value
}

func (stack *Item) Print(w io.Writer) {
	fmt.Fprintf(w, "stack:\n")
	for top := stack; top != nil; top = top.Next {
		if top.Typ == Number {
			fmt.Fprintf(w, "%d\n", top.Value)
			continue
		}
		fmt.Fprintf(w, "%s\n", top.Operation)
	}
}

func (stack *Item) Empty() bool {
	if stack == nil {
		return true
	}
	return false
}

func (stack *Item) Pop() (newstack *Item, top *Item) {
	top = stack
	newstack = stack.Next
	return
}

func (stack *Item) Push(item *Item) *Item {
	item.Next = stack
	return item
}

func (stack *Item) Peek() *Item {
	return stack
}

func prepareList(stringreps []string) []*Item {
	var list []*Item

	for _, str := range stringreps {
		n, err := strconv.Atoi(str)
		var item *Item
		if err == nil {
			item = &Item{Typ: Number, Value: n}
		} else {
			item = &Item{Typ: Operation, Operation: str}
		}
		list = append(list, item)
	}

	return list
}
