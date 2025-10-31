package main

import "fmt"

type person struct {
	name string
	age  int
}

// type employee1 struct {
// 	p        person
// 	salary   int
// 	position string
// }

// type employee2 struct {
// 	person
// 	salary   int
// 	position string
// }

// type employee3 struct {
// 	pkg.Personal
// 	salary   int
// 	position string
// }

// type employee4 struct {
// 	personal pkg.Personal
// 	salary   int
// 	position string
// }

func main() {
	//Pointer
	//Zero value

	p := person{
		name: "neng",
		age:  100,
	}

	fmt.Printf("person: %+v\n", p)
	fmt.Printf("person name: %s\n", p.name)

	// p := person{
	// 	name: "neng",
	// }

	// fmt.Printf("person: %+v\n", p)

	// var p person
	// p.name = "neng"
	// p.age = 100

	// fmt.Printf("person: %+v\n", p)

	/* --- struct without create type ---- */
	// s := struct {
	// 	position string
	// 	salary   int
	// }{
	// 	position: "Developer",
	// 	salary:   1000,
	// }

	// fmt.Printf("employee: %+v\n", s)

	/* ---- method ---- */

	// p := person{
	// 	name: "neng",
	// 	age:  100,
	// }

	// p.description()

	// p.changeName()

	// p.description()

	/* ------ modify struct ------ */
	// p := person{
	// 	name: "neng",
	// 	age:  100,
	// }

	// p.description()
	// changeName(p)
	// p.description()

	// change name with pointer
	// p.description()
	// changeNamePtr(&p)
	// p.description()

	/* ------- embed struct -------- */
	// s := employee{
	// 	p: person{
	// 		name: "neng",
	// 		age:  100,
	// 	},
	// 	salary:   1000,
	// 	position: "developer",
	// }

	// fmt.Printf("employee embed person %+v\n", s)

	// s.p.description()
	// s.description()

	/* ------- embed composition struct -------*/
	// s := employee2{
	// 	person: person{
	// 		name: "neng",
	// 		age:  100,
	// 	},
	// 	salary:   1000,
	// 	position: "developer",
	// }

	// fmt.Printf("employee embed composition person %+v\n", s)

	// s.description()
	// s.person.description()

	/*---- embed composition struct from other pakcage ------ */
	// s := employee3{
	// 	Personal: pkg.Personal{
	// 		Name: "Neng",
	// 		Age:  100,
	// 	},
	// 	salary:   1000,
	// 	position: "developer",
	// }

	// fmt.Printf("employee embed composition person %+v\n", s)

	// s.Myaddress()

	/*---- embed struct from other pakcage ------ */
	// s := employee4{
	// 	personal: pkg.Personal{
	// 		Name: "Chonlatee",
	// 		Age:  100,
	// 	},
	// 	salary:   1000,
	// 	position: "developer",
	// }

	// fmt.Printf("employee embed person %+v\n", s)
	// s.personal.Myaddress()

	// pass by value example
	// a := 5
	// someFunc(&a)
	// fmt.Printf("%d", a)
}

func (p person) description() {
	fmt.Printf("This is me, name: %s, age: %d\n", p.name, p.age)
}

func (p *person) changeName() {
	p.name = "Chonlatee"
}

// func (e employee1) description() {

// 	fmt.Printf("This is me, name: %s, age: %d, position: %s, salary: %d\n", e.p.name, e.p.age, e.position, e.salary)
// }

// func (e employee2) description() {
// 	fmt.Printf("This is me, name: %s, age: %d, position: %s, salary: %d\n", e.name, e.age, e.position, e.salary)
// }

// func changeName(p person) {
// 	p.name = "chonlatee"
// }

// func changeNamePtr(p *person) {
// 	p.name = "chonlatee"
// }

/*
	go alwasy pass by value

	- In Go, everything is passed by value: when you pass an argument to a function, Go copies the value and sends the copy to the function.
	- For primitive types (int, float, bool), this means the function works on a copy, and changes inside the function don’t affect the original value.
	- For structs, the entire struct is copied.
	- For slices, maps, and channels, the internal reference to the underlying data is copied, which allows the function to modify the data but not the container itself.
	- To modify the original value or struct, you need to pass a pointer to the function.
*/

// func someFunc(x *int) {
// 	*x = 2 // This changes the value that x is pointing to, so the caller's variable is now 2.

// 	y := 7 // A new local variable y is created with the value 7.

// 	x = &y // x is now pointing to y (which is a new local variable). However, this change to x does NOT affect the original pointer that was passed to the function!
// }

// go does not support the concept of pointers to pointers

// #include <stdio.h>

// void someFunc(int **x) {
//     **x = 2;  // Change the value of what the original pointer points to (affects the caller)

//     int y = 7;
//     *x = &y;  // Change the caller's pointer to point to y (now this change affects the caller)
// }

// int main() {
//     int a = 10;
//     int *ptr = &a;  // Create a pointer to a

//     printf("Before someFunc: %d\n", *ptr);  // Prints 10

//     someFunc(&ptr);  // Pass pointer to the pointer (int **)

//     printf("After someFunc: %d\n", *ptr);  // Prints 7, because the pointer now points to y
//     return 0;
// }
