package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const problemFilenae = "problems.csv"

func main() {
	f, err := os.Open(problemFilenae)
	if err != nil {
		fmt.Printf("failed err is: %v\n", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("failed err: %v\n", err)
		return
	}
	var correctAnswers int
	for i, record := range records {
		question, correctAnswer := record[0], record[1]
		fmt.Printf("%d. %s\n", i+1, question)
		var answer string
		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Printf("failed to scan: %v", err)
			return
		}
		if answer == correctAnswer {
			correctAnswers++
		}
	}
	fmt.Printf("Result: %d/%d\n", correctAnswers, len(records))

}
