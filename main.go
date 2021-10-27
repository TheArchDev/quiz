package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Thanks for starting a quiz!")

	file, _ := os.Open("problems.csv")
	reader := csv.NewReader(file)

	var user_answers [12]bool
	for i := 0; i < 12; i++ {
		line, _ := reader.Read()
		question := line[0]
		correct_answer := line[1]
		fmt.Printf("What is the answer to: %v? ", question)
		var user_answer string
		fmt.Scanln(&user_answer)
		fmt.Println()
		user_answers[i] = user_answer == correct_answer
	}
	fmt.Println(user_answers)
}
