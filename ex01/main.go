/*
Convert converts image formats.
It changes the format from the format specified by the -i flag
to the format specified by the -o flag.

It reports an error if a path is not given, or there are multiple paths given,
or the files in the path do not match the input format.

Usage:

	convert [flags] path

The flags are:

	-i
		Input file format, such as jpg, png, gif.
		Default format is jpg.
	-o
		Output file format, such as jpg, png, gif.
		Default format is png.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"vogsphere-v2.42tokyo.jp/vogsphere_admin/road-to-mercari-gopher-dojo-00-tafujise-01/ex01/imgconv"
)

var (
	rawInput  string
	rawOutput string
	input     imgconv.Format
	output    imgconv.Format
	err       error
	root      string
)

func init() {
	flag.StringVar(&rawInput, "i", "jpg", "input image format")
	flag.StringVar(&rawOutput, "o", "png", "output image format")
}

func main() {
	flag.Parse()
	root = parseArgs(flag.Args())
	input, err = imgconv.ParseFormat(rawInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	output, err = imgconv.ParseFormat(rawOutput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("%s: %w", path, errors.Unwrap(err))
			}
			if info.IsDir() {
				return nil
			}
			if input.Validate(path) == false {
				return fmt.Errorf("%s is not a valid file", path)
			}
			if input.Match(path) {
				err = imgconv.Convert(path, output)
				if err != nil {
					return err
				}
				return nil
			}
			return nil
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}

func parseArgs(args []string) string {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(1)
	}
	return args[0]
}
