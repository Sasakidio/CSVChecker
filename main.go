package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func compareCSVs() []string{

	csvOne, err := os.Open("data/csv1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvOne.Close()

	csvTwo, err := os.Open("data/csv1.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvTwo.Close()

	readerOne := csv.NewReader(csvOne)
	readerTwo := csv.NewReader(csvTwo)

	var csvOneContents []string

	var csvTwoContents []string

	matchedContents := make([]string, 0)

	for {

		valueOne, err := readerOne.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(valueOne) > 0 {
			csvOneContents = append(csvOneContents, valueOne[0])
		}

		valueTwo, err := readerTwo.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(valueTwo) > 0 {
			csvTwoContents = append(csvTwoContents, valueTwo[0])
		}
	}

	csvTwoMap := make(map[string]bool)
	for _, item := range csvTwoContents {
		csvTwoMap[item] = true
	}

	addedMap := make(map[string]bool)


	for _, element := range csvOneContents {
		for _, elementTwo := range csvTwoContents {
			if csvTwoMap[elementTwo] && !addedMap[element] {
			matchedContents = append(matchedContents, element)
			addedMap[element] = true
			}
		}
	}

	return matchedContents
}


func exportToCSV(list []string) error {
	exportFile, err := os.Create("results/exportData.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer exportFile.Close()


	writer := csv.NewWriter(exportFile)
	defer writer.Flush()


	for _, item := range list {
		err := writer.Write([]string{item})
		if err != nil {
			log.Fatal(err)
		}
	}


	return nil
}

func main() {
	fmt.Println("Starting comparison tool...")

	results := compareCSVs()
	resultsLength := len(results)

	fmt.Printf("Total Matches: %d \n", resultsLength)

	exportToCSV(results)

	fmt.Println("File Exported. Thank you!")
}

