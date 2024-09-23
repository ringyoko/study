package main

import (
	"fmt"
	"study/mathutils"
)

func main() {

	var number int

	fmt.Print("Enter number: ")
	fmt.Scan(&number)

	fmt.Printf("Factorial of %d is %d\n", number, mathutils.Factorial(number))

}
