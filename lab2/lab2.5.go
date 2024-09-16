package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {

	var myRect = Rectangle{width: 24, height: 30}
	solution := myRect.width * myRect.height

	fmt.Println("This rectangle's area equals", solution)
}
