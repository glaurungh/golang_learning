package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	result := strings.Join(os.Args[1:], " ")
	fmt.Println(result)
}
