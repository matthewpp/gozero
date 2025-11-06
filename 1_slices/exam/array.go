package exam

import "fmt"

func ArrayExam() {
	var a [5]int
	fmt.Println("emp:", a)
}

func AssignArray() {
	var a [5]int
	a[0] = 100
	fmt.Println("emp:", a)
}

func InitArray() {
	var a = [...]int{1, 2, 3, 4, 5}
	fmt.Println("emp:", a)
}