package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {

	person1 := Person{Name: "Jonny", Age: 30}
	person2 := Person{Name: "Alice", Age: 23}

	fmt.Println(person1.Info())
	fmt.Println(person2.Info())

	person1.Birthday()
	person2.Birthday()

	fmt.Println("\nAfter birthday:")
	fmt.Println(person1.Info())
	fmt.Println(person2.Info())
}
