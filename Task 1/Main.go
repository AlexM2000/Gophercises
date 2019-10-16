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
		"a csv file in the format of 'question,answer' (default problems.csv")
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
	var problemNum, userAnswer = 1, ""
	for question, correctAnswer := range quizQuestions {
		answerCh := make(chan string)
		go func() {
			fmt.Print("Problem №", problemNum, ": ", question, " = ")
			_, err := fmt.Scanln(&userAnswer)
			if err != nil {
				log.Fatalln(err)
			}
			answerCh <- userAnswer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d of %d \n", problemNum-1, len(quizQuestions)-1)
			return
		case userAnswer := <-answerCh:
			if userAnswer != correctAnswer {
				fmt.Println(userAnswer)
				fmt.Println("You scored ", problemNum-1, " of ", len(quizQuestions)-1)
				return
			}
			problemNum++
			timer.Reset(time.Duration(*limit) * time.Second)
		}

	}
	fmt.Printf("Congratulations! You scored %d of %d \n", len(quizQuestions)-1, len(quizQuestions)-1)
}
