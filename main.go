package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SaumitraLohokare/fAPI/parser"
)

const ASCII = "" +
	",------. ,---.  ,------. ,--.    \n" +
	"|  .---'/  O  \\ |  .--. '|  |    \n" +
	"|  `--,|  .-.  ||  '--' ||  |    \n" +
	"|  |`  |  | |  ||  | --' |  |    \n" +
	"`--'   `--' `--'`--'     `--'    \n"

func main() {
	args := os.Args
	progName := filepath.Base(args[0])

	if len(args) < 2 {
		fmt.Printf("Usage: %s <input_file>", progName)
		return
	} else if len(args) > 2 {
		fmt.Printf("%s only accepts one input file as of now.", progName)
		fmt.Printf("Usage: %s <input_file>", progName)
		return
	}

	fmt.Printf("%s\n", ASCII)

	fileName := args[1]

	root, err := parser.ParseFile(fileName)
	if err != nil {
		fmt.Println("[Error]:", err)
		return
	}

	fmt.Println("Selected Port:", root.Port)
}
