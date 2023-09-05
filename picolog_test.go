package picolog

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

//HACK: captureOutput captures the output of a function and returns it as a string.
// It's used to test the output of the logger.
func captureOutput(fun func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fun()

	w.Close()
	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)
	return string(out)
}

func TestLogLevels(t *testing.T) {
	logger := NewLogger("picolog/test_log_levels", Info, false)

	tests := []struct {
		level LogLevel
		message string
	}{
		{Info, "This is an info message"},
		{Debug, "This is a debug message"},
		{Warning, "This is a warning message"},
		{Error, "This is an error message"},
		{Ok, "This is an ok message"},
	}

	for _, tt := range tests {
		output := captureOutput(func () {
			logger.Log(tt.level, tt.message)
		})

		t.Logf("Testing log level %v with message: %v", tt.level, tt.message)

		if !strings.Contains(output, tt.message) {
			t.Errorf("Expected output to contain %s, got %s", tt.message, output)
		}
	}
}

func TestInvalidLogLevel(t *testing.T) {
	logger := NewLogger("picolog/test_invalid_log_level", LogLevel(-1), false)

	output := captureOutput(func() {
		logger.Log(LogLevel(-1), "This is invalid")
	})

	t.Logf("Testing an invalid log level %v with output : %v\n", logger.MinLogLevel, output)

	if !strings.Contains(output, "Picolog Error: Invalid log level provided, defaulting to INFO") {
		t.Errorf("Expected log message to mention invalid log level, got: %s", output)
	}
}

func TestLogWithOptionalError(t *testing.T) {
	logger := NewLogger("picolog/test_log_with_optional_error", Error, false)
	sampleError := "This is a sample error"

	output := captureOutput(func() {
		logger.Log(Error, "This is an error", errors.New(sampleError))
	})


	if !strings.Contains(output, sampleError) {
		t.Errorf("Expected log message to contain '%s', got: %s", sampleError, output)
	}
}

func TestLogWithoutOptionalError(t *testing.T) {
	logger := NewLogger("picolog/test_log_without_optional_error", Info, false)
	message := "This is sparta!"

	output := captureOutput(func() {
		logger.Log(Info, message)
	})

	if !strings.Contains(output, message) {
		t.Errorf("Expected log message to contain '%s', got: %s", message, output)
	}
}

func TestOptionalColors(t *testing.T) {
	logger := NewLogger("picolog/test_optional_colors", Info, true)

	output := captureOutput(func() {
		logger.Log(Info, "This should be in color")
	})

	if !strings.Contains(output, "\033[") {
		t.Errorf("Expected log message to contain ANSI escape code for color, got: %s", output)
	}

	logger.EnableColours = false

	output = captureOutput(func() {
		logger.Log(Info, "This should not be in color")
	})

	if strings.Contains(output, "\033[") {
		t.Errorf("Did not expect log message to contain ANSI escape code for color, got: %s", output)
	}
}
