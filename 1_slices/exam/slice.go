package exam

import "fmt"

func InitSlice() {
	// a[low:high]
	var a = []int{1, 2, 3, 4, 5} // var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("emp:", a)

	newA := a[1:4]
	fmt.Println("slice a[1:4]:", newA)

	// newB := a[10:20]
	// fmt.Println("slice a[10:20]:", newB)

	// len ความยาว cap ความจุ
	s := a[:0]
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s1 := a[2:]
	fmt.Println(" s1:", s1, "len:", len(s1), "cap:", cap(s1))

	s[0] = 100
	fmt.Println("after modify s", s, "len:", len(s), "cap:", cap(s))
}

func LenCapAppendExam() {
	var s []string //Slices are like references to arrays

	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "a")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "a")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "a")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "s")
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "h"
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
}

func LenCapAppendExam2() {

	// s := make([]string, 0, 3)
	// fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	//---

	s := make([]string, 3)
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "a")
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "a")
	fmt.Println("s2:", s, "len:", len(s), "cap:", cap(s))

	// s3 = append(s3, "a")
	// s3 = append(s3, "b")
	// fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))

}

func LoopOverSlice() {
	s := make([]string, 0, 3)
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c")

	for _, v := range s {
		fmt.Printf("value: %s\n", v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("value: %s\n", s[i])
	}
}

func CopySlice() {
	// src := []string{"a", "b", "c"}
	// dst := make([]string, len(src))
	// copy(dst, src)
	// fmt.Println("src:", src, "dst:", dst)

	// s := []string{"a", "b"}
	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("slice b:", b, "len:", len(b), "cap:", cap(b))

	// b[0] = "c"

	// fmt.Println("after change slice s:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("after change slice b:", b, "len:", len(b), "cap:", cap(b))

	/* cut slice */
	s := []string{"a", "b", "c", "d", "e", "f"}
	s = s[1:]
	fmt.Printf("s %+v\n", s)
	s = s[:len(s)-1]
	fmt.Printf("s %+v\n", s)
	s = s[2:4]

	fmt.Printf("s %+v\n", s)
	/* end cut slice */
}

func cutSlice() {
	fmt.Println("---- cut slice ----")
	s := []string{"a", "b", "c", "d", "e", "f"}
	s = s[1:]
	fmt.Printf("s %+v\n", s)
	s = s[:len(s)-1]
	fmt.Printf("s %+v\n", s)
	s = s[2:4]

	fmt.Printf("s %+v\n", s)
}
