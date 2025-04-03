package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tomato-net/incompatible"
)

func main() {
	var help bool
	flag.BoolVar(&help, "h", false, "Display help text")
	flag.Parse()

	if help {
		renderHelp()
		os.Exit(0)
	}

	results, err := incompatible.Analyse()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range results {
		fmt.Println(r)
	}

	if len(results) > 0 {
		os.Exit(1)
	}
}

func renderHelp() {
	fmt.Println(`Usage: ` + os.Args[0] + `
A linter that checks your go.mod to ensure there are no direct dependencies with non-modular packages greater than v1.
Flags:`)
	flag.PrintDefaults()
}
