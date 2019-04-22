package main

import (
	"fmt"

	"github.com/odysseus/vigenere"
)

func main() {
	key := "Typewriter"
	message := "Now is the time for all good men to come to the aid of their country"

	encoded := vigenere.Encipher(message, key)
	decoded := vigenere.Decipher(encoded, key)

	fmt.Print(encoded, decoded)
}
