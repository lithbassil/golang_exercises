package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// csv file source
const defaultProblemFilename = "problems.csv"

var (
	correctAnswers int
	totalRecords   int
)

func main() {
	var (
		flagProblemFileName = flag.String("f", defaultProblemFilename, "CSV file path")
		flagTimer           = flag.Duration("t", 30*time.Second, "Timer Value")
		flagShuffle         = flag.Bool("s", false, "Shuffle the questions")
	)
	// parse the value from the command line
	flag.Parse()

	// enuser the flags are not nil
	if flagProblemFileName == nil || flagTimer == nil || flagShuffle == nil {
		fmt.Println("Missing filename or timer value")
		return
	}

	// accept user input before start, scanln to get the user input
	fmt.Printf("Press enter to start the quiz. the quiz time is %d", *flagTimer)
	fmt.Scanln()

	// open the file
	f, err := os.Open(*flagProblemFileName)
	if err != nil {
		fmt.Printf("failed err is: %v\n", err)
		return
	}

	// let go close f func
	defer f.Close()

	// read csv file
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if *flagShuffle {
		fmt.Println("Shuffling....")
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}
	totalRecords = len(records)
	if err != nil {
		fmt.Printf("failed to read the csv file: %v\n", err)
		return
	}

	quizDone := startQuiz(records)
	quizTimer := time.NewTimer(*flagTimer).C

	select {
	case <-quizDone:
	case <-quizTimer:
	}

	// loop in the csv file to get the 2 index
	// print the file result
	fmt.Printf("Result: %d/%d\n", correctAnswers, totalRecords)

}



func startQuiz(records [][]string) chan bool {
	done := make(chan bool)
	go func() {
		for i, record := range records {
			question, correctAnswer := record[0], record[1]
			fmt.Printf("%d. %s\n", i+1, question)
			var answer string
			if _, err := fmt.Scan(&answer); err != nil {
				fmt.Printf("failed to scan: %v", err)
				return
			}
			answer = strings.TrimSpace(answer)
			answer = strings.ToLower(answer)
			if answer == correctAnswer {
				correctAnswers++
			}
		}
	}
}