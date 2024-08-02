package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

var fileFlag = flag.String("f", "problems.csv", "Path to csv file with problems")
var helpFlag = flag.Bool("h", false, "Show help message")
var timeFlag = flag.Int("t", 30, "Set timeout for answering all questions in seconds")
var shuffleFlag = flag.Bool("s", false, "Shuffle question list")

func prompt(ac chan<- string) error {
	input := new(string)
	fmt.Print("Answer: ")
	_, err := fmt.Scanln(input)
	if err != nil {
		return err
	}

	*input = strings.TrimSpace(*input)

	ac <- *input
	return nil
}

func loop(records [][]string, tChan <-chan time.Time) (total, correct int) {
	if *shuffleFlag {
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}

	total = len(records)
	ansChan := make(chan string)
	defer close(ansChan)

	for _, record := range records {
		question, correctAns := record[0], record[1]
		fmt.Println(question)
		go func() {
			err := prompt(ansChan)
			if err != nil {
				log.Fatal(err)
			}
		}()

		// Handling answer and timeout
		select {
		case ans := <-ansChan:
			if ans == correctAns {
				correct++
			}
		case <-tChan:
			fmt.Println()
			return
		}
	}

	return
}

func main() {
	// Parsing flags
	flag.Parse()
	if *helpFlag {
		flag.Usage()
		return
	}

	// Parsing content
	file, err := os.Open(*fileFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Preparing
	fmt.Printf("You'll have %d seconds\n", *timeFlag)
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()

	// Game loop
	timerChan := time.After(time.Duration(*timeFlag) * time.Second)

	t, c := loop(records, timerChan)
	fmt.Printf("Answered %d of %d questions\n", c, t)
}
