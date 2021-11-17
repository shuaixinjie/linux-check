package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

type Command struct {
}

func (c *Command) Test(cli *cli.Context) {
	uid := cli.Int("uid")
	username := cli.String("username")
	fmt.Println(uid, username)
}
