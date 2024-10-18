package parser

import (
	"encoding/json"
	"io"
	"os"
)

// Unmarshals the json into `Root`
func ParseFile(filename string) (ctx Context, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	byteContent, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteContent, &ctx)

	return
}
