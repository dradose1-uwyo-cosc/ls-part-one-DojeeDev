package functions

import (
	"os"
	"slices"
)

// func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S

//removes any hidden files from the dir listing, note not exported
func dirFilter(entries []os.DirEntry) []os.DirEntry {
	return slices.DeleteFunc(entries, func(d os.DirEntry) bool {
		return d.Name()[0] == '.'
	})
}
