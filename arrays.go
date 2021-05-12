package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("arra a: ", a)
	fmt.Println("index 4: ", a[4])

	fmt.Println("lenght: ", len(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("array b: ", b)

	var twoD [2][3]int
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}

	fmt.Println("Two Dimensional: ", twoD)
}