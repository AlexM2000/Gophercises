package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

func main() {
	StringPointer := flag.String("word", "AlExMilTo",
		"Word which you want to proceed")
	flag.Parse()
	fmt.Println(CamelCaseProblem(*StringPointer))
}

func CamelCaseProblem(str string) int {
	reg, err := regexp.Compile("[a-z]")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(str, "")
	return len(processedString)
}

//Caesar Cipher will be added later
