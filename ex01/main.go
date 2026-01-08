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

type CLIArgs struct {
	rawInput  *string
	rawOutput *string
	root      string
}

type ConvertOptions struct {
	input  imgconv.Format
	output imgconv.Format
}

func main() {
	args, err := parseCLIArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	ops, err := parseOptions(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	err = convertPath(args.root, ops)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}

}

func parseCLIArgs() (args CLIArgs, err error) {
	args.rawInput = flag.String("i", "jpg", "input image format")
	args.rawOutput = flag.String("o", "png", "output image format")
	flag.Parse()
	args.root, err = parseRootArg(flag.Args())
	return
}

func parseRootArg(args []string) (root string, err error) {
	if len(args) != 1 {
		return "", fmt.Errorf("invalid argument")
	}
	return args[0], nil
}

func parseOptions(args CLIArgs) (ops ConvertOptions, err error) {
	ops.input, err = imgconv.ParseFormat(*args.rawInput)
	if err != nil {
		return
	}
	ops.output, err = imgconv.ParseFormat(*args.rawOutput)
	if err != nil {
		return
	}
	if ops.input == ops.output {
		err = fmt.Errorf("input and output format must be different")
		return
	}
	return
}

func convertPath(root string, ops ConvertOptions) error {
	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("%s: %w", path, errors.Unwrap(err))
			}
			if info.IsDir() {
				return nil
			}
			if ops.input.Validate(path) == false {
				return fmt.Errorf("%s is not a valid file", path)
			}
			if ops.input.Match(path) {
				r, err := imgconv.OpenInput(path)
				if err != nil {
					return err
				}
				defer r.Close()
				w, err := imgconv.CreateOutput(path, ops.output)
				if err != nil {
					return err
				}
				defer w.Close()
				err = imgconv.Convert(r, w, ops.output)
				if err != nil {
					return err
				}
			}
			return nil
		})
}
