package main

import "strconv"

type Stack struct {
	elems []int
}

func NewStack(max int) Stack {
	return Stack{
		elems: make([]int, 0, max),
	}
}

func (s *Stack) Pop() int {
	last := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return last
}

func (s *Stack) Push(val int) {
	s.elems = append(s.elems, val)
}

func evalRPN(tokens []string) int {
	stack := NewStack(len(tokens))

	res := 0
	for _, t := range tokens {
		switch t {
		case "-":
			lhs := stack.Pop()
			rhs := stack.Pop()
			res = rhs - lhs
			stack.Push(res)
		case "+":
			lhs := stack.Pop()
			rhs := stack.Pop()
			res = rhs + lhs
			stack.Push(res)
		case "*":
			lhs := stack.Pop()
			rhs := stack.Pop()
			res = rhs * lhs
			stack.Push(res)
		case "/":
			lhs := stack.Pop()
			rhs := stack.Pop()
			res = rhs / lhs
			stack.Push(res)
		default:
			num, _ := strconv.Atoi(t)
			stack.Push(num)
		}

	}

	return stack.Pop()
}
