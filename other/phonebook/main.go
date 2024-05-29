package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	phoneBook := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to phonebook")
	fmt.Println("Available commans: add, get, delete, update, list, exit")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add", "update":
			record := strings.SplitN(parts[1], "=", 2)
			if len(record) != 2 {
				fmt.Println("Invalid format. Use: add name=number")
				continue
			}
			name, number := record[0], record[1]
			phoneBook[name] = number
			fmt.Printf("Added/Updated: %s -> %s\n", name, number)
		case "get":
			name := parts[1]
			number, exists := phoneBook[name]
			if !exists {
				fmt.Printf("Number for %s not found\n", name)
			} else {
				fmt.Printf("Number for %s is %s\n", name, number)
			}
		case "delete":
			name := parts[1]
			_, exists := phoneBook[name]
			if !exists {
				fmt.Printf("Number for %s not found\n", name)
			} else {
				delete(phoneBook, name)
				fmt.Printf("Number for %s has been deleted\n", name)
			}
		case "list":
			if len(phoneBook) == 0 {
				fmt.Println("Phonebook is empty")
			} else {
				for name, phone := range phoneBook {
					fmt.Printf("%s -> %s\n", name, phone)
				}
			}
		case "exit":
			fmt.Println("Exiting phonebook...")
			return
		default:
			fmt.Println("Unsupported command. Try 'add', 'get', 'delete', 'update', 'list' or 'exit'")
		}
	}
}
