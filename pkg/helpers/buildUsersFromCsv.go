package helpers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadUsersFromCsv(userList []string, tempPasswordList []string) ([]string, []string) {
	
	fmt.Print("Please enter the file path to read from: ")
	// Set up input reader for user interaction
	reader := bufio.NewReader(os.Stdin)
	filePath, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}


	filePath = strings.TrimSpace(filePath)

	if filePath == "" {
		log.Fatal("File path cannot be empty.")
	}

	// handle file path with " on the start and end
	if strings.HasPrefix(filePath, "\"") && strings.HasSuffix(filePath, "\"") {
		filePath = strings.TrimPrefix(filePath, "\"")
		filePath = strings.TrimSuffix(filePath, "\"")
	}

	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	
	defer f.Close()

	// Read CSV records
	r := csv.NewReader(f)
	r.FieldsPerRecord = 2

	records, err := r.ReadAll()

	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Process records
	for _, record := range records {
		userList = append(userList, record[0])
		tempPasswordList = append(tempPasswordList, record[1])
	}

	return userList, tempPasswordList
}