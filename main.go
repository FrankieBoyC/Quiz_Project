package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play the game!")
	} else {
		fmt.Println("You cannot play!")
	}
	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores Uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

	// You are at 55:40 on the video
	// https://www.youtube.com/watch?v=LHhsNa_Kgns&t=3529s
}
