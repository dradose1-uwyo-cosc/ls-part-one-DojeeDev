package functions

import (
	"fmt"
	"io"
	"os"
)

// List of colors
type color string

const (
	Reset   color = "\x1b[0m"
	Green   color = "\x1b[32m"
	Blue    color = "\x1b[34m"
)

//Prints appropriate entries in color, or colorless if regular file
func (c color) ColorPrint(w io.Writer, s string) {
	toPrint :=  string(c) + s + string(Reset)
	_, err := w.Write([]byte(toPrint))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error coloring string -> %s <- gave err: %v\n", s, err)
	}
}
