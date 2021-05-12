package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2

	fmt.Println("map:", m)

	fmt.Println("val:", m["a"])

	fmt.Println("len:", len(m))

	delete(m, "b")
	fmt.Println("del:", m)

	key, prs := m["a"]
	fmt.Println("key:", key, "prs:", prs)

	n := map[string]int{"one":1, "two":2}
	fmt.Println("dcl:", n)
}