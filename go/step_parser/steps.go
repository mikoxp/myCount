package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type CsvLine struct {
	Column1 string
	Column2 string
}

func main() {

	lines, err := readCsv("Kroki.csv")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i, line := range lines {
		if i > 0 {
			data := CsvLine{
				Column1: line[0],
				Column2: line[1],
			}
			x := data.Column1 + " " + data.Column2
			fmt.Println(x)
			writeToFile(f, x)
		}
	}
}

func readCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func writeToFile(f *os.File, line CsvLine) {
	s := fmt.Sprintf("%s,%s\n", line.Column1, line.Column2)
	_, err2 := f.WriteString(s)

	if err2 != nil {
		log.Fatal(err2)
	}
}
