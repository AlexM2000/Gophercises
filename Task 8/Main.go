package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"./db/table"
)

func main() {
	tableRows := getData()
	fmt.Println(tableRows)
	for _, t := range tableRows {
		normalNumber := normalizePhoneNum(t.PhoneNum)
		if normalNumber != t.PhoneNum {
			tempRow := findPhoneNum(normalNumber)
			if tempRow.PhoneNum != "" {
				deletePhoneNum(t.ID)
			} else {
				t.PhoneNum = normalNumber
				updatePhoneNum(t)
			}
		}
	}
	fmt.Println(getData())
}

func normalizePhoneNum(phonenum string) string {
	reg, err := regexp.Compile("[^0-9]")
	if err != nil {
		log.Fatal(err)
	}
	normalNum := reg.ReplaceAllString(phonenum, "")
	return normalNum
}

func deletePhoneNum(id int) {
	client := &http.Client{}
	t := table.Table{}
	t.ID = id
	JSONbytes, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(
		"DELETE", "http://localhost:8000/deletePhone/"+strconv.Itoa(t.ID), bytes.NewReader(JSONbytes),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var deleteMsg string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &deleteMsg)
		if err != nil {
			panic(err)
		}
	}
}

func findPhoneNum(number string) table.Table {
	client := &http.Client{}
	t := table.Table{}
	t.PhoneNum = number
	JSONbytes, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(
		"GET", "http://localhost:8000/getPhone/"+number, bytes.NewReader(JSONbytes),
	)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &t)
		if err != nil {
			panic(err)
		}
	}
	return t
}

func updatePhoneNum(t table.Table) table.Table {
	client := http.Client{}
	JSONbytes, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(
		"PUT", "http://localhost:8000/updatePhone/"+strconv.Itoa(t.ID), bytes.NewReader(JSONbytes),
	)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &t)
		if err != nil {
			panic(err)
		}
	}
	return t
}

func getData() []table.Table {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "http://localhost:8000/get", nil,
	)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var data []table.Table
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			panic(err)
		}
	}
	return data
}
