package main

import (
	"flag"
	"log"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	var stack []rune
	for _, char := range expr {
		if isOpen(char) {
			push(&stack, char)
		} else if isClosed(char) {
			if isMatched(stack, char) {
				pop(&stack)
			} else {
				push(&stack, char)
			}
		}
	}
	return len(stack) == 0
}

func contains(list []rune, char rune) bool {
	for _, item := range list {
		if item == char {
			return true
		}
	}

	return false
}

func isOpen(char rune) bool {
	open := []rune{'{', '[', '('}
	return contains(open, char)
}

func isClosed(char rune) bool {
	close := []rune{'}', ']', ')'}
	return contains(close, char)
}

func isMatched(stack []rune, char rune) bool {
	bracketsMap := map[rune]rune{'{': '}', '[': ']', '(': ')'}
	top := len(stack) - 1
	return bracketsMap[stack[top]] == char
}

func push(stack *[]rune, char rune) {
	*stack = append(*stack, char)
}

func pop(stack *[]rune) {
	top := len(*stack) - 1
	(*stack) = (*stack)[:top]
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
