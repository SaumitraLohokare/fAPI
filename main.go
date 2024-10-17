package main

import (
	"fmt"
	"net/http"
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

	ctx, err := parser.ParseFile(fileName)
	if err != nil {
		fmt.Println("[Error]:", err)
		return
	}

	// fmt.Printf("%+v\n\n", ctx) // Printing JSON context for debugging

	fmt.Println("Selected Port:", ctx.Port)

	GenerateHandlers(&ctx)
	fmt.Printf("Generated handlers for %d routes.\n", len(ctx.Routes))

	fmt.Printf("Starting server on: http://127.0.0.1:%d\n\n", ctx.Port)

	fmt.Println("----------------------------------------------------------")

	address := fmt.Sprintf(":%d", ctx.Port)
	http.ListenAndServe(address, nil)
	// TODO: Add graceful handling of ^C
}
