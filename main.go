package main

import (
	"Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is your fav programing language!\n", user.Username)
	fmt.Printf("What do you want?\n")
	repl.Start(os.Stdin, os.Stdout)
}
