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
    csvFilenamePointer := flag.String("csv", "problems.csv",
    	"a csv file in the format of 'question,anwser' (default problems.csv")
    limitPointer := flag.Int("limit", 30,"the time for the quiz in seconds (default 30)")
    flag.Parse()
    quizQuestions := make(map[string]string)
    timer := time.NewTimer(time.Duration(*limitPointer) * time.Second)

    csvfile, err := os.Open(*csvFilenamePointer)
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
    for question, realAnswer := range quizQuestions {
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
			fmt.Println("You scored ", problemNum-1, " of ", len(quizQuestions)-1)
			return
		case userAnwser := <-anwserCh:
			if userAnwser != realAnswer {
				fmt.Println(userAnwser)
				fmt.Println("You scored ", problemNum-1, " of ", len(quizQuestions)-1)
				return
			}
			problemNum++
			timer.Reset(time.Duration(*limitPointer) * time.Second)
		}

	}
	fmt.Println("Congratulations! You scored ", len(quizQuestions)-1, " of ", len(quizQuestions)-1)
	//Comment for pull request 
}