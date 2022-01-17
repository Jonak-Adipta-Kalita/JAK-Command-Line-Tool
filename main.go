package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "JAK Command Line Tool"
	app.Usage = "Get details about JAK!!"
	app.Description = "A Command Line Tool made by JAK!!"
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
			Name: "class",
			Usage: "Get JAK's Class!!",
			Action: func(c *cli.Context) error {
				fmt.Println("8 A")
				return nil
			},
		},
		{
			Name: "website",
			Usage: "Get JAK's Website's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://jonakadiptakalita.herokuapp.com/")
				return nil
			},
		},
		{
			Name: "api",
			Usage: "Get JAK's API's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://jak-api-dot-com.herokuapp.com/")
				return nil
			},
		},
		{
			Name: "discord_bot",
			Usage: "Get JAK's Discord Bot's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://jak-discord-bot.vercel.app/")
				return nil
			},
		},
		{
			Name: "youtube",
			Usage: "Get JAK's YouTube Channel's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://www.youtube.com/channel/UC6IPfVhkqfcfBZCko6Q9mnQ")
				return nil
			},
		},
		{
			Name: "instagram",
			Usage: "Get JAK's Instagram's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://www.instagram.com/xxjonakadiptaxx/")
				return nil
			},
		},
		{
			Name: "github",
			Usage: "Get JAK's Github's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://github.com/Jonak-Adipta-Kalita")
				return nil
			},
		},
		{
			Name: "twitter",
			Usage: "Get JAK's Twitter's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://twitter.com/AdiptaJonak")
				return nil
			},
		},
		{
			Name: "discord",
			Usage: "Get JAK's Discord Server's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://discord.gg/S3UfGkW")
				return nil
			},
		},
		{
			Name: "reddit",
			Usage: "Get JAK's Reddit Comunity's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://www.reddit.com/r/BeastNight_TV/")
				return nil
			},
		},
		{
			Name: "twitch",
			Usage: "Get JAK's Twitch's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://www.twitch.tv/jonakadiptakalita_2596/")
				return nil
			},
		},
		{
			Name: "itch_io",
			Usage: "Get JAK's Itch.io's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://jonak-adipta-kalita.itch.io/")
				return nil
			},
		},
		{
			Name: "spotify",
			Usage: "Get JAK's Spotify's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://open.spotify.com/user/31cypdycu52u6rj3bsfcldmqrlji")
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
		{
			Name: "mother_name",
			Usage: "Get JAK's Mother's Full Name",
			Action: func(c * cli.Context) error {
				fmt.Println("Namita Deka")
				return nil
			},
		},
		{
			Name: "father_name",
			Usage: "Get JAK's Father's Full Name",
			Action: func(c *cli.Context) error {
				fmt.Println("Hemanta Kalita")
				return nil
			},
		},
		{
			Name: "vscode_extension",
			Usage: "Get JAK's VSCode Extension's Link!!",
			Action: func(c *cli.Context) error {
				fmt.Println("Opening in Browser!!")
				openBrowser("https://marketplace.visualstudio.com/items?itemName=JAKVSCodeExtension.jak-vscode-extension")
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
