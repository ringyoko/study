package main

import (
	"fmt"
)

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	slice := arr[1:4]
	fmt.Println("Initial slice:", slice)

	slice = append(slice, 6)
	fmt.Println("After appending:", slice)

	indexToRemove := 1
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("After removing:", slice)
}
