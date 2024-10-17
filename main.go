package main

import (
	"fmt"

	"github.com/SaumitraLohokare/fAPI/parser"
)

func main() {
	fileName := "example.json"

	root, err := parser.ParseFile(fileName)
	if err != nil {
		fmt.Println("[Error]:", err)
		return
	}

	fmt.Println("Selected Port:", root.Port)
}
