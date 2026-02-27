package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func printColor(w io.Writer, filename string, dir string) {
	info, err := os.Lstat(filepath.Join(dir, filename))
	if err != nil {

		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}
	mode := info.Mode()
	if info.IsDir() {
		Blue.ColorPrint(w, filename+"\n")
	} else if mode.IsRegular() && (mode & 0111) != 0 {
		Green.ColorPrint(w, filename+"*\n")
		// on my system ls adds a * to the end of executables
	} else {
		Reset.ColorPrint(w, filename+"\n")
	}
}

//A simple ls for when no flags provided
func SimpleLS(w io.Writer, args []string, useColor bool) {
	printHeaders := len(args) > 1

	for _, a := range args {
		file, err := os.Lstat(a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			continue
		}
		// printing for directories
		if file.IsDir() {
			dirs, err := os.ReadDir(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting directory entries: %v\n", err)
				return 
			}
			dirs = dirFilter(dirs)

			slices.SortFunc(dirs, func(a, b os.DirEntry) int {
				return strings.Compare(strings.ToLower(a.Name()), strings.ToLower(b.Name()))
			})

			if printHeaders {
				dirHeader := "\n" + a + ":\n"
				_, err := w.Write([]byte(dirHeader))
				if err != nil {
					fmt.Fprintf(os.Stderr, "error writing to stdout: %v\n", err)
					return 
				}
			}

			for _, d := range dirs {
				if useColor {
					printColor(w, d.Name(), a)
				} else {
					_, err := w.Write(append([]byte(d.Name()), '\n') )
					if err != nil {
						fmt.Fprintf(os.Stderr, "error writing to stdout: %v\n", err)
						return 
					}
				}

			}
		} else { // printing for files
			if useColor {
				printColor(w, a, ".")
			} else {
				_, err := w.Write(append([]byte(a), '\n') )
				if err != nil {
					fmt.Fprintf(os.Stderr, "error writing to stdout: %v\n", err)
					return 

				}
			}

		}

	}
}
