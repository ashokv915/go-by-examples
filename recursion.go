package main

import "fmt"

func fact(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	}
	return fact(n-1) * n;
}

func main() {

	fmt.Println("factorial of 5", fact(5))
}