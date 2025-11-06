package exam

import (
	"fmt"
)

func BasicExam() {
	var a int = 10
	fmt.Println("Value of a:", a)

	var c, name, isActive = 20, "John", true
	fmt.Println("Value of c:", c)
	fmt.Println("Value of name:", name)
	fmt.Println("Value of isActive:", isActive)

	var (
		x int     = 5
		y float64 = 10.5
		z string  = "Golang"
	)
	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)
	fmt.Println("Value of z:", z)

	// short from
	var i, j int = 1, 2
	k := 3
	cc, python, java := true, false, "no!"

	fmt.Println(i, j, k, cc, python, java)
}

func ExampleVariable() {

	// example var
	// bool
	var isActive bool = true

	// string
	var name string = "John"

	// int  int8  int16  int32  int64
	var age int = 30

	// uint uint8 uint16 uint32 uint64 uintptr
	var score uint = 100

	// byte // alias for uint8
	var grade byte = 'A'

	// rune // alias for int32
	var unicode rune = 'A'
	//      // represents a Unicode code point

	// float32 float64
	var pi float64 = 3.14159

	// complex64 complex128
	var complexNum complex128 = complex(5, 7)

	fmt.Println(isActive, name, age, score, grade, unicode, pi, complexNum)

}

func ZeroValueExam() {
	var a int
	var b float64
	var c string
	var d bool

	fmt.Printf("Default value of int: %d\n", a)
	fmt.Printf("Default value of float64: %f\n", b)
	fmt.Printf("Default value of string: '%s'\n", c)
	fmt.Printf("Default value of bool: %t\n", d)
}

func ConvertExam() {
	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(, y, z)
}

func TypeInferenceExam() {
	a := 10
	b := 20.5
	c := "Golang"
	d := true

	fmt.Printf("Type of a: %T, Value: %v\n", a, a)
	fmt.Printf("Type of b: %T, Value: %v\n", b, b)
	fmt.Printf("Type of c: %T, Value: %v\n", c, c)
	fmt.Printf("Type of d: %T, Value: %v\n", d, d)
}

func ConstantExame() {
	const Pi = 3.14
	const Greeting = "Hello, World!"

	fmt.Println("Value of Pi:", Pi)
	fmt.Println("Value of Greeting:", Greeting)
}

func PointerExam() {
	var a int = 42
	var p *int = &a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", p)
	fmt.Println("Value at address p:", *p)

	*p = 100
	fmt.Println("New value of a:", a)	
}