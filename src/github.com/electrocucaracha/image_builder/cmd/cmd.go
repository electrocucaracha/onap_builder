package cmd

import (
	"os"

	"github.com/kballard/go-shellquote"

	"github.com/electrocucaracha/image_builder/utils"
)

type Cmd struct {
	Name   string
	Args   []string
	Stdin  *os.File
	Stdout *os.File
	Stderr *os.File
	Dir    string
}

func NewCmd(cmd string) *Cmd {
	cmds, err := shellquote.Split(cmd)
	utils.Check(err)

	name := cmds[0]
	args := make([]string, 0)
	for _, arg := range cmds[1:] {
		args = append(args, arg)
	}
	return &Cmd{
		Name:   name,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (cmd *Cmd) WithArg(arg string) *Cmd {
	cmd.Args = append(cmd.Args, arg)
	return cmd
}

func (cmd *Cmd) WithArgs(args ...string) *Cmd {
	for _, arg := range args {
		cmd.WithArg(arg)
	}
	return cmd
}
