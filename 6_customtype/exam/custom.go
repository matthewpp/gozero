package exam

import "fmt"

type InsuranceType string

var UL InsuranceType = "Unit Link"
var UN InsuranceType = "Universal Life"
var WL InsuranceType = "Whole Life"

func (i InsuranceType) ShowFullDisplay() {
	fmt.Printf("This is full display of insurance: %s\n", i)
}
