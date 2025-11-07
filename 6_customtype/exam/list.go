package exam

import "fmt"

type InsuranceList []InsuranceType

func (i *InsuranceList) InitializeInsurance() {
	*i = append(*i, UL)
	*i = append(*i, UN)
	*i = append(*i, WL)
}

func (i InsuranceList) Display() {
	for _, v := range i {
		fmt.Printf("insurance %s\n", v)
	}
}
