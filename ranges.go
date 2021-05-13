package main

import (
	"fmt"
)

func main() {

	nums := []int{5,6,7,8,9}
	sum := 0

	//We can ignore i by _
	for i, num := range nums {

		fmt.Println("index:",i, "elem:",num)
		sum += num
	}

	fmt.Println("total:", sum)

	kvs := map[int]string{1:"a", 2:"b", 3:"c", 4:"d"}

	for k,v := range kvs {
		fmt.Printf("%d -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for i, c := range "ASHOK" {
		fmt.Println(i, c)
	}
}