package exam

import (
	"fmt"
	"gozero/chonlatee/pkg"
)

func BasicExam() {
	fmt.Println("")
	fmt.Println("--- BasicExam ---")
	var a int = 10
	fmt.Println("Value of a:", a)

	p := pkg.Personal{
		Name:    "Ohm",
		Address: "Bangkok",
		Age:     35,
	}

	fmt.Printf("person: %+v\n", p)
}

func ZeroStructExam() {
	fmt.Println("")
	fmt.Println("--- ZeroStructExam ---")
	var p pkg.Personal
	fmt.Printf("person: %+v\n", p)
}

func PointerStructExam() {
	fmt.Println("")
	fmt.Println("--- PointerStructExam ---")
	var p *pkg.Personal = &pkg.Personal{
		Name:    "Ohm",
		Address: "Bangkok",
		Age:     35,
	}

	fmt.Printf("person: %+v\n", p)
}

func MethodStructExam() {
	fmt.Println("")
	fmt.Println("--- MethodStructExam ---")
	p := pkg.PersonalP{
		Name: "Ohm",
		Age:  35,
	}
	p.SetAddress("Bangkok")
	p.MyAddress()
}

func EmbedComExam() {
	fmt.Println("")
	fmt.Println("--- EmbedComExam ---")
	e := pkg.Employee{
		Personal: pkg.PersonalP{
			Name: "Ohm",
			Age:  35,
		},
		Salary:   50000,
		Position: "Developer",
	}

	fmt.Printf("employee: %+v\n", e)
}
