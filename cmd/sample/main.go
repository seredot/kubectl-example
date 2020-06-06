package main

import (
	"fmt"
	"os"
)

func printUsage(failure bool) {
	if failure {
		fmt.Println("Not enough arguments")
	}
	fmt.Println("Usage: sample <RESOURCE_NAME>")
	if failure {
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage(true)
		return
	}

	param := args[1]
	if param == "-h" || param == "--help" {
		printUsage(false)
		return
	}

	resource, err := resourceName(param)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resource)
}
