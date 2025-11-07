package main

import (
	"exercise/exam"
	"fmt"
)

/* create custom type function */

/*optional pattern */

func seq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func add(v int) func(int) int {
	return func(i int) int {
		return v + i
	}
}

func main() {

	s := exam.ToString(123)
	fmt.Printf("s %s\n", s)

	l := exam.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, exam.IsEven)
	fmt.Println("---- Filter 1 ------")
	for _, v := range l {
		fmt.Printf("v %v", v)

	}
	fmt.Println("")

	l2 := exam.FilterV2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, exam.IsEven)
	fmt.Println("---- Filter 2 ------")
	for _, v := range l2 {
		fmt.Printf("v %v", v)
	}
	fmt.Println("")

	l3 := exam.FilterV2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println("---- Filter 3 ------")
	for _, v := range l3 {
		fmt.Printf("v %v", v)
	}
	fmt.Println("")

	server := exam.NewServer(changeReadTimeout(), changeWriteTimeout(20))
	fmt.Printf("server %+v\n", server)

}

func changeReadTimeout() exam.Option {
	return func(s *exam.Server) {
		s.ReadTimeout = 10
	}
}

func changeWriteTimeout(v int) exam.Option {
	return func(s *exam.Server) {
		s.WriteTimeout = v
	}
}
