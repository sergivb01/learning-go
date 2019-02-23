package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	// Define the CSV flag and set the default value
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the limit for the quiz in seconds")

	flag.Parse()

	// Get file contents
	lines := readLines(csvFilename)

	// Parse all the problems
	problems := parseCSV(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	//problemloop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
			//break problemloop
		case answer := <-answerCh:
			if answer == problem.answer {
				correct++
			}
		}

	}

	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func readLines(csvFilename *string) [][]string {
	// Open CSV file with a pointer to the filename
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Create a CSV reader from the file
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provied CSV file.")
	}

	return lines
}

func parseCSV(lines [][]string) []problem {
	//Create an array of problem with a defined length
	res := make([]problem, len(lines))

	for i, line := range lines {
		res[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return res
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
