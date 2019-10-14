package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"unicode"
)

func main() {
	StringPointer := flag.String("word", "AlExMilTo",
		"Word which you want to proceed")
	flag.Parse()
	fmt.Println(CamelCaseProblem(*StringPointer))
	fmt.Println(CaesarCipherProblem(*StringPointer, 10))
}

func CamelCaseProblem(str string) int {
	reg, err := regexp.Compile("[a-z]")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(str, "")
	return len(processedString)
}

func CaesarCipherProblem(str string, shift int32) string {
	runed := []rune(str)
	for i := 0; i < len(runed); i++ {
		switch {
		case unicode.IsLower(runed[i]) && runed[i]+shift >= 122:
			runed[i] = (runed[i]+shift-(int32)('a'))%26 + (int32)('a')
		case unicode.IsUpper(runed[i]) && runed[i]+shift >= 90:
			runed[i] = (runed[i]+shift-(int32)('A'))%26 + (int32)('A')
		default:
			runed[i] += shift
		}
	}
	return string(runed)
}
