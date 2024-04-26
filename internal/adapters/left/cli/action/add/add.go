package add

import (
	"fmt"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/data/command"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
	"os"
)

const helpTemplate = `To add command you have to add stored command name as a parameter, and a runnable command
$ scriptRunner add "command name" "runnable command"
`

type add types.ActionStruct

func Create(c *ports.Left) *add {
	return &add{
		Cli:        c,
		Parameters: getParameters(),
	}
}

func (a add) Run() {
	cli := *a.Cli
	cmd := command.Command{Name: a.Parameters["commandName"], Commands: []string{a.Parameters["command"]}}
	err := cli.AddCommand(cmd)

	if err != nil {
		log.Fatalf("cannot add command %s: %v", a.Parameters["command"], err)
	}

	fmt.Println("Command added successfully")
}

func (a add) Help() {
	renderer.Render(helpTemplate, nil)
}

func getParameters() map[string]string {
	params := make(map[string]string)

	if len(os.Args) < 4 {
		log.Fatalf("Lack of required parameters, type --help to check how to use this command")
	}

	params["commandName"] = os.Args[2]
	params["command"] = os.Args[3]

	return params
}
