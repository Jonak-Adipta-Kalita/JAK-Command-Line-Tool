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
			Usage: "Get JAK's Full Name!!",
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
		{
			Name: "website",
			Usage: "Get JAK's Website's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://jonakadiptakalita.herokuapp.com/")
				return nil
			},
		},
		{
			Name: "api",
			Usage: "Get JAK's API's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://jak-api-dot-com.herokuapp.com/")
				return nil
			},
		},
		{
			Name: "discord_bot",
			Usage: "Get JAK's Discord Bot's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://discord.com/oauth2/authorize?client_id=756402881913028689&scope=bot")
				return nil
			},
		},
		{
			Name: "youtube",
			Usage: "Get JAK's YouTube Channel's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://www.youtube.com/channel/UC6IPfVhkqfcfBZCko6Q9mnQ")
				return nil
			},
		},
		{
			Name: "instagram",
			Usage: "Get JAK's Instagram's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://www.instagram.com/xxjonakadiptaxx/")
				return nil
			},
		},
		{
			Name: "github",
			Usage: "Get JAK's Github's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://github.com/Jonak-Adipta-Kalita")
				return nil
			},
		},
		{
			Name: "twitter",
			Usage: "Get JAK's Twitter's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://twitter.com/AdiptaJonak")
				return nil
			},
		},
		{
			Name: "discord",
			Usage: "Get JAK's Discord Server's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://discord.gg/S3UfGkW")
				return nil
			},
		},
		{
			Name: "reddit",
			Usage: "Get JAK's Reddit Comunity's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://www.reddit.com/r/BeastNight_TV/")
				return nil
			},
		},
		{
			Name: "twitch",
			Usage: "Get JAK's Twitch's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://www.twitch.tv/jonakadiptakalita_2596/")
				return nil
			},
		},
		{
			Name: "itch_io",
			Usage: "Get JAK's Itch.io's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://jonak-adipta-kalita.itch.io/")
				return nil
			},
		},
		{
			Name: "spotify",
			Usage: "Get JAK's Spotify's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("https://open.spotify.com/user/31cypdycu52u6rj3bsfcldmqrlji")
				return nil
			},
		},
		{
			Name: "school",
			Usage: "Get JAK's School's Full Name",
			Action: func(c *cli.Context) error {
				fmt.Println("Kendriya Vidyalaya, Mangaldoi")
				return nil
			},
		},
	}
	
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
