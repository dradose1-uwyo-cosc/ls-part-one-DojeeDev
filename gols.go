package main

import (
	"fmt"
	"main/functions"
	"os"
	"slices"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting current working directory: %v\n", err)
			return
		}
		functions.SimpleLS(os.Stdout, []string{cwd}, functions.IsTerminal(os.Stdout))
	} else {
		var filepaths []string
		var dirpaths  []string

		for _, a := range args[1:] {
			file, err := os.Lstat(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
				continue
			}

			if file.IsDir() {
				dirpaths = append(dirpaths, a)
			} else {
				filepaths = append(filepaths, a)
			}
		}
		slices.SortFunc(filepaths, func(a, b string) int {
			return strings.Compare(strings.ToLower(a), strings.ToLower(b))
		})
		slices.SortFunc(dirpaths, func(a, b string) int {
			return strings.Compare(strings.ToLower(a), strings.ToLower(b))
		})

		isterm := functions.IsTerminal(os.Stdout)
		functions.SimpleLS(os.Stdout, append(filepaths, dirpaths...), isterm)


	}
}
