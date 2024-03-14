package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Get the path to the directory containing the executable
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error getting executable directory:", err)
		return
	}

	// Extract ffmpeg executable
	ffmpegPath := filepath.Join(exeDir, "ffmpeg.exe")
	err = extractFFmpeg(exeDir, ffmpegPath)
	if err != nil {
		fmt.Println("Error extracting ffmpeg executable:", err)
		return
	}

	// Prompt the user for the MP3 filename
	fmt.Println("Enter the filename of the MP3 file:")
	mp3Filename := getUserInput()

	// Prompt the user for the CSV timestamp filename
	fmt.Println("Enter the filename of the CSV timestamp file:")
	csvFilename := getUserInput()

	// Load MP3 file into memory
	mp3Data, err := os.ReadFile(mp3Filename)
	if err != nil {
		fmt.Println("Error reading MP3 file:", err)
		return
	}

	// Load CSV timestamp file into memory
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		fmt.Println("Error opening CSV timestamp file:", err)
		return
	}
	defer csvFile.Close()

	// Prepare ffmpeg command
	cmd := exec.Command(ffmpegPath, "-i", "pipe:0", "-f", "segment", "-segment_times", getSegmentTimes(csvFile), "output_%03d.mp3")
	cmd.Dir = exeDir

	// Connect pipes for input and output
	cmd.Stdin = strings.NewReader(string(mp3Data))

	// Run ffmpeg command
	fmt.Println("Splitting MP3 file based on timestamps...")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running ffmpeg:", err)
		return
	}

	fmt.Println("MP3 file successfully split!")
}

func extractFFmpeg(destinationDir, ffmpegPath string) error {
	// Check if ffmpeg executable already exists in the destination directory
	_, err := os.Stat(ffmpegPath)
	if err == nil {
		return nil // ffmpeg executable already exists, no need to extract again
	}

	// Read ffmpeg executable from resources
	ffmpegData, err := ioutil.ReadFile("ffmpeg.exe")
	if err != nil {
		return err
	}

	// Write ffmpeg executable to the destination directory
	err = ioutil.WriteFile(ffmpegPath, ffmpegData, 0755)
	if err != nil {
		return err
	}

	return nil
}

func getSegmentTimes(csvFile *os.File) string {
	var segmentTimes []string
	scanner := bufio.NewScanner(csvFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			startTime, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				fmt.Println("Error parsing start time:", err)
				return ""
			}
			segmentTimes = append(segmentTimes, fmt.Sprintf("%.2f", startTime))
		}
	}
	return strings.Join(segmentTimes, ",")
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
