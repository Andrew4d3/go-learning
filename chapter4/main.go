package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Exercise 1")
	randomNumbers := exercise1()
	exercise2(randomNumbers)
	exercise3()
}

// 1. Write a for loop that puts 100 random numbers between 0 and 100 into an int slice
func exercise1() []int {
	randomNumbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		randomNumbers[i] = rand.Intn(100)
	}

	return randomNumbers
}

// 2. Loop over the slice you created in exercise 1. For each value in the slice, apply the
// following rules:
// a. If the value is divisible by 2, print “Two!”
// b. If the value is divisible by 3, print “Three!”
// c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// d. Otherwise, print “Never mind”
func exercise2(randomNumbers []int) {
	for _, number := range randomNumbers {
		resultMod2 := number % 2
		resultMod3 := number % 3

		switch {
		case resultMod2 == 0 && resultMod3 == 0:
			fmt.Println(number, "Six!")

		case resultMod2 == 0:
			fmt.Println(number, "Two!")

		case resultMod3 == 0:
			fmt.Println(number, "Three!")
		}
	}
}

// 3. Start a new program. In main, declare an int variable called total. Write a for
// loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
// The body of the for loop should be as follows:
// total := total + i
// fmt.Println(total)
// After the for loop, print out the value of total. What is printed out? What is the
// likely bug in this code?
func exercise3() {
	var total int
	for i := range 10 {
		total := total + i
		// The reason this is always printing the value and no the sum is because we're declaring new
		// total at each interation that always uses the total value from the outter scope
		// To fix it we should assign total like this total = total + i
		fmt.Println(total)
	}

	// this will not print the total sum in the loop since, that total values is only available within that scope
	fmt.Println(total)
}
