package main

import (
	"bufio"
	"fmt"
	"one-time-pad/pad"
	"os"
)

func main() {
	fmt.Println("--- One Time Pad Demo ---")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a key to create a new one-time pad: ")
	scanner.Scan()
	key := scanner.Text()

	fmt.Println("Enter a value to encrypt/decrypt: ")
	scanner.Scan()
	value := scanner.Text()

	oneTimePad := pad.New(key)
	encrypted, err := oneTimePad.Encrypt(value)
	if err != nil {
		fmt.Printf("Error encrypting text: %s\n", err)
		return
	}
	fmt.Printf("Encrypted text: %s\n", encrypted)

	decrypted, err := oneTimePad.Decrypt(encrypted)
	if err != nil {
		fmt.Printf("error decrypting text: %s\n", err)
		return
	}
	fmt.Printf("Text was decrypted back to: %s\n", decrypted)
}
