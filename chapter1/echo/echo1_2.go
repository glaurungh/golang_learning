// Echoing command line arguments

package main

import (
	"fmt"
	"os"
)

func main() {

	for i, v := range os.Args[1:] {
		fmt.Printf("%d %s\n", i+1, v)
	}

}
