package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	data     *Node // data values
	maxValue *Node // max data, kept in sync with data
}

var (
	NilStack = errors.New("nil Stack")
	NoData   = errors.New("no stack data")
)

func main() {
	stack := &Stack{data: nil, maxValue: nil}
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		stack.push(n)
	}

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
	s.data = s.data.next
	s.maxValue = s.maxValue.next
	return val, nil
}

func (s *Stack) push(val int) {
	if s.data == nil {
		s.data = &Node{value: val}
		s.maxValue = &Node{value: val}
		return
	}
	max := s.maxValue.value
	s.data = &Node{value: val, next: s.data}
	if s.data.value > max {
		max = s.data.value
	}
	s.maxValue = &Node{value: max, next: s.maxValue}
}
