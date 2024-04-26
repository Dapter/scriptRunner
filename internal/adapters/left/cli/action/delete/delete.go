package delete

import (
	"fmt"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
	"os"
)

const helpTemplate = `To delete command you have to add stored command name only as a parameter
$ scriptRunner delete "command name"
`

type delete types.ActionStruct

func Create(c *ports.Left) *delete {
	return &delete{
		Cli:        c,
		Parameters: getParameters(),
	}
}

func (a delete) Run() {
	cli := *a.Cli
	err := cli.RemoveCommand(a.Parameters["command"])

	if err != nil {
		log.Fatalf("cannot delete command %s: %v", a.Parameters["command"], err)
	}

	fmt.Println("Command deleted successfully")
}

func (a delete) Help() {
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
