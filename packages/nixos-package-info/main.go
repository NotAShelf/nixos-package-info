package main

import (
	"flag"
	"fmt"
	"os"

	utils "notashelf.dev/nixos-package-info/internal"
)

func main() {
	input := flag.String("input", "", "Input file")
	fullFlag := flag.Bool("full", false, "Include extra data")

	flag.Parse()

	if *input == "" {
		fmt.Println("Please provide an input file")
		os.Exit(1)
	}

	packages, err := utils.ReadFile(*input, *fullFlag)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	jsonData, err := utils.OutputJSON(packages, *fullFlag)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
