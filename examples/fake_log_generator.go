package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/VentGrey/picolog" // Update this import path
)

func main() {
	// Create a new random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a log file
	logFile, err := os.Create("fake_logs.txt")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()

	// Redirect standard output to file
	oldStdout := os.Stdout
	os.Stdout = logFile
	defer func() { os.Stdout = oldStdout }()

	// Generate random log entries
	for i := 0; i < 10000; i++ {
		logger := picolog.NewLogger("main/fake_logs", picolog.Info, false)
		logLevel := getRandomLogLevel(r)
		message := getRandomMessage(r)

		// Using the actual picolog methods
		switch logLevel {
		case picolog.Info:
			logger.Log(picolog.Info, message)
		case picolog.Debug:
			logger.Log(picolog.Debug, message)
		case picolog.Warning:
			logger.Log(picolog.Warning, message)
		case picolog.Error:
			logger.Log(picolog.Error, message, errors.New("Generic Error"))
		case picolog.Ok:
			logger.Log(picolog.Ok, message)
		}

		// Simulate some delay
		time.Sleep(5 * time.Millisecond)
	}
}

func getRandomLogLevel(r *rand.Rand) picolog.LogLevel {
	levels := []picolog.LogLevel{picolog.Info, picolog.Debug, picolog.Warning, picolog.Error, picolog.Ok}
	return levels[r.Intn(len(levels))]
}

func getRandomMessage(r *rand.Rand) string {
	messages := []string{
		"User logged in",
		"File not found",
		"Network error",
		"Operation successful",
		"Data saved",
	}
	return messages[r.Intn(len(messages))]
}
