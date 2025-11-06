package main

import (
	"exercise/exam"
	"fmt"
)

func main() {

	/*
		rune is Go's solution for handling Unicode characters effectively in a world of global text and diverse languages.
	*/

	exam.InitRune()
	fmt.Println("")

	exam.NormalString()
	fmt.Println("")

	exam.RuneEmoji()

}
