package main

import "fmt"

var counter int // изначально 0

type stack struct {
	elements []int
}

func (s *stack) add(v int) {
	s.elements = append(s.elements, v)
	// fmt.Println(s.elements)
}

func (s *stack) pop() int {
	if len(s.elements) == 0 {
		return 0
	}
	v := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return v
}

func (s *stack) isexist(v int) bool {
	for counter := 0; counter < len(s.elements); counter++ {
		if s.elements[counter] == v {
			return true
		}
	}
	return false
}

func main() {
	stack := new(stack)
	stack.add(20)
	stack.add(20)
	fmt.Println(stack.pop())
	fmt.Println(stack.isexist(20))
}
