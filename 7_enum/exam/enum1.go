package exam

import "fmt"

type Day int

// Using iota to define enum values
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func Exam1(today Day) {
	switch today {
	case Monday:
		fmt.Println("Start of the workweek")
	case Friday:
		fmt.Println("End of the workweek")
	default:
		fmt.Println("It's a regular day")
	}
}
