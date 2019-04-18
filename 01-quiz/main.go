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
	"time"
)

type Quiz struct {
	input  string
	answer int
}

func (q Quiz) isRightAnswer(input int) bool {
	return q.answer == input
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

func checkQuiz(input string, quiz Quiz) bool {
	got := strings.TrimSpace(input)
	if got == "" {
		return false
	}
	parsedGot, err := strconv.Atoi(got)
	if err != nil {
		fmt.Println("Expected integer. Going to the next question!")
		return false
	}
	return quiz.isRightAnswer(parsedGot)
}

func readUserInput(ch chan<- string) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		if !sc.Scan() && sc.Err() != nil {
			log.Fatal("error while reading input")
		}
		ch <- sc.Text()
	}
}

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	shuffle := flag.Bool("shuffle", true, `set to false if dont want shuffled quizzes: '-shuffle=false'`)
	timeForAnswer := flag.Int("time", 3, "time which is given for the answer")
	flag.Parse()

	fmt.Println("Checking for file:", *filename)
	quizzes := parseQuizzesFromFile(*filename)
	if *shuffle {
		shuffleQuizzes(quizzes)
	}

	inputChannel := make(chan string)
	go readUserInput(inputChannel)

	solved := 0
	for i, quiz := range quizzes {
		fmt.Printf("Problem #%d: %s = ", i+1, quiz.input)
		select {
		case input := <-inputChannel:
			if checkQuiz(input, quiz) {
				solved++
			}
		case <-time.After(time.Duration(*timeForAnswer) * time.Second):
			fmt.Println("Times up! Going to next question.")
		}
	}
	fmt.Printf("you solved %d of %d tasks!\n", solved, len(quizzes))
}
