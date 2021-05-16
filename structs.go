package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 11})

	fmt.Println(person{name: "Ali"})

	fmt.Println(&person{name: "alice", age: 90})

	fmt.Println(newPerson("Jon"))


	s := person{name: "PA", age: 33}
	fmt.Println("name",s.name, "age", s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age= 12
	fmt.Println(sp.age)

}