package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"phonebook/book"
	"phonebook/logger"
)

type handler func(book.PhoneBook, []string) error

func main() {
	phoneBook := make(book.PhoneBook)

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
		command, args := parts[0], parts[1:]

		switch command {
		case "add":
			handleCommand(phoneBook, doAdd, args)
		case "update":
			handleCommand(phoneBook, doUpdate, args)
		case "get":
			handleCommand(phoneBook, doGet, args)
		case "delete":
			handleCommand(phoneBook, doDelete, args)
		case "list":
			handleCommand(phoneBook, doList, args)
		case "exit":
			logger.Info("Exiting phonebook...")
			return
		default:
			logger.Warn(errors.New("unsupported command. Try 'add', 'get', 'delete', 'update', 'list' or 'exit'"))
		}
	}
}

func handleCommand(book book.PhoneBook, cmd handler, args []string) {
	if err := cmd(book, args); err != nil {
		logger.Warn(err, "command failed")
	}
}

func doGet(book book.PhoneBook, args []string) error {
	if len(args) < 1 {
		return errors.New("invalid format. Use: get name")
	}

	name := args[0]

	record, err := book.Get(name)
	if err != nil {
		return err
	}

	unixUpdatedAt := time.Unix(record.LastUpdatedAt, 0)

	logger.Info(
		fmt.Sprintf(
			"Number for %s is %s (last upated at %s)\n",
			name,
			record.Number,
			unixUpdatedAt.Format("2006-01-02 15:04:05"),
		),
	)

	return nil
}

func doList(book book.PhoneBook, _ []string) error {
	if len(book) == 0 {
		return errors.New("phonebook is empty")
	} else {
		for name, record := range book {
			logger.Info(
				fmt.Sprintf(
					"%s -> %s (updated at: %s)\n",
					name,
					record.Number,
					time.Unix(record.LastUpdatedAt, 0).Format("2006-01-02 15:04:05"),
				),
			)
		}
	}
	return nil
}

func doAdd(book book.PhoneBook, args []string) error {
	record := strings.SplitN(args[0], "=", 2)
	if len(record) != 2 {
		return errors.New("invalid format. Use: add name=number")
	}
	name, number := record[0], record[1]
	err := book.Add(name, number)
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("Added an entry: %s -> %s\n", name, number))
	return nil
}

func doUpdate(book book.PhoneBook, args []string) error {
	record := strings.SplitN(args[0], "=", 2)
	if len(record) != 2 {
		return errors.New("invalid format. Use: update name=number")
	}
	name, number := record[0], record[1]
	err := book.Update(name, number)
	if err != nil {
		return err
	}
	logger.Info(
		fmt.Sprintf(
			"Updated an entry: %s -> %s\n",
			name,
			number,
		),
	)
	return nil
}

func doDelete(book book.PhoneBook, args []string) error {
	if len(args) < 1 {
		return errors.New("invalid format. Use: delete name")
	}

	name := args[0]

	err := book.Delete(name)
	if err != nil {
		return err
	}
	return nil
}
