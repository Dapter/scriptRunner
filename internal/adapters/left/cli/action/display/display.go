package display

import (
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
	"os"
)

const helpTemplate = `To display command details you have to add stored command name only as a parameter
$ scriptRunner display "command name"
`

const template = `Command name: "{{.Name}}"
Runnable commands: {{range $i, $value := .Commands}}
	- "{{$value}}"{{end}}
`

type display types.ActionStruct

func Create(c *ports.Left) *display {
	return &display{
		Cli:        c,
		Parameters: getParameters(),
	}
}

func (a display) Run() {
	cli := *a.Cli
	cmd, err := cli.GetCommand(a.Parameters["command"])

	if err != nil {
		log.Fatalf("cannot display command details for %s: %v", a.Parameters["command"], err)
	}

	renderer.Render(template, cmd)
}

func (a display) Help() {
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
