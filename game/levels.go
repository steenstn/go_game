package game

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadLevel(filename string) []int {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		panic("Failed to load level")
	}
	fileAsString := string(fileContent)

	result := make([]int, len(fileAsString))

	for i, rune := range strings.ReplaceAll(fileAsString, "\n", "") {
		result[i], _ = strconv.Atoi(string(rune))
	}

	return result
}
