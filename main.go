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

func createQuiz() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Helping you create a quiz!")
	fmt.Println("What should the quiz be called?")
	title, _ := reader.ReadString('\n')
	title = strings.TrimRight(title, "\n")
	fmt.Println("Really " + title + " ? okay....")
	fmt.Println("You will now begin entering Question and Answer pairs")
	quiz := make(map[string]string)
	do := true
	for do {
		fmt.Print("Q: ")
		q, _ := reader.ReadString('\n')
		q = strings.TrimRight(q, "\n")
		fmt.Print("A: ")
		a, _ := reader.ReadString('\n')
		a = strings.TrimRight(a, "\n")
		quiz[q] = a
		fmt.Println("Another? y/n")
		res, _ := reader.ReadString('\n')
		res = strings.TrimRight(res, "\n")
		if res == "y" {
		} else {
			do = false
		}
	}
	fmt.Println(quiz)
	filePath := "Quizes/" + title + ".csv"
	file, err := os.Create(filePath)
	check(err)
	defer file.Close()
	for key, value := range quiz {
		q := key
		a := value
		_, err := file.WriteString(q + "," + a + "\n")
		check(err)
	}
	os.Exit(0)
}

func timer(t int, right *float64, total *float64) {
	dur := time.Duration(t) * time.Second
	t1 := time.NewTimer(dur)
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
	fmt.Println("(1) Play Quiz (2) Create Quiz")
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	res = strings.TrimRight(res, "\n")
	if res == "1" {
		game(fileFlag, timerFlag)
	} else {
		createQuiz()
	}
}
