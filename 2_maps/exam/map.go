package exam

import (
	"fmt"
	"maps"
)

func InitMap() {
	// make map
	fmt.Println("------InitMap ----")

	var m map[string]int
	fmt.Println("init map m:", m)

	m = make(map[string]int)
	// fmt.Println("init map m:", m)

	// assign key-value
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	// print map
	fmt.Printf("maps m: %+v\n", m)

	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	var nameScore = map[string]string{
		"Tom":  "100",
		"Matt": "90",
	}
	for k, v := range nameScore {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func CheckKeyInMap() {
	fmt.Println("------CheckKeyInMap ----")

	m := make(map[string]string)
	m["a"] = "first"
	m["b"] = "second"

	if v, ok := m["a"]; ok {
		fmt.Printf("value is: %s\n", v)
	} else {

		fmt.Printf("key not found\n")
	}

	if _, ok := m["c"]; ok {
		fmt.Printf("key c exist\n")
	} else {
		fmt.Printf("key c not found\n")
	}
}

func MutateMap() {
	fmt.Println("------MutateMap ----")

	m := make(map[string]string)
	m["a"] = "first"
	m["b"] = "second"
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
	delete(m, "a")
	fmt.Println("after delete key a")

	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func Clear() {
	fmt.Println("------Clear ----")

	m := make(map[string]string)
	m["a"] = "first"
	m["b"] = "second"
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
	clear(m)

	fmt.Println("after clear maps m:")
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func CompareMap() {
	fmt.Println("------CompareMap ----")

	m1 := map[string]string{
		"a": "first",
		"b": "second",
		"c": "third",
	}

	m2 := map[string]string{
		"a": "first",
		"b": "second",
		"c": "cccc",
	}

	fmt.Printf("m1 == m2: %t\n", maps.Equal(m1, m2))

	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"

	fmt.Printf("maps before edit %+v\n", m)

	editMap(m)

	fmt.Printf("maps after edit %+v\n", m)
}

func editMap(m map[string]string) {
	m["a"] = "this new a"
}
