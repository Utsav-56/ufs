package ulog

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Box drawing characters
const (
	topLeft     = "╭"
	topRight    = "╮"
	bottomLeft  = "╰"
	bottomRight = "╯"
	horizontal  = "─"
	vertical    = "│"
)

// Colors for different log levels
var (
	warningColor = color.New(color.FgYellow).SprintFunc()
	messageColor = color.New(color.FgBlue).SprintFunc()
	infoColor    = color.New(color.Reset).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
	successColor = color.New(color.FgGreen).SprintFunc()
	ongoingColor = color.New(color.FgHiYellow).SprintFunc() // Orange-like
	tagColor     = color.New(color.Bold).SprintFunc()
)

// Logger is a utility for logging with box-style outputs and colors
type Logger struct {
	showTimestamp bool
	padding       int
}

// NewLogger creates a new Logger instance
func NewLogger(showTimestamp bool, padding int) *Logger {
	if padding < 1 {
		padding = 1
	}
	return &Logger{
		showTimestamp: showTimestamp,
		padding:       padding,
	}
}

// Default logger instance with default settings
var DefaultLogger = NewLogger(true, 1)

// formatBox creates a box around the given message with the specified color function
func (l *Logger) formatBox(message string, tag string, colorFunc func(a ...interface{}) string) string {
	lines := strings.Split(message, "\n")

	// Find the longest line to determine box width
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	// Add space for tag if provided
	if tag != "" {
		if len(tag)+4 > maxLength {
			maxLength = len(tag) + 4
		}
	}

	// Add padding
	maxLength += l.padding * 2

	// Create the box
	var result strings.Builder

	// Top border with tag if provided
	topBorder := topLeft + strings.Repeat(horizontal, maxLength) + topRight
	if tag != "" {
		tagDisplay := " " + tagColor(tag) + " "
		topBorderParts := strings.SplitN(topBorder, "", 2)
		result.WriteString(colorFunc(topBorderParts[0]+tagDisplay+topBorderParts[1]) + "\n")
	} else {
		result.WriteString(colorFunc(topBorder) + "\n")
	}

	// Add timestamp if enabled
	if l.showTimestamp {
		timestamp := time.Now().Format("15:04:05")
		paddedLine := vertical + strings.Repeat(" ", l.padding) + timestamp
		paddedLine += strings.Repeat(" ", maxLength-len(timestamp)) + vertical
		result.WriteString(colorFunc(paddedLine) + "\n")
	}

	// Message lines
	for _, line := range lines {
		paddedLine := vertical + strings.Repeat(" ", l.padding) + line
		paddedLine += strings.Repeat(" ", maxLength-len(line)) + vertical
		result.WriteString(colorFunc(paddedLine) + "\n")
	}

	// Bottom border
	result.WriteString(colorFunc(bottomLeft + strings.Repeat(horizontal, maxLength) + bottomRight))

	return result.String()
}

// Warning logs a warning message in yellow
func (l *Logger) Warning(message string, tag ...string) {
	tagStr := ""
	if len(tag) > 0 {
		tagStr = tag[0]
	}
	fmt.Println(l.formatBox(message, tagStr, warningColor))
}

// Message logs a message in blue
func (l *Logger) Message(message string, tag ...string) {
	tagStr := ""
	if len(tag) > 0 {
		tagStr = tag[0]
	}
	fmt.Println(l.formatBox(message, tagStr, messageColor))
}

// Info logs an info message in default terminal color
func (l *Logger) Info(message string, tag ...string) {
	tagStr := ""
	if len(tag) > 0 {
		tagStr = tag[0]
	}
	fmt.Println(l.formatBox(message, tagStr, infoColor))
}

// Error logs an error message in red
func (l *Logger) Error(message string, tag ...string) {
	tagStr := ""
	if len(tag) > 0 {
		tagStr = tag[0]
	}
	fmt.Println(l.formatBox(message, tagStr, errorColor))
}

// Success logs a success message in green
func (l *Logger) Success(message string, tag ...string) {
	tagStr := ""
	if len(tag) > 0 {
		tagStr = tag[0]
	}
	fmt.Println(l.formatBox(message, tagStr, successColor))
}

// Ongoing logs an ongoing operation message in orange-like color
func (l *Logger) Ongoing(message string, tag ...string) {
	tagStr := ""
	if len(tag) > 0 {
		tagStr = tag[0]
	}
	fmt.Println(l.formatBox(message, tagStr, ongoingColor))
}

// Global convenience functions that use the default logger

// Warning logs a warning message in yellow using the default logger
func Warning(message string, tag ...string) {
	DefaultLogger.Warning(message, tag...)
}

// Message logs a message in blue using the default logger
func Message(message string, tag ...string) {
	DefaultLogger.Message(message, tag...)
}

// Info logs an info message in default terminal color using the default logger
func Info(message string, tag ...string) {
	DefaultLogger.Info(message, tag...)
}

// Error logs an error message in red using the default logger
func Error(message string, tag ...string) {
	DefaultLogger.Error(message, tag...)
}

// Success logs a success message in green using the default logger
func Success(message string, tag ...string) {
	DefaultLogger.Success(message, tag...)
}

// Ongoing logs an ongoing operation message in orange-like color using the default logger
func Ongoing(message string, tag ...string) {
	DefaultLogger.Ongoing(message, tag...)
}
