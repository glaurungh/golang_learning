// Echoing command line arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	var output, sep string

	for _, v := range os.Args[1:] {
		output += sep + v
		sep = " "
	}

	fmt.Printf("%s\n", output)

}
