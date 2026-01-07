/*
Ft_cat imitates the UNIX cat command.
It concatenates files and print on the standard output.
Unlike the UNIX cat, it doesn't have any options.

Usage:

	ft_cat [FILE]
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		err := conCat(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}
	for _, fileName := range os.Args[1:] {
		err := readFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
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
	err = conCat(os.Stdout, f)
	if err != nil {
		return err
	}
	return nil
}

func conCat(w io.Writer, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner: %s", err)
	}
	return nil
}
