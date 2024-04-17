package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	lines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		lines[input.Text()]++
	}

	for line, cnt := range lines {
		if cnt > 1 {
			fmt.Printf("%d\t%s\n", cnt, line)
		}
	}
}
