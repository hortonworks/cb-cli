package main

import (
	"github.com/urfave/cli"
	"fmt"
)

func Commander() []cli.Command {
	return []cli.Command{
		{
			Name:        "krisz",
			Description: "testcommand",
			Usage:       "invoke the command and see what happens",
			Action:      MyCommand,
		},
	}
}

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}

var Greeter greeting

func main() {
}
