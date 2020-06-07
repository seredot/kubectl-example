package main

import (
	"bufio"
	"fmt"
	"net/http"
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

func downloadAndPrintResource(name string) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/seredot/k8s-resource-sample/master/resources/%s.yaml", name)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading resource %v: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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

	downloadAndPrintResource(resource)
}
