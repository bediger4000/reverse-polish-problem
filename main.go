package main

import (
	"errors"
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

	stack := prepareStack(os.Args[1:])
	stack.Print(os.Stdout)

	value, stack, err := eval(stack)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem: %v\n", err)
		return
	}
	fmt.Printf("%d\n", value)
	stack.Print(os.Stdout)
}

func eval(stack *Item) (int, *Item, error) {
	var op, left, right *Item

	for !stack.Empty() {
		stack, op = stack.Pop()
		if op.Typ == Number {
			return op.Value, stack, nil
		}
		stack, right = stack.Pop()
		stack, left = stack.Pop()
		fmt.Printf("%d %s %d\n", left.Value, op.Operation, right.Value)

		var val int
		switch op.Operation {
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
	return -1, stack, errors.New("shouldn't get here")
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

func prepareStack(stringreps []string) *Item {
	var stack *Item

	for _, str := range stringreps {
		n, err := strconv.Atoi(str)
		if err == nil {
			stack = stack.Push(&Item{Typ: Number, Value: n})
			continue
		}
		stack = stack.Push(&Item{Typ: Operation, Operation: str})
	}

	return stack
}
