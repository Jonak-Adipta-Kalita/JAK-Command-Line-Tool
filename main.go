package main

import (
	"log"
	"os"

	"github.com/Jonak-Adipta-Kalita/JAK-Command-Line-Tool/networking"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "JAK Command Line Tool"
	app.Description = "This CLI tool combines networking management and a package manager specific to JAK's language, catering to the needs of developers and system administrators."

	app.Commands = networking.Commands

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
