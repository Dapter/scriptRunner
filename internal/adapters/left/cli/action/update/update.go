package update

import (
	"fmt"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/data/command"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
	"os"
)

const helpTemplate = `To edit command you have to add stored command name as a parameter, and a runnable command to be replaced
$ scriptRunner edit "command name" "runnable command"
`

type update types.ActionStruct

func Create(c *ports.Left) *update {
	return &update{
		Cli:        c,
		Parameters: getParameters(),
	}
}

func (a update) Run() {
	cli := *a.Cli
	cmd := command.Command{Name: a.Parameters["commandName"], Commands: []string{a.Parameters["command"]}}
	err := cli.UpdateCommand(cmd)

	if err != nil {
		log.Fatalf("cannot edit command %s: %v", a.Parameters["command"], err)
	}

	fmt.Println("Command edited successfully")
}

func (a update) Help() {
	renderer.Render(helpTemplate, nil)
}

func getParameters() map[string]string {
	params := make(map[string]string)

	fmt.Println(os.Args)

	if len(os.Args) < 4 {
		log.Fatalf("Lack of required parameters, type --help to check how to use this command")
	}

	params["commandName"] = os.Args[2]
	params["command"] = os.Args[3]

	return params
}
