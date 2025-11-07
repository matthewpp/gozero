package main

import (
	"exercise/exam"
	"fmt"
)

var unitLink = "Unit Link"

func main() {

	fmt.Printf("insurance type: %+v \ntype:%T\n", exam.UL, exam.UL)
	fmt.Printf("string type: %+v\n type:%T\n", unitLink, unitLink)

	s := isInsurance()
	fmt.Printf("insurance: %s", s)

	/* --- add method to type --- */
	var is exam.InsuranceList

	is.InitializeInsurance()

	is.Display()

	/* --- implement interface error  --- */
	e := exam.RetErrState()
	fmt.Println(e.Error())

	/* --- add method to insurance type ---- */
	// UL.ShowFullDisplay()

}

func isInsurance() exam.InsuranceType {
	return exam.UL
}

// func isInsurance() string {
// 	return string(exam.UL)
// }
