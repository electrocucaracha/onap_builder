package cmd

import (
	"errors"
	"os"

	"github.com/kballard/go-shellquote"
	pkgerrors "github.com/pkg/errors"
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
	if bashCmd == "" {
		return nil, errors.New("Invalid Argument")
	}
	params, err := shellquote.Split(bashCmd)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "Split bash arguments")
	}

	name := params[0]
	cmd := Cmd{
		Name:   name,
		Args:   params[1:],
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return &cmd, nil
}

// WithArg adds a new argument to the existing list of arguments
func (cmd *Cmd) WithArg(arg string) *Cmd {
	if arg != "" {
		cmd.Args = append(cmd.Args, arg)
	}
	return cmd
}

// WithArgs adds a list of arguments to the existing list of arguments
func (cmd *Cmd) WithArgs(args ...string) *Cmd {
	for _, arg := range args {
		cmd.WithArg(arg)
	}
	return cmd
}
