package book

import (
	"fmt"
	"time"
)

type PhoneBook map[string]PhoneNumber

type PhoneNumber struct {
	Number        string
	LastUpdatedAt int64
}

func (book *PhoneBook) Add(name, number string) error {
	if _, exists := (*book)[name]; exists {
		return fmt.Errorf("name %s already exists", name)
	}
	(*book)[name] = PhoneNumber{
		Number:        number,
		LastUpdatedAt: time.Now().Unix(),
	}
	return nil
}

func (book *PhoneBook) Update(name, number string) error {
	if _, exists := (*book)[name]; !exists {
		return fmt.Errorf("name %s doesn't exist", name)
	}
	(*book)[name] = PhoneNumber{
		Number:        number,
		LastUpdatedAt: time.Now().Unix(),
	}
	return nil
}

func (book *PhoneBook) Delete(name string) error {
	if _, exists := (*book)[name]; !exists {
		return fmt.Errorf("name %s doesn't exist", name)
	}
	delete(*book, name)
	return nil
}

func (book *PhoneBook) Get(name string) (PhoneNumber, error) {
	record, exists := (*book)[name]
	if !exists {
		return PhoneNumber{}, fmt.Errorf("name %s doesn't exist", name)
	}
	return record, nil
}
