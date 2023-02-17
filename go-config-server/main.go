package main

import (
	"log"
	"os"

	"github.com/gota33/go-config-server/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx, cancel := internal.NewAppContext()
	defer cancel()

	app := cli.App{
		Name:     "config-server",
		Version:  "dev",
		Commands: []*cli.Command{internal.CmdWeb},
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Fatal("Error: ", err.Error())
	}
}
