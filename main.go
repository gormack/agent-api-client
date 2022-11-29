package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiBaseUri = "http://localhost:8000"

type T struct {
	Memory struct {
		Total string `json:"total"`
		Used  string `json:"used"`
		Free  string `json:"free"`
	} `json:"memory"`
	Ip string `json:"ip"`
}

func main() {
	resp, err := http.Get(ApiBaseUri + "/api")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var jsonData T
	if err := json.Unmarshal(body, &jsonData); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("IP: ", jsonData.Ip)
	fmt.Println("Total memory: ", jsonData.Memory.Total)
	fmt.Println("Free memory: ", jsonData.Memory.Free)
	fmt.Println("Used memory: ", jsonData.Memory.Used)
}
