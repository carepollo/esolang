package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/carepollo/sexlang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Initializing SEEEXXXXX programming language, user: ", user.Username)
	fmt.Println("You may now start writting")
	repl.Start(os.Stdin, os.Stdout)
}
