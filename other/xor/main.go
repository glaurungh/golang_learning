package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cypherer"
)

var (
	mode      = flag.String("mode", "cypher", "Choose mode: cypher or decypher.")
	secretKey = flag.String("secret", "", "Your secret key. Mandatory.")
)

func main() {
	flag.Parse()

	if len(*secretKey) == 0 {
		fmt.Println("No secret key provided.")
		os.Exit(1)
	}

	switch *mode {
	case "cypher":
		textToEncrypt := getUserInput("Enter your text to encrypt: ")
		encryptedText, err := cypherer.Encrypt(textToEncrypt, *secretKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf(
			"Result:\n%s\n",
			encryptedText,
		)
	case "decypher":
		textToDecrypt := getUserInput("Enter encypted text: ")
		decryptedText, err := cypherer.Decrypt(textToDecrypt, *secretKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf(
			"Result:\n%s\n",
			decryptedText,
		)
	default:
		fmt.Println("Invalid mode. Choose 'cypher' or 'decypher'")
		os.Exit(1)
	}

}

func getUserInput(prompt string) string {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error occured while reading string")
			continue
		}
		return strings.TrimRight(result, "\n")
	}
}
