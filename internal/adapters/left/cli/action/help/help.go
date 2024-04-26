package help

import (
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
)

const (
	template = `Available commands are
- list				Displays names list of available commands to work with
- add				Adds a command to command list
- edit				Edit a command using command name as identifier
- delete			Deletes a command using command name as identifier
- run 				Runs a command using command name as identifier
- display			Displays command details using command name as identifier
- help				Displays this message

If you want additional information about how to use specific command, type
$ scriptRunner [command] --help
`
)

type help types.ActionStruct

func Create() *help {
	return &help{}
}

func (a help) Run() {
	renderer.Render(template, nil)
}

func (a help) Help() {
	a.Run()
}
