package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Enter hex number or 'exit'")
	var input string
	number := new(big.Int)
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if input == "exit" {
			break
		}
		if _, ok := number.SetString(removeHexPrefix(input), 16); !ok {
			fmt.Println("Invalid number")
			continue
		}
		fmt.Println(number)
	}
}

func removeHexPrefix(s string) string {
	return strings.TrimPrefix(s, "0x")
}
