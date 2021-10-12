package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Feed struct {
	ContractAddress string   `json:"contractAddress"`
	ContractVersion int      `json:"contractVersion"`
	DecimalPlaces   int      `json:"decimalPlaces"`
	Heartbeat       int64    `json:"heartbeat"`
	History         bool     `json:"history"`
	Multiply        string   `json:"multiply"`
	Name            string   `json:"name"`
	Symbol          string   `json:"symbol"`
	Pair            []string `json:"pair"`
	Path            string   `json:"path"`
	NodeCount       int      `json:"nodeCount"`
	Status          string   `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var configs = "testnet_config.json"

	var config []Feed

	jsonFile, err := os.Open(configs)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleRequests() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func main() {
	handleRequests()
}
