package parsehelper

import (
	"fmt"
	"strings"
	"testing"
)

// A parser for tests

type testParser struct {
	// Information to show as the result of the parse
	numberOfINFO int
}

func (p *testParser) Parse(line string) {
	if strings.Index(line, "INFO") != -1 {
		p.numberOfINFO++
	}
}

func (p *testParser) PrintResult() {
	fmt.Println("Number of INFO: ", p.numberOfINFO)
}

// Tests

func TestNewParseHelper(t *testing.T) {
	// Put empty information
	files := []string{}
	parsers := []Parser{}
	helper := ParseHelper{}
	if helper.isInitialized == true {
		t.Errorf("The default value of isInitialized was wrong.")
	}
	if helper.isParseDone == true {
		t.Errorf("The value of isParseDone was wrong.")
	}

	helper.NewParseHelper(files, parsers)

	if helper.isInitialized == false {
		t.Errorf("The initialization was not correctly done.")
	}
	if helper.isParseDone == true {
		t.Errorf("The value of isParseDone was wrong.")
	}
}

func TestParse(t *testing.T) {
	var helper ParseHelper
	var err error
	var files []string
	var parsers []Parser

	// Error case - Calling Parse() before the initialization
	helper = ParseHelper{}
	err = helper.Parse()
	if err == nil {
		t.Errorf("The func didn't return an error.")
	}

	// Put empty information
	files = []string{}
	parsers = []Parser{}
	helper = ParseHelper{}
	helper.NewParseHelper(files, parsers)
	err = helper.Parse()
	if err != nil {
		t.Errorf("The target func returns an error: [%s].", err)
	}
	if helper.isParseDone == false {
		t.Errorf("The value of isParseDone was wrong.")
	}

	// Normal test - Use test data
	files = []string{}
	files = append(files, "test_data.txt")
	parsers = []Parser{}
	var p Parser = &testParser{}
	parsers = append(parsers, p)
	helper = ParseHelper{}
	helper.NewParseHelper(files, parsers)
	err = helper.Parse()
	if err != nil {
		t.Errorf("The target func returns an error: [%s].", err)
	}
	if helper.isParseDone == false {
		t.Errorf("The value of isParseDone was wrong.")
	}
}

func TestShowResult(t *testing.T) {
	var helper ParseHelper
	var err error
	var files []string
	var parsers []Parser

	// Normal test - Use test data
	files = []string{}
	files = append(files, "test_data.txt")
	parsers = []Parser{}
	tp := testParser{}
	var p Parser = &tp
	parsers = append(parsers, p)
	helper = ParseHelper{}
	helper.NewParseHelper(files, parsers)
	err = helper.Parse()
	if err != nil {
		t.Errorf("The target func returns an error: [%s].", err)
	}
	if helper.isParseDone == false {
		t.Errorf("The value of isParseDone was wrong.")
	}
	err = helper.ShowResult()
	if err != nil {
		t.Errorf("The target func returns an error: [%s].", err)
	}
	if tp.numberOfINFO != 6 {
		t.Errorf("Parsing the test data failed. numberOfINFO: [%d].", tp.numberOfINFO)
	}
}
