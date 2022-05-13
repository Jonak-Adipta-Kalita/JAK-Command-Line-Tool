package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "JAK Command Line Tool"
	app.Description = "A Command Line Tool made by JAK!!"
	
	netFlag := []cli.Flag{
		cli.StringFlag {
			Name: "host",
			Value: "jonakadiptakalita.vercel.app",
		},
	}
	
	app.Commands = []cli.Command{
		{
			Name: "name-servers",
			Usage: "Looks up the Name Servers for a Perticular Host",
			Flags: netFlag,
			Action: func(ctx * cli.Context) error {
				ns, err := net.LookupNS(ctx.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				
				return nil
			},
			Aliases: []string{"ns"},
		},
		{
			Name: "ip",
			Usage: "Looks up the IPs for a particular host",
			Flags: netFlag,
			Action: func(ctx * cli.Context) error {
				ip, err := net.LookupIP(ctx.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				
				return nil
			},
		},
		{
			Name: "cname",
			Usage: "Looks up the CNAME for a particullar host",
			Flags: netFlag,
			Action: func(ctx * cli.Context) error {
				cname, err := net.LookupCNAME(ctx.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name: "mx",
			Usage: "Looks up the MX Records for a particullar host",
			Flags: netFlag,
			Action: func(ctx * cli.Context) error {
				mx, err := net.LookupMX(ctx.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
