package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func averageAge(people map[string]int) float64 {
	var sum int
	var count int

	for _, age := range people {
		sum += age
		count++
	}

	return float64(sum) / float64(count)
}

func deletePerson(people map[string]int, name string) {

	if _, exists := people[name]; exists {
		fmt.Printf("You sure? It's a person, you know. Y/n ")
		reader := bufio.NewReader(os.Stdin)
		confirmation, _ := reader.ReadString('\n')
		confirmation = strings.TrimSpace(strings.ToLower(confirmation))

		if confirmation == "y" || confirmation == "yes" {
			delete(people, name)
			fmt.Printf("I'll have you know, you are a menace to society. Total disgrace. I hope you die alone.\n")
			fmt.Printf("A person named %s was succesfully eliminated.\n", name)
		} else {
			fmt.Println("Ight.")
		}
	} else {
		fmt.Printf("A person named %s was sadly not found.\n", name)
	}
}

func main() {

	people := map[string]int{
		"Alice":        23,
		"Leland Coyle": 40,
		"Rick Trager":  43,
	}

	people["Jonny Crane"] = 31

	fmt.Println("List:")
	for name, age := range people {
		fmt.Printf("%s: %d y.o.\n", name, age)
	}

	avg := averageAge(people)
	fmt.Printf("\nAverage age: %.2f y.o.\n", avg)

	fmt.Print("\nA person to eliminate: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	deletePerson(people, name)

	fmt.Println("\nRenewed list:")
	for name, age := range people {
		fmt.Printf("%s: %d y.o.\n", name, age)
	}
	avg2 := averageAge(people)
	fmt.Printf("\nAverage age: %.2f y.o.\n", avg2)
}
