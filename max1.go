package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// stack nodes: linked list type. Easy to add a head node,
// or to remove a head node. Nil value means "empty list".
type Node struct {
	value int
	next  *Node
}

// struct Stack holds pointers to heads of two stacks:
// one for data added via push() method
// 2nd for the max data value at that level of the data stack
// Both stacks use structs Node to hold items pushed on stack
type Stack struct {
	data     *Node // data values
	maxValue *Node // max data, kept in sync with data
}

var (
	NilStack = errors.New("nil Stack")
	NoData   = errors.New("no stack data")
)

// Go can do interface types to enforce (func signature)
// compliance with abstract types.
type AbstractMaxStack interface {
	max() (int, error)
	pop() (int, error)
	push(val int)
}

var _ AbstractMaxStack = (*Stack)(nil)

func main() {
	stack := &Stack{data: nil, maxValue: nil}
	for _, str := range os.Args[1:] {
		if str == "--" {
			popAll(stack)
			continue
		}
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		stack.push(n)
	}
	popAll(stack)
}

func popAll(stack *Stack) {
	var n int
	var err error
	for err == nil {
		m, err := stack.max()
		if err != nil {
			fmt.Printf("max() returns %v\n", err)
			break
		}
		n, err = stack.pop()
		if err != nil {
			fmt.Printf("pop() returns %v\n", err)
			break
		}
		fmt.Printf("max %d, %d\n", m, n)
	}
}

func (s *Stack) max() (int, error) {
	if s == nil {
		return 0, NilStack
	}

	if s.data == nil {
		return 0, NoData
	}

	return s.maxValue.value, nil
}

func (s *Stack) pop() (int, error) {
	if s == nil {
		return 0, NilStack
	}

	if s.data == nil {
		return 0, NoData
	}
	val := s.data.value
	// pop both stacks in unison
	s.data = s.data.next
	s.maxValue = s.maxValue.next
	return val, nil
}

func (s *Stack) push(val int) {
	if s.data == nil {
		// empty stacks, put first item on both
		s.data = &Node{value: val}
		s.maxValue = &Node{value: val}
		return
	}
	// careful with this part
	max := s.maxValue.value
	// push val on data stack
	s.data = &Node{value: val, next: s.data}
	if s.data.value > max {
		max = s.data.value
	}
	// push max value at this level of stack on max value stack
	s.maxValue = &Node{value: max, next: s.maxValue}
}
