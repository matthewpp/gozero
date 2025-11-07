package exam

import "strconv"

func ToString(n int) string {
	return strconv.Itoa(n)
}

type keep func(v int) bool

func Filter(l []int, fn keep) []int {
	var r []int
	for _, v := range l {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}

func IsEven(val int) bool {
	return val%2 == 0
}

func FilterV2(l []int, fn func(v int) bool) []int {
	var r []int
	for _, v := range l {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}
