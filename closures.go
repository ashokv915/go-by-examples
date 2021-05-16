package main

import "fmt"

func intSeq() func() int {
	i:=0

	return func() int {
		i++
		return i
	}
}

func main() {

	value := intSeq()

	fmt.Println(value())
	fmt.Println(value())
	fmt.Println(value())

	newvalue := intSeq()
	fmt.Println(newvalue())
}