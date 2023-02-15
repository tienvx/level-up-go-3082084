package main

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	var tests = map[string]struct {
		input []rune
		want  bool
	}{
		"empty slice":  {[]rune{}, false},
		"not contains": {[]rune{'a', 'b', 'c'}, false},
		"contains":     {[]rune{'a', 'b', 'c', 'd'}, true},
	}
	char := 'd'
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := contains(test.input, char)
			if got != test.want {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
func TestIsOpen(t *testing.T) {
	var tests = map[string]struct {
		input rune
		want  bool
	}{
		"{":        {'{', true},
		"(":        {'(', true},
		"[":        {'[', true},
		"alphabet": {'a', false},
		"number":   {'1', false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := isOpen(test.input)
			if got != test.want {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
func TestIsClosed(t *testing.T) {
	var tests = map[string]struct {
		input rune
		want  bool
	}{
		"}":        {'}', true},
		")":        {')', true},
		"]":        {']', true},
		"alphabet": {'a', false},
		"number":   {'1', false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := isClosed(test.input)
			if got != test.want {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}

func TestIsMatched(t *testing.T) {
	stack := []rune{'{', '['}
	var tests = map[string]struct {
		char rune
		want bool
	}{
		"not matched 1": {'}', false},
		"not matched 2": {')', false},
		"matched":       {']', true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := isMatched(stack, test.char)
			if got != test.want {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}

func TestPush(t *testing.T) {
	stack := []rune{'{', '['}
	want := []rune{'{', '[', '('}
	push(&stack, '(')
	if !reflect.DeepEqual(stack, want) {
		t.Errorf("Expected %v, got %v", want, stack)
	}
}

func TestPop(t *testing.T) {
	stack := []rune{'{', '[', '('}
	want := []rune{'{', '['}
	pop(&stack)
	if !reflect.DeepEqual(stack, want) {
		t.Errorf("Expected %v, got %v", want, stack)
	}
}

func TestIsBalanced(t *testing.T) {
	var tests = map[string]struct {
		expr string
		want bool
	}{
		"not balanced 1": {"{test[}", false},
		"not balanced 2": {"({welcome[here])}", false},
		"balanced":       {"[hello(world)]", true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := isBalanced(test.expr)
			if got != test.want {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
