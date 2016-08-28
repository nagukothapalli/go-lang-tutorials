package main

import (
	"fmt"
	"os"
)

func main() {

	// Some Randome junk to make hands to be dirty
	fmt.Println("Hello World , Welcome to GO.")
	// Declaring the variable and constants ( Enums in java size)
	name := "Hello World"
	const age uint8 = 24
	const (
		names = "Go lang"
		exp   = 10 * 20
	)
	fmt.Printf("Name: %s, with age %d , exp %d", name, age, exp)

	// Data Types
	var numberOfBytesprinted int
	var theError error
	var pi float32 = 3.14

	fmt.Printf(" Float %f", pi)
	//	var age uint8 = 24
	fmt.Printf("Age %d", age)
	// if else

	if numberOfBytesprinted, theError = fmt.Printf("Pi %.2f age %d", pi, age); theError != nil {
		os.Exit(2)
	}
	fmt.Println(numberOfBytesprinted)

	// calling functions  with error handing
	greeting, error := greet("Nagu")

	if error == nil {
		fmt.Println("Greating  call value " + greeting)
	}

	// arrays
	words := make([]string, 4)
	words[0] = "GUNA"
	fmt.Println(words)
	fmt.Println(len(words))
}

func greet(name string) (string, error) {

	return ("Hello World " + name), nil
}
