package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	fmt.Println(normalizePhoneNum("(123)456- 7892"))
}

func normalizePhoneNum(phonenum string) string {
	reg, err := regexp.Compile("[^0-9]")
	if err != nil {
		log.Fatal(err)
	}
	normalNum := reg.ReplaceAllString(phonenum, "")
	return normalNum
}
