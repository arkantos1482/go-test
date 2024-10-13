package main

import (
	"log"
	"runtime"
)

func main() {
	// Simulate a function call chain
	firstFunction()
}

func firstFunction() {
	secondFunction()
}

func secondFunction() {
	thirdFunction()
}

func thirdFunction() {
	logStackTrace()
}

func logStackTrace() {
	// Create a buffer to hold the stack trace
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false) // Get the current stack trace
	stackTrace := string(buf[:n])

	// Log the stack trace
	log.Printf("Call Stack:\n%s", stackTrace)
}
