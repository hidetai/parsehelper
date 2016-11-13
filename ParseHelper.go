package parsehelper

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Parser is an interface for parsers which the parse helper helps
type Parser interface {
	// Parse a line of text files
	Parse(line string)

	// Print the result of the parase
	PrintResult()
}

// Helper to parse text files

// ParseHelper is a struct which helps parse files.
type ParseHelper struct {
	textFiles []string
	parsers   []Parser

	isInitialized bool
	isParseDone   bool
}

// NewParseHelper is a func which initialize the ParseHelper struct.
func (p *ParseHelper) NewParseHelper(textFiles []string, parsers []Parser) {
	p.textFiles = textFiles
	p.parsers = parsers

	p.isInitialized = true
	p.isParseDone = false
}

// Parse is a func which parse text files.
func (p *ParseHelper) Parse() error {
	if p.isInitialized == false {
		// A caller need call NewParseHelper() func before calling this func.
		return errors.New("This helper has been not initialized yet.")
	}

	// Parse all files
	for _, filePath := range p.textFiles {
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
			os.Exit(1)
			return err
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		// To avoid "token too long" error, set a large buffer.
		max := 10000 * 1024
		buf := []byte{}
		scanner.Buffer(buf, max)

		for scanner.Scan() {
			line := scanner.Text()
			for _, parser := range p.parsers {
				parser.Parse(line)
			}

		}
		if serr := scanner.Err(); serr != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, serr)
			continue
		}
	}

	p.isParseDone = true
	return nil
}

// ShowResult is a func which shows the result of the parse.
func (p *ParseHelper) ShowResult() error {
	if p.isParseDone == false {
		return errors.New("Parse has been not done yet.")
	}

	fmt.Println("--- Result ---")
	for _, parser := range p.parsers {
		parser.PrintResult()
	}
	return nil
}
