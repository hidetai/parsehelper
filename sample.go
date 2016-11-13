package main

import (
	"fmt"
	"strings"

	"github.com/hidetai/parsehelper"
)

// A sample parser to show how to use ParseHelper.
// The sample parser counts the number of line which contains "INFO" in text files.

// Step 1. Define a type for your parser.
type sampleParser struct {
	// Information to show as the result of the parse
	numberOfINFO int
}

// Step 2.Implement what your parser does each line of text files.
// This method is called every time ParseHelper read a new line.
func (p *sampleParser) Parse(line string) {
	if strings.Index(line, "INFO") != -1 {
		p.numberOfINFO++
	}
}

// Step 3. Implement what you print as the result of parsing.
// This method is called once when ParseHelper.ShowResult func is called.
func (p *sampleParser) PrintResult() {
	fmt.Println("Number of INFO: ", p.numberOfINFO)
}

// Show how to use your original parser based on ParseHelper

func main() {
	// files := []string{"hoge.txt", "moge.txt"}
	files := []string{}

	parsers := []parsehelper.Parser{}
	var p parsehelper.Parser = &sampleParser{}
	parsers = append(parsers, p)

	// Step 4. Create an instance of ParseHelper.
	helper := parsehelper.ParseHelper{}

	// Step 5. Register target files and your parsers into the helper
	helper.NewParseHelper(files, parsers)

	// Step 6. Execute parsing the target files
	helper.Parse()

	// Step 7. Show the result of parsing
	helper.ShowResult()
}
