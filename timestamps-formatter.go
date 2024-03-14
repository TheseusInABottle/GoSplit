package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToSeconds(timestamp string) int {
	parts := strings.Split(timestamp, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	seconds, _ := strconv.Atoi(parts[2])

	totalSeconds := hours*3600 + minutes*60 + seconds
	return totalSeconds
}

func main() {
	// Prompt the user for the filename
	fmt.Println("Enter the filename of the timestamps.csv file:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create the output file
	outputFile, err := os.Create("newtimestamps.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Read the CSV file line by line and convert timestamps
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for _, line := range lines {
		if len(line) > 0 {
			// Assuming the start time is in the first column and end time in the second column
			startTime := line[0]
			endTime := line[1]

			// Convert start time to seconds
			startSeconds := convertToSeconds(startTime)

			// Convert end time to seconds
			endSeconds := convertToSeconds(endTime)

			// Create a new line with the converted timestamps
			newLine := make([]string, len(line))
			copy(newLine, line)
			newLine[0] = strconv.Itoa(startSeconds)
			newLine[1] = strconv.Itoa(endSeconds)

			// Write the new line to the output file
			err := writer.Write(newLine)
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}

	fmt.Println("Conversion completed. Output file: newtimestamps.csv")
}
