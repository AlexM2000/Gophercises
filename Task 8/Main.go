package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func getPhoneNums() []string {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "http://localhost:8000/getPhoneNums", nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var phoneNums []string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &phoneNums)
		if err != nil {
			panic(err)
		}
	}
	return phoneNums
}
