package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	start := time.Now()
	defer timeTrack(start)
	data, err1 := readCSVFile("problems.csv")
	if err1 != nil {
		fmt.Println("Error reading file: ", err1)
		os.Exit(1)
	}
	reader, err2 := parseCSVData(data)
	if err2 != nil {
		fmt.Println("Error creating CSV Reader: ", err2)
	}
	processCSV(reader)
}

func timeTrack(start time.Time) {
	fmt.Printf("This program took %s", time.Since(start))
}

func readCSVFile(fileName string) ([]byte, error) {
	f, err1 := os.Open(fileName)

	if err1 != nil {
		return nil, err1
	}

	defer f.Close()
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

func processCSV(reader *csv.Reader) {
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error Reading CSV File: ", err)
			break
		}
		fmt.Println(record)
	}
}
