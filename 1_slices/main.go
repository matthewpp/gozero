package main

import "fmt"

/* a slice is a dynamic, flexible, and more powerful version of an array. */

func main() {

	fmt.Println("matthew")
	/* initialize array */
	// var a [5]int
	// fmt.Println("emp:", a)

	/* assign value to array */
	// a[0] = 100
	// fmt.Println("emp:", a)

	/* initialize array with value ( compiler wil check size for you ) */
	// var a = [...]int{1, 2, 3, 4, 5}
	// fmt.Println("emp:", a)

	/* slice topic */
	/* initialize slice */
	// var a []int
	// a = []int{1, 2, 3, 4, 5}
	// fmt.Println("emp:", a)

	// var a = []int{1, 2, 3, 4, 5}
	// fmt.Println("emp:", a)

	// a := []int{1, 2, 3, 4, 5}
	// fmt.Println("emp:", a)

	// a = append(a, 6)
	// fmt.Println("emp:", a)

	/* ---------- start len cap append topic -------------- */
	// var s []string
	// fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = make([]string, 3)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s[0] = "a"
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s := make([]string, 3)
	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s := make([]string, 0, 3)
	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	/* take carefully when append with len and cap*/
	// s := make([]string, 0, 6)
	// s = append(s, "a")
	// s = append(s, "b")

	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))

	/* take carefully when append with len and cap is the same*/
	// s := make([]string, 6)
	// s = append(s, "a")
	// s = append(s, "b")

	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))

	/* loop over slice */
	// s = make([]string, 0, 3)
	// s = append(s, "a")
	// s = append(s, "b")
	// s = append(s, "c")

	// for _, v := range s {
	// 	fmt.Printf("value: %s\n", v)
	// }

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("value: %s\n", s[i])
	// }

	/* --------------------- start copy slice -------------------*/
	// var s []string
	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("emp:", b, "len:", len(b), "cap:", cap(b))

	/* if destition slice have cap less than source slice */
	// var s []string
	// s = append(s, "a")
	// s = append(s, "b")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("emp:", b, "len:", len(b), "cap:", cap(b))

	/* if destition slice have cap less than source slice */
	// var s []string
	// s = append(s, "a")
	// s = append(s, "b")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// var b []string
	// copy(b, s)

	// fmt.Println("emp:", b, "len:", len(b), "cap:", cap(b))

	/* copy and then change destination slice */
	// var s []string
	// s = append(s, "a")
	// s = append(s, "b")
	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("slice b:", b, "len:", len(b), "cap:", cap(b))

	// b[0] = "c"

	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("slice b:", b, "len:", len(b), "cap:", cap(b))

	/* --------------------- end copy slice -------------------*/

	/* cut slice */
	// s := []string{"a", "b", "c", "d", "e", "f"}
	// s = s[1:]
	// fmt.Printf("s %+v\n", s)
	// s = s[:len(s)-1]
	// fmt.Printf("s %+v\n", s)
	// s = s[2:4]

	// fmt.Printf("s %+v\n", s)
	/* end cut slice */

	/* ----------------- start modify slice topic ----------------*/
	/* if cap less than len */

	/* modify slice with index 0*/
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call modify:", s, "len:", len(s), "cap:", cap(s))

	// modifySlice(s)

	// fmt.Println("slice s after call modify:", s, "len:", len(s), "cap:", cap(s))

	/* modify slice with append before index 0 not modified */
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call modify:", s, "len:", len(s), "cap:", cap(s))

	// modifySliceWithAppendBefore(s)

	// fmt.Println("slice s after call modify:", s, "len:", len(s), "cap:", cap(s))

	/* modify slice with append after index 0 is modified */
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call modify:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Printf("address slice s before call %p\n", s)

	// modifySliceWithAppendAfter(s)

	// fmt.Printf("address slice s after call %p\n", s)

	// fmt.Println("slice s after call modify:", s, "len:", len(s), "cap:", cap(s))

	// /*modify Slice with append  with cap more than len with edit index*/
	// s := make([]string, 6)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call modify:", s, "len:", len(s), "cap:", cap(s))

	// modifySliceWithEditIndex(s)

	// fmt.Println("slice s after call modify:", s, "len:", len(s), "cap:", cap(s))

	/* change slice by return value from func */
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call retSlice:", s, "len:", len(s), "cap:", cap(s))

	// s = retSlice(s)

	// fmt.Println("slice s after call retSlice:", s, "len:", len(s), "cap:", cap(s))

	/* change slice by poiter */

	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call ptrSlice:", s, "len:", len(s), "cap:", cap(s))

	// ptrSlice(&s) // yellow flag if slice growth

	// fmt.Println("slice s after call ptrSlice:", s, "len:", len(s), "cap:", cap(s))

	/* re-slice by pointer */
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// b := s

	// fmt.Println("slice s before call reSlice:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("slice b before call reSlice:", b, "len:", len(b), "cap:", cap(b))

	// reSlice(&s) // yellow flag in here

	// fmt.Println("slice s after call reSlice:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("slice b before call reSlice:", b, "len:", len(b), "cap:", cap(b))

	/* ----------------- end modify slice topic ----------------*/

	/*
		in Go, a slice itself is a reference type, meaning it contains a pointer to the underlying array.
		So when you pass a slice to a function, you're already passing a reference to the data, not a copy
		same with channel, maps

	*/

	// fmt.Printf("slice s with append %+v\n", getSliceWithAppend())

	/* if know length use this pattern are more efficient */
	// fmt.Printf("slice s with assign %+v\n", getSliceWithAssign())
}

func getSliceWithAppend() []string {
	var r []string
	for _, v := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		r = append(r, v)
	}

	return r
}

func getSliceWithAssign() []string {
	s := []string{"a", "b", "c", "d", "e", "f", "g"}
	r := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[i]
	}

	return r
}

func modifySlice(s []string) {
	s[0] = "z"
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
}

func modifySliceWithAppendBefore(s []string) {
	s = append(s, "x") // new slice
	s[0] = "z"         // change new slice not effect origin slice
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
}

func modifySliceWithAppendAfter(s []string) {
	s[0] = "z"         // edit slice effect origin slice
	s = append(s, "x") // create new slice not effect origin slice
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
}

func modifySliceWithEditIndex(s []string) {
	s[0] = "z"
	s[3] = "x"
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
}

func retSlice(s []string) []string {
	s = append(s, "new value")
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
	return s
}

func ptrSlice(s *[]string) {
	*s = append(*s, "new value")
	fmt.Println("slice s in func:", s, "len:", len(*s), "cap:", cap(*s))
}

func reSlice(s *[]string) {
	*s = []string{"a"}
}
