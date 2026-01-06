package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, fileName := range os.Args {
		err := readFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			os.Exit(1)
		}
	}
}

func readFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scanerr: %s", err)
	}
	return nil
}
