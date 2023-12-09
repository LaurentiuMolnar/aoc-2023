package utils

import (
	"cmp"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	if len(os.Args) < 2 || os.Args[1] != "--sample" {
		input = getInput(day)
	} else {
		input = sample
	}

	return strings.TrimSuffix(input, "\n")
}

func GetInputOrSampleLines(day uint, sample string) []string {
	lines := GetInputOrSample(day, sample)

	return strings.Split(lines, "\n")
}

func Sum(nums []int) int64 {
	var sum int64 = 0
	for _, n := range nums {
		sum += int64(n)
	}
	return sum
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func PadInputAndSplit(input string, block rune) []string {
	lines := strings.Split(input, "\n")
	lineLen := len(lines[0])

	var horizontalLine string = ""
	var result []string = make([]string, len(lines)+2)

	for i := 1; i <= lineLen+2; i++ {
		horizontalLine = strings.Join([]string{horizontalLine, string(block)}, "")
	}

	for i := range result {
		if i == 0 || i == len(result)-1 {
			result[i] = horizontalLine
		} else {
			result[i] = strings.Join([]string{string(block), lines[i-1], string(block)}, "")
		}
	}

	return result
}

func PadInput(input string, block rune) string {
	return strings.Join(PadInputAndSplit(input, block), "\n")
}

func IsDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func Min[T cmp.Ordered](items []T) T {
	var min = items[0]

	for i := 1; i < len(items); i++ {
		if items[i] < min {
			min = items[i]
		}
	}

	return min
}

func SliceContains[T comparable](item T, slice []T) bool {
	for i := 0; i < len(slice); i++ {
		if item == slice[i] {
			return true
		}
	}
	return false
}

func MapStringsToInts(strs []string) []int {
	var result []int = make([]int, len(strs))

	for i, s := range strs {
		num, err := strconv.Atoi(s)

		if err != nil {
			log.Fatalf("Cannot convert %s to int\n", s)
		}
		result[i] = num
	}
	return result
}
