package main

import (
	"github.com/urfave/cli"
	"fmt"
)

func MyCommand(c *cli.Context) {
	fmt.Println("invoked: " + c.Command.Name)
}
