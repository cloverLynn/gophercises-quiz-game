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

func timer(timer *int, right *float64, total *float64) {
	t1 := time.NewTimer(timer * time.Second)
	<-t1.C
	fmt.Println("Timer expired")
	endGame(right, total)
}

func endGame(right *float64, total *float64) {
	perc := math.Floor((*right / *total) * 100)
	fmt.Printf("You got a %d/%d \n", int(*right), int(*total))
	fmt.Printf("%d%%\n", int(perc))
	os.Exit(0)
}

func startGame() bool {
	reader := bufio.NewReader(os.Stdin)
	if ask(*reader, "Ready? y/n", "y") {
		return true
	} else {
		os.Exit(0)
		return false
	}
}

func game(fileFlag *string, timerFlag *int) {
	right := float64(0)
	quiz := importCSV(*fileFlag)
	maxPossible := float64(len(quiz))
	reader := bufio.NewReader(os.Stdin)
	startGame()
	go timer(*timerFlag, &right, &maxPossible)
	for q, a := range quiz {
		if ask(*reader, q, a) {
			right++
		} else {
		}
	}
	endGame(&right, &maxPossible)
}

func main() {
	fileFlag := flag.String("f", "./quiz.txt", "a filepath")
	timerFlag := flag.Int("t", 30, "timer seconds")
	flag.Parse()
	game(fileFlag, timerFlag)
}
