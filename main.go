package main

import (
	"log"
	"os"
	"time"

	"github.com/Jonak-Adipta-Kalita/JAK-Command-Line-Tool/networking"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "JAK Command Line Tool"
	app.Compiled = time.Now()
	app.Description = "This CLI tool combines DNS inspector and a package manager specific to JAK's language, catering to the needs of developers and system administrators."
	app.Version = "v1.0.1"
	app.Copyright = "(c) 2021-now Jonak-Adipta-Kalita"

	app.Commands = []cli.Command{
		{
			Name:        "networking",
			Aliases:     []string{"net"},
			Usage:       "Lookup DNS related Information of a Host.",
			Subcommands: networking.Commands,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
