package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Unmarshals the json into `Root`
func ParseFile(filename string) (root Root, err error) {
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

	err = json.Unmarshal(byteContent, &root)

	return
}
