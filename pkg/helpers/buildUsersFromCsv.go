package helpers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadUsersFromCsv reads user data from a CSV file and returns two slices containing usernames and temporary passwords
// Parameters:
//   userList: Slice to store usernames
//   tempPasswordList: Slice to store temporary passwords
// Returns:
//   []string: Updated slice of usernames
//   []string: Updated slice of temporary passwords
func ReadUsersFromCsv(userList []string, tempPasswordList []string) ([]string, []string) {
	
	// Prompt user for input file path
	fmt.Print("Please enter the file path to read from: ")
	// Set up input reader for user interaction
	reader := bufio.NewReader(os.Stdin)
	filePath, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Remove leading/trailing whitespace from file path
	filePath = strings.TrimSpace(filePath)

	// Validate that file path is not empty
	if filePath == "" {
		log.Fatal("File path cannot be empty.")
	}

	// handle file path with " on the start and end
	// Remove enclosing quotes if present
	if strings.HasPrefix(filePath, "\"") && strings.HasSuffix(filePath, "\"") {
		filePath = strings.TrimPrefix(filePath, "\"")
		filePath = strings.TrimSuffix(filePath, "\"")
	}

	// Open the CSV file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Ensure file is closed after function completes
	defer f.Close()

	// Create CSV reader and configure it to expect 2 fields per record
	r := csv.NewReader(f)
	r.FieldsPerRecord = 2

	// Read all records from CSV file
	records, err := r.ReadAll()

	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Process records by appending username and password to respective slices
	for _, record := range records {
		userList = append(userList, record[0])        // Add username from first column
		tempPasswordList = append(tempPasswordList, record[1]) // Add password from second column
	}

	return userList, tempPasswordList
}
