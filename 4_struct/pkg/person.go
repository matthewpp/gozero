package pkg

import "fmt"

type Personal struct {
	Name    string // First capital letter mean export can access in public
	Age     int
	Address string // cannot access via public
}

// func (p Personal) address() {
// 	fmt.Println("Bangkok")
// }

func (p Personal) Myaddress() {
	fmt.Println("Bangkok")
}

type PersonalP struct {
	Name    string // First capital letter mean export can access in public
	Age     int
	address string // cannot access via public
}

func (p *PersonalP) SetAddress(addr string) {
	p.address = addr
}

func (p *PersonalP) MyAddress() {
	fmt.Println("Address:", p.address)
}

type Employee struct {
	Personal PersonalP
	Salary   int
	Position string
}
