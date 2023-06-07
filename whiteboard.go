package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const defaultProblemFileName = "01-Quiz/problems.csv" // Const is a way to define variable as global and defualt.

func main() {

	f, err := os.Open(defaultProblemFileName)
	if err != nil {
		fmt.Printf("Faild to open the file: %v\n", err)
		return
	}
	defer f.Close()

	// read CSV file by use csv package

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Faild to Read the file: %v\n", err)
		return
	}

	var correctAnswers int
	// display one Q at the time
	for i, record := range records {
		question, correctAnswer := record[0], record[1]
		fmt.Printf("%d. %s?\n", i+1, question)
		// get answer, then proceed next one immediatley
		var answer string
		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Printf("Faild to scan: %v\n", err)
			return
		}
		if answer == correctAnswer {
			correctAnswers++
		}

	}
	// output number of Quses total + correct
	fmt.Printf("Result: %d/%d\n", correctAnswers, len(records))

}
