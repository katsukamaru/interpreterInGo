package main

import (
	"interpreterInGo/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stderr)
}
