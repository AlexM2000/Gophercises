package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv",
		"a csv file in the format of 'question,anwser' (default problems.csv")
	limit := flag.Int("limit", 30, "the time for the quiz in seconds (default 30)")
	flag.Parse()
	quizQuestions := make(map[string]string)
	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	csvfile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal("Couldn't open csv file: ", err)
	}
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quizQuestions[record[0]] = record[1]
	}
	var problemNum, userAnwser = 1, ""
	for question, correctAnswer := range quizQuestions {
		anwserCh := make(chan string)
		go func() {
			fmt.Print("Problem №", problemNum, ": ", question, " = ")
			_, err := fmt.Scanln(&userAnwser)
			if err != nil {
				log.Fatalln(err)
			}
			anwserCh <- userAnwser
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d of %d \n", problemNum-1, len(quizQuestions)-1)
			return
		case userAnwser := <-anwserCh:
			if userAnwser != correctAnswer {
				fmt.Println(userAnwser)
				fmt.Println("You scored ", problemNum-1, " of ", len(quizQuestions)-1)
				return
			}
			problemNum++
			timer.Reset(time.Duration(*limit) * time.Second)
		}

	}
	fmt.Printf("Congratulations! You scored %d of %d \n", len(quizQuestions)-1, len(quizQuestions)-1)
}
