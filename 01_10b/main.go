package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) (*float64, error){
	"+": func(x, y float64) (*float64, error) { val := x + y; return &val, nil },
	"-": func(x, y float64) (*float64, error) { val := x - y; return &val, nil },
	"*": func(x, y float64) (*float64, error) { val := x * y; return &val, nil },
	"/": func(x, y float64) (*float64, error) {
		if y == 0 {
			return nil, fmt.Errorf("Can not divide by zero")
		} else {
			val := x / y
			return &val, nil
		}
	},
}

// parseOperand parses a string to a float64
func parseOperand(op string) (*float64, error) {
	parsedOp, error := strconv.ParseFloat(op, 64)
	if error != nil {
		return nil, fmt.Errorf("Can not parse number: %s", error)
	}
	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (*float64, error) {
	ops := strings.Fields(expr)
	if len(ops) != 3 {
		return nil, fmt.Errorf("Missing operand or operator")
	}
	left, error := parseOperand(ops[0])
	if error != nil {
		return nil, fmt.Errorf("can not parse left operand: %s", error)
	}
	right, error := parseOperand(ops[2])
	if error != nil {
		return nil, fmt.Errorf("can not parse right operand: %s", error)
	}
	f, ok := operators[ops[1]]
	if !ok {
		return nil, fmt.Errorf("Operator %s is not supported", ops[1])
	}
	result, error := f(*left, *right)
	if error != nil {
		return nil, fmt.Errorf("Can not evaluate expression %s: %s", expr, error)
	}
	return result, nil
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result, error := calculate(*expr)
	if error != nil {
		log.Fatalf("Something wrong: %s\n", error)
	}
	log.Printf("%s = %.2f\n", *expr, *result)
}
