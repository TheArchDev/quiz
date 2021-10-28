package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var filename string
var timer int

func init() {
	flag.StringVar(&filename, "file", "problems.csv", "Quiz file")
	flag.IntVar(&timer, "timer", 30, "Timer length in seconds")
}


func main() {
	flag.Parse()
	fmt.Printf("Starting a quiz from %v!\n", filename)
	fmt.Printf("Hit Enter to start the %v second(s) timer!\n", timer)
	fmt.Scanln()

	file, _ := os.Open(filename)
	reader := csv.NewReader(file)

	correct_answers := 0
	question_number := 0
	for keep_reading := true; keep_reading; question_number++ {
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
			if user_answer == correct_answer {
				correct_answers += 1
			}
		}
	}
	fmt.Printf("Correctly answered %v out of %v question(s)\n", correct_answers, question_number)
}
