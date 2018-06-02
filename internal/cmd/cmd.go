package cmd

import (
	"os"

	"github.com/kballard/go-shellquote"
)

// Cmd is a representation of a bash instruction
type Cmd struct {
	Name   string
	Args   []string
	Stdin  *os.File
	Stdout *os.File
	Stderr *os.File
	Dir    string
}

// NewCmd creates a new command instance from a string bash instruction
func NewCmd(bashCmd string) (*Cmd, error) {
	params, err := shellquote.Split(bashCmd)
	if err != nil {
		return nil, err
	}

	name := params[0]
	args := make([]string, len(params)-1)
	for _, arg := range params[1:] {
		args = append(args, arg)
	}
	cmd := Cmd{
		Name:   name,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return &cmd, nil
}

// WithArg adds a new argument to the existing list of arguments
func (cmd *Cmd) WithArg(arg string) *Cmd {
	cmd.Args = append(cmd.Args, arg)
	return cmd
}

// WithArgs adds a list of arguments to the existing list of arguments
func (cmd *Cmd) WithArgs(args ...string) *Cmd {
	for _, arg := range args {
		cmd.WithArg(arg)
	}
	return cmd
}
