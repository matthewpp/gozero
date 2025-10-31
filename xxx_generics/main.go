package main

import (
	"fmt"
)

/*
	NOTE: generics in go work in compile time

	the Go compiler generates specific versions of the generic code based on the concrete types you use in your program
*/

type Addable interface {
	int | float32 | float64 | string
}

// Generic function with a custom constraint
func Add[T Addable](a, b T) T {
	return a + b
}

/* with out generics any */
// type stack struct {
// 	val []int
// }

// func (s *stack) Push(v int) {
// 	s.val = append(s.val, v)
// }

// func (s *stack) Pop() int {
// 	v := s.val[len(s.val)-1]
// 	s.val = s.val[:len(s.val)-1]

// 	return v
// }

// func (s *stack) Peek() int {
// 	return s.val[len(s.val)-1]
// }

// func newStack() *stack {
// 	return &stack{
// 		val: make([]int, 0),
// 	}
// }

/* stack with generics */
type stack[T any] struct {
	val []T
}

func (s *stack[T]) Push(v T) {
	s.val = append(s.val, v)
}

func (s *stack[T]) Pop() T {
	v := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]

	return v
}

func (s *stack[T]) Peek() T {
	return s.val[len(s.val)-1]
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		val: make([]T, 0),
	}
}

// func Max[T comparable](a, b T) T {
//     if a > b {
//         return a
//     }
//     return b
// }

func main() {
	// fmt.Println(Add(10, 5))
	// fmt.Println(Add(10.0, 5.0))
	// fmt.Println(Add("a", "b"))

	// s := newStack[int]()

	// s.Push(10)
	// s.Push(50)
	// s.Push(60)
	// s.Push(100)
	// fmt.Println(s.Pop())
	// fmt.Println(s.Peek())

	s := newStack[string]()

	s.Push("G")
	s.Push("O")
	s.Push("L")
	s.Push("A")
	s.Push("N")
	s.Push("G")
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())

}
