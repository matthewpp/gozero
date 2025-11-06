package main

import (
	"exercise/exam"
	"fmt"
)

/* a slice is a dynamic, flexible, and more powerful version of an array. */

func main() {

	fmt.Println("---- Array ----")
	/* initialize array */
	// exam.InitArray()

	// exam.ArrayExam()

	// /* assign value to array */
	// exam.AssignArray()

	/* ---------- start len cap append topic -------------- */
	fmt.Println("---- Slice ----")
	exam.LenCapAppendExam2()

	/* loop over slice */
	exam.LoopOverSlice()
	/* --------------------- start copy slice -------------------*/
	fmt.Println("---- Copy Slice ----")
	exam.CopySlice()

}
