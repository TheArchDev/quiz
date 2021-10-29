package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var filename string
var timer int

func init() {
	flag.StringVar(&filename, "file", "problems.csv", "Quiz file")
	flag.IntVar(&timer, "timer", 30, "Timer length in seconds")
}

func countdown(timer int) {
	time.Sleep(time.Duration(timer) * time.Second)
	fmt.Println("Timer over")
}

func run_quiz(data [][]string) int {
	correct_answers := 0
	for _, value := range(data) {
		question := value[0]
		correct_answer := value[1]

		fmt.Printf("What is the answer to: %v? \n", question)
		var user_answer string
		fmt.Scanln(&user_answer)
		fmt.Println()

		if user_answer == correct_answer {
			correct_answers += 1
		}
	}
	return correct_answers
}

func main() {
	flag.Parse()

	fmt.Printf("Starting a quiz from %v!\n", filename)
	fmt.Printf("Hit Enter to start your %v second(s) timer!\n", timer)
	fmt.Scanln()

	go countdown(timer)

	file, _ := os.Open(filename)
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	number_of_questions := len(data)

	correct_answers := run_quiz(data)

	fmt.Printf("Correctly answered %v out of %v question(s)\n", correct_answers, number_of_questions)

}
