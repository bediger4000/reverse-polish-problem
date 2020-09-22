package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var stack []int
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err == nil {
			fmt.Printf("Push %d on stack\n", n)
			stack = append(stack, n)
			continue
		}
		l := len(stack)
		fmt.Printf("stack %d deep, %v\n", l, stack)
		left := stack[l-2]
		right := stack[l-1]
		stack = stack[:l-2]
		var result int
		switch str {
		case "+":
		case "*":
			result = left * right
		case "-":
			result = left - right
		case "/":
			result = left / right
		case "%":
			result = left % right
		}
		fmt.Printf("Push result %d on stack\n", result)
		stack = append(stack, result)
	}
	fmt.Printf("Answer: %d\n", stack[len(stack)-1])
	if len(stack) > 0 {
		fmt.Printf("Stack remaining: %v\n", stack)
	}
}
