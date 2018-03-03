package gosh

import (
	"io"
	"os/exec"
	"strings"
)

type Proc interface {
	SetStderr(io.Writer)
	Run() error
}

type OutProc interface {
	SetStdout(io.Writer)
}

type InProc interface {
	SetStdin(io.Reader)
}

type StdProc struct {
	*exec.Cmd
}

func NewStdProc(command string) (*StdProc, error) {
	cmd := strings.Split(command, " ")
	if len(cmd) > 1 {
		return &StdProc{
			Cmd: exec.Command(cmd[0], cmd[1:]...),
		}, nil
	}
	return &StdProc{
		Cmd: exec.Command(cmd[0]),
	}, nil
}

func (sp *StdProc) SetStderr(w io.Writer) { sp.Cmd.Stderr = w }
func (sp *StdProc) SetStdout(w io.Writer) { sp.Cmd.Stdout = w }
func (sp *StdProc) SetStdin(r io.Reader)  { sp.Cmd.Stdin = r }
func (sp *StdProc) Run() error            { return sp.Cmd.Run() }
