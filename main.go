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

	var user_answers []bool
	for i, keep_reading := 0, true; keep_reading; i++ {
		line, err := reader.Read()
		if err != nil {
			keep_reading = false
		} else {
			question := line[0]
			correct_answer := line[1]
			fmt.Printf("What is the answer to: %v? ", question)
			var user_answer string
			fmt.Scanln(&user_answer)
			fmt.Println()
			user_answers = append(user_answers, user_answer == correct_answer)
		}
	}
	fmt.Println(user_answers)
}
