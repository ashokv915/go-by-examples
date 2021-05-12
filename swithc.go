package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon")
	default:
		fmt.Println("It's after Noon")
	}

	whatAmI := func(i interface{})  {
		switch t:= i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Int")
		default:
			fmt.Println("Other", t)
		}
	}

	whatAmI(true)
	whatAmI(12)
	whatAmI("qwe")


}