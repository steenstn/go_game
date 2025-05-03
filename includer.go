package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
Parses a file, searches for // #include "filename" and replaces the import
statement with the content of filename
*/
func javascriptParseInclude(filename string) {

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

/*
Scans the given file for an enum and exports it to the defined filename

Example:
// #export "enum.js"
const (

	Level        SetupMessageSubType = 0
	PlayerSprite SetupMessageSubType = 1

)

Will generate enum.js:
const SetupMessageSubType_Level = 0;
const SetupMessageSubType_PlayerSprite = 1;
*/
func javascriptExport(filename string) {

	fmt.Printf("Parsing %s\n", filename)

	inputFileForRemoval, err2 := os.Open(filename)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer inputFileForRemoval.Close()

	regex := regexp.MustCompile("#export \"(.*)\"")

	scannerForRemoval := bufio.NewScanner(inputFileForRemoval)

	// Remove exported files so they can be recreated
	for scannerForRemoval.Scan() {
		if strings.Contains(scannerForRemoval.Text(), "#export") {
			filePath := regex.FindStringSubmatch(scannerForRemoval.Text())[1] // enums.js
			println("Removing " + "client/" + filePath)
			err := os.Remove("client/" + filePath)
			if err != nil {
				println(err.Error())
			}
		}
	}

	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	constRegex := regexp.MustCompile("const \\(")
	enumRegex := regexp.MustCompile("(\\w+)\\s+(\\w+)\\s*=\\s*(\\d+)")

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#export") {
			filePath := regex.FindStringSubmatch(scanner.Text())[1] // enums.js
			outputFile, err := os.OpenFile("client/"+filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			scanner.Scan() // Walk forward to get to the "const (" line

			if !constRegex.MatchString(scanner.Text()) {
				log.Fatalf("%s does not match regex", scanner.Text())
			}
			// Walk through all the consts in the enum
			for scanner.Scan() {
				matches := enumRegex.FindStringSubmatch(scanner.Text())
				if len(matches) == 0 {
					break
				}
				if len(matches) != 4 {
					log.Fatalf("Expected 3 matches, found %d", len(matches))
				}
				fmt.Printf("%s_%s = %s;\n", matches[2], matches[1], matches[3])
				outputFile.WriteString("const " + matches[2] + "_" + matches[1] + " = " + matches[3] + ";\n")
			}

		}
	}

}
