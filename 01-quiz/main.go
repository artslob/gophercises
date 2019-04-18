package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Quiz struct {
	input  string
	answer int
}

func parseQuizzesFromFile(filename string) []Quiz {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(file))
	var quizzes []Quiz
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if len(line) != 2 {
			log.Fatal("expected line with 2 fields")
		}
		expected, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("error parsing 'expected' field: %s", err)
		}
		quizzes = append(quizzes, Quiz{line[0], expected})
	}
	return quizzes
}

func shuffleQuizzes(slice []Quiz) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	shuffle := flag.Bool("shuffle", true, `set to false if dont want shuffled quizzes: '-shuffle=false'`)
	flag.Parse()

	fmt.Println("Checking for file:", *filename)
	quizzes := parseQuizzesFromFile(*filename)
	if *shuffle {
		shuffleQuizzes(quizzes)
	}

	scanner := bufio.NewScanner(os.Stdin)
	solved := 0
	for i, quiz := range quizzes {
		fmt.Printf("Problem #%d: %s = ", i+1, quiz.input)
		if !scanner.Scan() && scanner.Err() != nil {
			log.Fatal("error while reading input")
		}
		got := strings.TrimSpace(scanner.Text())
		if got == "" {
			continue
		}
		parsedGot, err := strconv.Atoi(got)
		if err != nil {
			log.Fatalf("error parsing your input: %s", err)
		}
		if parsedGot == quiz.answer {
			solved++
		}
	}
	fmt.Printf("you solved %d of %d tasks!\n", solved, len(quizzes))
}
