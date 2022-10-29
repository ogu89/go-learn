package main

import "fmt"

func main(){
	fmt.Println("Welcom to my quiz game!")

	fmt.Printf("Enter\n your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10{
		fmt.Println("Yay you can play!")
	}else{
		fmt.Println("You can't play!")
		return 
	}

	score := 0
	num_questions := 0


	fmt.Println("What is better, The RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)



	if answer + " " + answer2 == "RTX 3090"{
		fmt.Println("Correct!")
		score ++
	}else if answer + " " + answer2 == "rtx 3090"{
		fmt.Println("Correct!")
	}else if answer + " " + answer2 == "rtX 3090"{
		fmt.Println("Correct!")
	}else{
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have")
	var corse uint
	fmt.Scan(&corse)

	if corse == 12{
		fmt.Println("Correct!")
		score++
	}else{
		fmt.Println("Incorrect")
	}


	fmt.Printf("You scored %v out of %v. \n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("you scored: %v%%.", percent)
}