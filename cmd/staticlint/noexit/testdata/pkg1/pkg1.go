package main

import (
	"os"
)

func Exit(code int) {
	os.Exit(code)
}

func foo() {
	os.Exit(1)
}

func main() {
	Exit(0)

	foo()
	bar()

	func() {
		os.Exit(1) // want "call os.Exit function"
	}()

	os.Exit(0) // want "call os.Exit function"
}
