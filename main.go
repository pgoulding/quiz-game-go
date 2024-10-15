package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	start := time.Now()
	defer timeTrack(start)

	csvFileName := flag.String("csv", "problems.csv", " a csv file in the format of 'question,answer'")
	flag.Parse()

	data, err1 := readCSVFile(*csvFileName)
	if err1 != nil {
		fmt.Println("Error reading file: ", err1)
		os.Exit(1)
	}
	reader, err2 := parseCSVData(data)
	if err2 != nil {
		fmt.Println("Error creating CSV Reader: ", err2)
	}
	problems := processCSV(reader)

	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect :( ")
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func timeTrack(start time.Time) {
	fmt.Printf("This program took %s \n", time.Since(start))
}

func readCSVFile(fileName string) ([]byte, error) {
	f, err1 := os.Open(fileName)

	if err1 != nil {
		return nil, err1
	}

	defer f.Close() // close the stream at the end no matter what
	data, err2 := io.ReadAll(f)

	if err2 != nil {
		return nil, err2
	}
	return data, nil
}

func parseCSVData(d []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(d))
	return reader, nil
}

func processCSV(reader *csv.Reader) []problem {
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}
