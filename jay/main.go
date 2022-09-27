package main

import (
	"context"
	"fmt"
	"jay/jay/repl"
	"os"
	"os/signal"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	signal.NotifyContext(ctx, os.Interrupt, os.Kill)

	fmt.Printf("Hello %s! This is the Jay programming language\n", user.Username)

	fmt.Printf("Please input Jay code to execute...\n")
	repl.Start(ctx, os.Stdin, os.Stdout)
}
