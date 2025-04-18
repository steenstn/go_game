package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func includeStuff(filename string) {

	fmt.Printf("Parsing %s\n", filename)

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("out/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	regex := regexp.MustCompile("#include \"(.*)\"")

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#include") {
			filePath := regex.FindStringSubmatch(scanner.Text())[1]
			fmt.Printf("Including %s\n", filePath)
			includeFile, includeErr := os.Open(filePath)
			if includeErr != nil {
				log.Fatal(includeErr)
				panic("Could not open " + filePath)
			}
			includeScanner := bufio.NewScanner(includeFile)
			for includeScanner.Scan() {
				outputFile.WriteString(includeScanner.Text() + "\n")
			}
			includeFile.Close()
		} else {
			outputFile.WriteString(scanner.Text() + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
