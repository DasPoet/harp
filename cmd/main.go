package main

import (
	"log"
	"os"

	"github.com/daspoet/harp/cmd/cli"
)

func main() {
	if err := cli.App.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
