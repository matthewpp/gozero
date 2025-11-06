package main

import (
	"exam/exam"
)

func main() {

	/*
		in Go, maps are not safe for concurrent use by default.
		If multiple goroutines attempt to read from and write to a map simultaneously,
		this can lead to race conditions or runtime panics. Specifically,
		the Go runtime will panic with the error
		fatal error: concurrent map read and map write when it detects concurrent access to a map.
	*/

	/* ---------- initlize map with make ------- */
	exam.InitMap()
	exam.CheckKeyInMap()
	exam.MutateMap()
	exam.Clear()
	exam.CompareMap()

}

