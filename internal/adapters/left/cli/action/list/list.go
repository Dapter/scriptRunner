package list

import (
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/renderer"
	"github.com/Dapter/scriptRunner/internal/adapters/left/cli/action/types"
	"github.com/Dapter/scriptRunner/internal/interfaces/ports"
	"log"
)

const template = `Available command list{{range $key, $value := .}}
	- {{$key}}{{end}}
`

const helpTemplate = `List command simply doesn't require any additional parameters
`

type list types.ActionStruct

func Create(c *ports.Left) *list {
	return &list{
		Cli: c,
	}
}

func (a list) Run() {
	cli := *a.Cli
	mainCommands, err := cli.GetList()

	if err != nil {
		log.Fatalf("cannot render command list: %v", err)
	}

	renderer.Render(template, mainCommands)
}

func (a list) Help() {
	renderer.Render(helpTemplate, nil)
}
