package run

import (
	"fmt"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
	"os"
)

const helpTemplate = `To run command you have to add stored command name only as a parameter
$ scriptRunner run "command name"
`

type run types.ActionStruct

func Create(c *ports.Left) *run {
	return &run{
		Cli:        c,
		Parameters: getParameters(),
	}
}

func (a run) Run() {
	cli := *a.Cli
	outputs := cli.Run(a.Parameters["command"])

	fmt.Println(outputs)
}

func (a run) Help() {
	renderer.Render(helpTemplate, nil)
}

func getParameters() map[string]string {
	params := make(map[string]string)

	if len(os.Args) < 3 {
		log.Fatalf("Lack of required parameters, type --help to check how to use this command")
	}

	params["command"] = os.Args[2]

	return params
}
