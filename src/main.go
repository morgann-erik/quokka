package main

import (
	"os"

	"github.com/morgann-erik/quokka/repl"
)

func main() {
    repl.Start(os.Stdin, os.Stdout)
}
