package main

import (
	"fmt"
	"os"

	"github.com/dgrachov/monkey/monkey/repl"
)

func main() {
	fmt.Println("Hello! This the the Monkey Programming language!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
