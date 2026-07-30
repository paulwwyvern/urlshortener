package main

import (
	myos "os"
)

func main() {
	myos.Exit(0) // want "call os.Exit function"
}
