package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var dupFiles []string
	files := os.Args[1:]
	if len(files) == 0 {
		if hasDups(os.Stdin) {
			fmt.Println("Stdin")
		}
	} else {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if hasDups(f) {
				dupFiles = append(dupFiles, fname)
			}
			f.Close()
		}
	}
	
	for _, file := range dupFiles {
		fmt.Printf("%s\n", file)
	}

}

func hasDups(f *os.File) bool {
	cntMap := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		cntMap[s]++
		if cntMap[s] > 1 {
			return true
		}
	}
	return false
}
