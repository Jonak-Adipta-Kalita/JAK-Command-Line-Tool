package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "JAK Command Line Tool"
	app.Usage = "Get details about JAK!!"
	
	app.Commands = []cli.Command {
		{
			Name: "name",
			Usage: "Get JAK's full name!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Jonak Adipta Kalita")
				return nil
			},
		},
		{
			Name: "class_and_section",
			Usage: "Get JAK's Class and Section!!",
			Action: func(c *cli.Context) error {
				fmt.Println("8 A")
				return nil
			},
		},
		{
			Name: "class",
			Usage: "Get JAK's Class!!",
			Action: func(c *cli.Context) error {
				fmt.Println("8")
				return nil
			},
		},
		{
			Name: "section",
			Usage: "Get JAK's Section!!",
			Action: func(c *cli.Context) error {
				fmt.Println("A")
				return nil
			},
		},
	}
	
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
