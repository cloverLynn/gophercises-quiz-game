package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		fmt.Println("FILE BROKE")
		panic(e)
	}
}

func importCSV(path string) map[string]string {
	fmt.Println("chill...Opening Quiz")
	time.Sleep(time.Second)
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	quiz := make(map[string]string)

	reader := csv.NewReader(file)
	records, recordsErr := reader.ReadAll()

	// Checks for the error
	if recordsErr != nil {
		fmt.Println("Error reading records")
	}

	// Loop to iterate through
	// and print each of the string slice
	for _, line := range records {
		q := line[0]
		a := line[1]
		quiz[q] = a
	}

	return quiz
}

func ask(reader bufio.Reader, q string, a string) bool {
	fmt.Print(q + "\n")
	res, _ := reader.ReadString('\n')
	res = strings.TrimRight(res, "\n")
	res = strings.TrimSpace(res)

	res = strings.ToLower(res)
	a = strings.ToLower(a)

	if strings.Compare(a, res) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fileFlag := flag.String("f", "./quiz.txt", "a filepath")
	flag.Parse()
	fmt.Println(*fileFlag)
	right := float64(0)
	quiz := importCSV(*fileFlag)
	maxPossible := float64(len(quiz))
	reader := bufio.NewReader(os.Stdin)
	for q, a := range quiz {
		if ask(*reader, q, a) {
			right++
		} else {

		}
	}
	perc := math.Floor((right / maxPossible) * 100)
	fmt.Printf("You got a %d/%d \n", int(right), int(maxPossible))
	fmt.Printf("%d%%\n", int(perc))

}
