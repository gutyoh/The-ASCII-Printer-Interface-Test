package main

import (
	"fmt"
	"log"
	"os"
)

type FileName string

// Do not change the contents of the PrintAscii() method!
func (f FileName) PrintAscii() {
	b, err := os.ReadFile(string(f))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(b))
}

// Create the AsciiPrinter interface with the PrintAscii() method below:
type ? interface {
	?
}

func main() {
	// Create the variable 'a' of the AsciiPrinter interface type below:
	var a ?

	// Open and read the file "ascii_art.txt" with the 'a' AsciiPrinter interface:
	a = FileName("ascii_art.txt")

	// Call the PrintAscii() method on the 'a' AsciiPrinter interface below:
	a.?
}
