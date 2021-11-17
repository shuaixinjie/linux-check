package main

import (
	"fmt"
	"github.com/shuaixinjie/linux-check/cmd"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "test",
			Usage:  "test --uid=x --username=y",
			Action: (&cmd.Command{}).Test,
			Flags: []cli.Flag{
				cli.IntFlag{Name: "uid", Usage: "--uid"},
				cli.StringFlag{Name: "username", Usage: "--username"},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Print("command error :" + err.Error())
	}
}
