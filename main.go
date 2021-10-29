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

func run_quiz(data [][]string, c chan int) {
	var number_correct int
	c <- number_correct
	for _, value := range(data) {
		question := value[0]
		correct_answer := value[1]

		fmt.Printf("What is the answer to: %v? \n", question)
		var user_answer string
		fmt.Scanln(&user_answer)
		fmt.Println()

		if user_answer == correct_answer {
			val := <- c
			c <- val + 1
		}
	}
}

func main() {
	flag.Parse()

	fmt.Printf("Starting a quiz from %v!\n", filename)
	fmt.Printf("Hit Enter to start your %v second(s) timer!\n", timer)
	fmt.Scanln()

	file, _ := os.Open(filename)
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	number_of_questions := len(data)

	c := make(chan int, 2)

	go run_quiz(data, c)

	countdown(timer)

	number_correct := <-c

	fmt.Printf("Correctly answered %v out of %v question(s)\n", number_correct, number_of_questions)

}
