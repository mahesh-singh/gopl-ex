/*
Modify the dup 2 to print the names of all the files in which each duplicate line occours.
CTRD+D will break the line
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupFiles)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			countLines(f, counts, dupFiles)
			f.Close()
		}
	}

	for line, count := range counts {
		fmt.Printf("%d\t%s\t%v\n", count, line, dupFiles[line])
	}
}

func countLines(f *os.File, counts map[string]int, dupFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		dupFiles[text] = append(dupFiles[text], f.Name())
		counts[text]++
	}
}
