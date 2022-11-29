package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const ApiBaseUri = "http://localhost:8000"

func main() {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := httpClient.Get(ApiBaseUri + "/api")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var jsonData struct {
		Memory struct {
			Total string `json:"total"`
			Used  string `json:"used"`
			Free  string `json:"free"`
		} `json:"memory"`
		Ip string `json:"ip"`
	}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		log.Fatalln(err)
	}

	var properties = map[string]string{
		"IP":           jsonData.Ip,
		"Total memory": jsonData.Memory.Total,
		"Free memory":  jsonData.Memory.Free,
		"Used memory":  jsonData.Memory.Used,
	}
	for name, value := range properties {
		fmt.Println(name+":", value)
	}
}
