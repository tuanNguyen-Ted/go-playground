package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

/**Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to
a user keeping track of how many questions they get right and how many they get incorrect.
Regardless of whether the answer is correct or wrong the next question should be asked immediately afterward.
The CSV file should default to problems.csv, but the user should be able to customize the filename
via a flag.*/

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		_, _ = fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(lines))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string //question
	a string //answer
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
