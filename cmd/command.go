package cmd

import (
	"fmt"
	"github.com/shuaixinjie/linux-check/process"
	"github.com/urfave/cli"
)

type Command struct {
}

func (c *Command) Test(cli *cli.Context) {
	uid := cli.Int("uid")
	username := cli.String("username")
	fmt.Println(uid, username)
}

func (c *Command) GetProcessInfo(cli *cli.Context) {
	svc := process.NewProcSvc()
	get := svc.Get()
	for _, proc := range get {
		fmt.Println(proc)
	}
}
