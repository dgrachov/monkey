package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dgrachov/monkey/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This the the Monkey Programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
