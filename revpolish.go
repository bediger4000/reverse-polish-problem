package main

import (
	"fmt"
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

// Item represents a number or an operation, and if the latter,
// what arithmetic operation.
// Has a Next pointer element so it can function as a stack, too,
// it's an old school intrusive data structure.
type Item struct {
	Typ       ItemType
	Value     int
	Operation string
	Next      *Item
}

func main() {

	list := prepareList(os.Args[1:])
	result := eval(list) // explicit stack version
	fmt.Printf("%d\n", result)
	result2 := eval2(list) // implicit stack version
	fmt.Printf("%d\n", result2)
}

// eval runs the input slice doing RPN arithmetic,
// and returns the result so calculated.
// The problem statement says you can assume a correct
// input expression, so this lacks any input syntax
// error handling. Uses *Item instances as both input
// commands, and elements of a FIFO stack.
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

// Pop returns the new stack, and whatever used to be at
// the top of the old stack. Use like: stack, top := stack.Pop()
func (stack *Item) Pop() (newstack *Item, top *Item) {
	top = stack
	newstack = stack.Next
	return
}

// Push puts a *Item on a stack. Returns the new stack.
// Use it like: stack = stack.Push(&Item{})
func (stack *Item) Push(item *Item) *Item {
	item.Next = stack
	return item
}

// prepareList turns a command line (slice of string) into
// a slice of *Item. The output of this function is what the
// problem statement says is the input to the desired function.
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

// eval2 destructively evaluates the RPN expression in the list slice argument.
// The slice gets used as a stack implictly.
func eval2(list []*Item) int {
	for len(list) > 1 {

		for i := 0; i < len(list); i++ {
			if list[i].Typ == Operation {
				left := list[i-2]
				right := list[i-1]
				var val int
				switch list[i].Operation {
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
				list[i] = &Item{Typ: Number, Value: val}
				// The only tricky part: excising the two Number-type
				// elements of the slice list
				list = append(list[:i-2], list[i:]...)
				break
			}
		}
	}
	return list[0].Value
}
