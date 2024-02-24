package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Scanner is a struct that holds a scanner
type Scanner struct {
	scanner *bufio.Scanner
}

// NewScanner returns a new Scanner
func NewScanner() *Scanner {
	return &Scanner{scanner: bufio.NewScanner(os.Stdin)}
}

// Scan returns the next token from the scanner
func (s *Scanner) Scan() string {
	s.scanner.Scan()
	return strings.TrimSpace(s.scanner.Text())
}

// Prompt prints a prompt to the console
func (s *Scanner) Prompt(prompt string) {
	fmt.Print(prompt)
}

// Close closes the scanner
func (s *Scanner) Close() {
	s.scanner = nil
}

// IsClosed returns true if the scanner is closed
func (s *Scanner) IsClosed() bool {
	return s.scanner == nil
}

// IsEmpty returns true if the scanner is empty
func (s *Scanner) IsEmpty() bool {
	return s.scanner.Text() == ""
}

// IsEmptyOrClosed returns true if the scanner is empty or closed
func (s *Scanner) IsEmptyOrClosed() bool {
	return s.IsEmpty() || s.IsClosed()
}

