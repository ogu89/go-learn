package main

import "fmt"

func main(){
	fmt.Println("Welcom to my quiz game!")

	fmt.Printf("Enter\nTom your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	// fmt.Println(age >= )

}