package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var filename string
var timer int
var shuffle bool

func init() {
	flag.StringVar(&filename, "file", "problems.csv", "Quiz file")
	flag.IntVar(&timer, "timer", 30, "Timer length in seconds")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle questions")
}

func countdown(timer int) {
	time.Sleep(time.Duration(timer) * time.Second)
	fmt.Println("Timer over")
}

func shuffle_questions(data [][]string, number_of_questions int) [][]string {
	shuffled_data := make([][]string, number_of_questions)
	randomised := rand.Perm(number_of_questions)
	for index, value := range randomised {
		shuffled_data[index] = data[value]
	}
	return shuffled_data
}

func run_quiz(data [][]string, c chan int) {
	var number_correct int
	c <- number_correct
	for _, value := range data {
		question := value[0]
		correct_answer := value[1]

		fmt.Printf("What is the answer to: %v? \n", question)
		var user_answer string
		fmt.Scanln(&user_answer)
		fmt.Println()

		if user_answer == correct_answer {
			val := <-c
			c <- val + 1
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
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
	if shuffle {
		data = shuffle_questions(data, number_of_questions)
	}

	c := make(chan int, 2)

	go run_quiz(data, c)

	countdown(timer)

	number_correct := <-c

	fmt.Printf("Correctly answered %v out of %v question(s)\n", number_correct, number_of_questions)

}
