package parser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Parser struct {
	// The parsed output is a list of (`request line`, `expected output`) as [2]string{}
	Output [][2]string
}

// Parses the given file and stores the output in Parser.Output
func (p *Parser) ParseFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	byteContent, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var accumulator strings.Builder
	var previousOutput string
	var previousRequest string = ""
	content := string(byteContent)
	for _, line := range strings.Split(content, "\n") {
		if isCommandLine(line) {
			if previousRequest != "" {
				previousOutput = accumulator.String()
				accumulator.Reset()
				p.Output = append(p.Output, [2]string{previousRequest, previousOutput})
			}

			previousRequest = line
		} else {
			accumulator.WriteString(line)
		}
	}
	previousOutput = accumulator.String()
	accumulator.Reset()
	p.Output = append(p.Output, [2]string{previousRequest, previousOutput})
}

func isCommandLine(s string) bool {
	return strings.HasPrefix(s, "GET") || strings.HasPrefix(s, "POST") || strings.HasPrefix(s, "PUT") || strings.HasPrefix(s, "DELETE")
}
