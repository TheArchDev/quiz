package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "problems.csv", "Specify quiz file")
}

func main() {
	flag.Parse()
	fmt.Printf("Starting a quiz from %v!\n", filename)
	file, _ := os.Open(filename)
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
