package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	fmt.Println("Exercise 1: Slices")
	// Original slice
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	// Subslice 1: First two values
	sub1 := greetings[0:2]

	// Subslice 2: Second, third, and fourth values
	sub2 := greetings[1:4]

	// Subslice 3: Fourth and fifth values
	sub3 := greetings[3:5]

	// Print all slices
	fmt.Println("Original slice:", greetings)
	fmt.Println("Subslice 1 (first two):", sub1)
	fmt.Println("Subslice 2 (second to fourth):", sub2)
	fmt.Println("Subslice 3 (fourth and fifth):", sub3)
}

func exercise2() {
	fmt.Println("Exercise 2: Runes and strings")
	message := "Hi 👩 and 👨"
	// fourthChar := rune(message[3])
	fmt.Printf("Fourth char: %c\n", message[3])
}

type employee struct {
	firstName string
	lastName  string
	id        int
}

func exercise3() {
	fmt.Println("Exercise 3: Structs")

	employee1 := employee{
		"andrew",
		"davis",
		1,
	}

	employee2 := employee{
		firstName: "andrew",
		lastName:  "davis",
		id:        1,
	}

	var employee3 employee

	employee3.firstName = "andrew"
	employee3.lastName = "davis"
	employee3.id = 1

	fmt.Println(employee1, employee2, employee3)
}
