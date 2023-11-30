package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	dotenv "github.com/joho/godotenv"
)

const baseUrl = "https://adventofcode.com/2023/day/%d/input"

func getToken() string {
	err := dotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	token, tokenExists := os.LookupEnv("TOKEN")

	if !tokenExists {
		log.Fatal("Token not found")
	}

	return token
}

func getInput(day uint) string {
	token := getToken()
	url := fmt.Sprintf(baseUrl, day)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", fmt.Sprintf("session=%v", token))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error getting input file", err)
	}

	if res.StatusCode != 200 {
		log.Fatal("Get input response not 200:", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body", err)
	}
	res.Body.Close()

	return string(body)
}

func GetInputOrSample(day uint, sample string) string {
	var input string
	if len(os.Args) < 2 {
		input = getInput(day)
	} else {
		input = sample
	}

	return input
}
