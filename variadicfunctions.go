//Variadic functions can be called with any number of trailing arguments.
//For example, fmt.Println is a common variadic function
package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("nums:", nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println("total:", total)
}

func main() {

	sum(1,2)
	sum(1,2,3)

	a := []int{1,2,3,4,5}
	sum(a...)
}